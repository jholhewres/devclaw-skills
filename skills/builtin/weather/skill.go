// Package weather implements the weather skill for GoClaw.
// Uses free APIs (wttr.in and Open-Meteo) — no API key required.
package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// WeatherSkill provides weather information and forecasts.
type WeatherSkill struct {
	provider    string // "wttr" or "openmeteo"
	defaultCity string
	units       string // "metric" or "imperial"
	language    string
	client      *http.Client
}

// New creates a new WeatherSkill instance.
func New(config map[string]any) (*WeatherSkill, error) {
	s := &WeatherSkill{
		provider:    "wttr",
		defaultCity: "São Paulo",
		units:       "metric",
		language:    "en",
		client:      &http.Client{Timeout: 15 * time.Second},
	}
	s.applyConfig(config)
	return s, nil
}

// Init initializes the skill with the provided configuration.
func (s *WeatherSkill) Init(_ context.Context, config map[string]any) error {
	s.applyConfig(config)
	return nil
}

// Execute runs the skill with raw text input.
func (s *WeatherSkill) Execute(ctx context.Context, input string) (string, error) {
	location := strings.TrimSpace(input)
	if location == "" {
		location = s.defaultCity
	}
	result, err := s.GetWeather(ctx, map[string]any{"location": location})
	if err != nil {
		return "", err
	}
	b, _ := json.MarshalIndent(result, "", "  ")
	return string(b), nil
}

// Shutdown releases resources.
func (s *WeatherSkill) Shutdown() error {
	s.client.CloseIdleConnections()
	return nil
}

// ---------- Tools ----------

// GetWeather returns current weather for a location.
func (s *WeatherSkill) GetWeather(ctx context.Context, args map[string]any) (any, error) {
	location := s.resolveLocation(args)
	return s.fetchWttr(ctx, location, "current")
}

// GetForecast returns a multi-day weather forecast.
func (s *WeatherSkill) GetForecast(ctx context.Context, args map[string]any) (any, error) {
	location := s.resolveLocation(args)
	days := 3
	if d, ok := toInt(args["days"]); ok && d >= 1 && d <= 7 {
		days = d
	}
	return s.fetchWttrForecast(ctx, location, days)
}

// GetMoon returns the current moon phase.
func (s *WeatherSkill) GetMoon(ctx context.Context, args map[string]any) (any, error) {
	location := s.resolveLocation(args)
	return s.fetchWttrMoon(ctx, location)
}

// ---------- wttr.in ----------

// fetchWttr fetches current weather from wttr.in JSON API.
func (s *WeatherSkill) fetchWttr(ctx context.Context, location, _ string) (any, error) {
	reqURL := fmt.Sprintf("https://wttr.in/%s?format=j1&lang=%s",
		url.PathEscape(location), s.language)

	body, err := s.httpGet(ctx, reqURL)
	if err != nil {
		return nil, fmt.Errorf("wttr.in: %w", err)
	}

	var data map[string]any
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return nil, fmt.Errorf("parsing wttr.in response: %w", err)
	}

	// Extract current conditions from the JSON response.
	return s.parseWttrCurrent(data), nil
}

func (s *WeatherSkill) parseWttrCurrent(data map[string]any) map[string]any {
	result := map[string]any{
		"provider": "wttr.in",
	}

	// Extract nearest area info.
	if areas, ok := data["nearest_area"].([]any); ok && len(areas) > 0 {
		if area, ok := areas[0].(map[string]any); ok {
			result["location"] = extractWttrName(area, "areaName")
			result["region"] = extractWttrName(area, "region")
			result["country"] = extractWttrName(area, "country")
		}
	}

	// Extract current condition.
	if conditions, ok := data["current_condition"].([]any); ok && len(conditions) > 0 {
		if cond, ok := conditions[0].(map[string]any); ok {
			if s.units == "imperial" {
				result["temperature_f"] = cond["temp_F"]
				result["feels_like_f"] = cond["FeelsLikeF"]
			} else {
				result["temperature_c"] = cond["temp_C"]
				result["feels_like_c"] = cond["FeelsLikeC"]
			}
			result["humidity"] = cond["humidity"]
			result["wind_speed_kmph"] = cond["windspeedKmph"]
			result["wind_direction"] = cond["winddir16Point"]
			result["pressure_mb"] = cond["pressure"]
			result["visibility_km"] = cond["visibility"]
			result["uv_index"] = cond["uvIndex"]
			result["cloud_cover"] = cond["cloudcover"]

			// Weather description.
			if descs, ok := cond["weatherDesc"].([]any); ok && len(descs) > 0 {
				if desc, ok := descs[0].(map[string]any); ok {
					result["description"] = desc["value"]
				}
			}
			// Try localized description.
			langKey := "lang_" + s.language
			if descs, ok := cond[langKey].([]any); ok && len(descs) > 0 {
				if desc, ok := descs[0].(map[string]any); ok {
					result["description"] = desc["value"]
				}
			}
		}
	}
	return result
}

// fetchWttrForecast fetches multi-day forecast.
func (s *WeatherSkill) fetchWttrForecast(ctx context.Context, location string, days int) (any, error) {
	reqURL := fmt.Sprintf("https://wttr.in/%s?format=j1&lang=%s",
		url.PathEscape(location), s.language)

	body, err := s.httpGet(ctx, reqURL)
	if err != nil {
		return nil, fmt.Errorf("wttr.in forecast: %w", err)
	}

	var data map[string]any
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	forecasts, ok := data["weather"].([]any)
	if !ok {
		return nil, fmt.Errorf("no forecast data available")
	}

	result := make([]map[string]any, 0, days)
	for i, f := range forecasts {
		if i >= days {
			break
		}
		if day, ok := f.(map[string]any); ok {
			entry := map[string]any{
				"date": day["date"],
			}
			if s.units == "imperial" {
				entry["max_temp_f"] = day["maxtempF"]
				entry["min_temp_f"] = day["mintempF"]
			} else {
				entry["max_temp_c"] = day["maxtempC"]
				entry["min_temp_c"] = day["mintempC"]
			}
			entry["avg_temp_c"] = day["avgtempC"]
			entry["sun_hours"] = day["sunHour"]
			entry["uv_index"] = day["uvIndex"]

			// Extract hourly summary (astronomy).
			if astro, ok := day["astronomy"].([]any); ok && len(astro) > 0 {
				if a, ok := astro[0].(map[string]any); ok {
					entry["sunrise"] = a["sunrise"]
					entry["sunset"] = a["sunset"]
					entry["moonrise"] = a["moonrise"]
					entry["moonset"] = a["moonset"]
					entry["moon_phase"] = a["moon_phase"]
				}
			}

			// Hourly snapshots (simplified).
			if hourly, ok := day["hourly"].([]any); ok {
				hours := make([]map[string]any, 0, len(hourly))
				for _, h := range hourly {
					if hr, ok := h.(map[string]any); ok {
						hEntry := map[string]any{
							"time":              hr["time"],
							"temp_c":            hr["tempC"],
							"chance_of_rain":    hr["chanceofrain"],
							"chance_of_snow":    hr["chanceofsnow"],
							"wind_speed_kmph":   hr["windspeedKmph"],
							"wind_direction":    hr["winddir16Point"],
							"humidity":          hr["humidity"],
							"cloud_cover":       hr["cloudcover"],
							"precipitation_mm":  hr["precipMM"],
						}
						if descs, ok := hr["weatherDesc"].([]any); ok && len(descs) > 0 {
							if d, ok := descs[0].(map[string]any); ok {
								hEntry["description"] = d["value"]
							}
						}
						hours = append(hours, hEntry)
					}
				}
				entry["hourly"] = hours
			}
			result = append(result, entry)
		}
	}
	return result, nil
}

// fetchWttrMoon returns moon phase info.
func (s *WeatherSkill) fetchWttrMoon(ctx context.Context, location string) (any, error) {
	// wttr.in moon endpoint returns ASCII art; use the JSON API instead.
	reqURL := fmt.Sprintf("https://wttr.in/%s?format=j1&lang=%s",
		url.PathEscape(location), s.language)

	body, err := s.httpGet(ctx, reqURL)
	if err != nil {
		return nil, fmt.Errorf("wttr.in moon: %w", err)
	}

	var data map[string]any
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	// Extract astronomy from today's weather.
	if forecasts, ok := data["weather"].([]any); ok && len(forecasts) > 0 {
		if day, ok := forecasts[0].(map[string]any); ok {
			if astro, ok := day["astronomy"].([]any); ok && len(astro) > 0 {
				if a, ok := astro[0].(map[string]any); ok {
					return map[string]any{
						"date":             day["date"],
						"moon_phase":       a["moon_phase"],
						"moon_illumination": a["moon_illumination"],
						"moonrise":         a["moonrise"],
						"moonset":          a["moonset"],
						"sunrise":          a["sunrise"],
						"sunset":           a["sunset"],
					}, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("no moon data available")
}

// ---------- Helpers ----------

func extractWttrName(area map[string]any, key string) string {
	if vals, ok := area[key].([]any); ok && len(vals) > 0 {
		if v, ok := vals[0].(map[string]any); ok {
			if name, ok := v["value"].(string); ok {
				return name
			}
		}
	}
	return ""
}

func (s *WeatherSkill) resolveLocation(args map[string]any) string {
	if loc, ok := args["location"].(string); ok && loc != "" {
		return loc
	}
	return s.defaultCity
}

func (s *WeatherSkill) httpGet(ctx context.Context, reqURL string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "GoClaw/1.0")

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *WeatherSkill) applyConfig(config map[string]any) {
	if p, ok := config["provider"].(string); ok {
		s.provider = p
	}
	if city, ok := config["default_city"].(string); ok {
		s.defaultCity = city
	}
	if units, ok := config["units"].(string); ok {
		s.units = units
	}
	if lang, ok := config["language"].(string); ok {
		s.language = lang
	}
}

func toInt(v any) (int, bool) {
	switch n := v.(type) {
	case int:
		return n, true
	case float64:
		return int(n), true
	case int64:
		return int(n), true
	case json.Number:
		i, err := n.Int64()
		return int(i), err == nil
	}
	return 0, false
}
