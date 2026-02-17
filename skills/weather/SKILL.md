---
name: weather
version: 0.3.0
author: devclaw
description: "Weather information and forecasts — no API key required"
category: builtin
tags: [weather, forecast, temperature, wttr, open-meteo, moon]
---
# Weather

You can check weather for any location worldwide. **No API key needed.**

## Current weather

```bash
# Quick one-line weather
curl -s "wttr.in/CITY?format=3"

# Detailed current conditions
curl -s "wttr.in/CITY?format=%l:+%c+%t+(feels+like+%f)+💧%h+💨%w"

# Full forecast (3 days, colored terminal output)
curl -s "wttr.in/CITY?lang=pt"

# Compact forecast
curl -s "wttr.in/CITY?format=v2&lang=pt"
```

## JSON format (for structured data)

```bash
# Full weather data as JSON
curl -s "wttr.in/CITY?format=j1" | jq '{
  location: .nearest_area[0].areaName[0].value,
  country: .nearest_area[0].country[0].value,
  temp_c: .current_condition[0].temp_C,
  feels_like_c: .current_condition[0].FeelsLikeC,
  humidity: .current_condition[0].humidity,
  description: .current_condition[0].weatherDesc[0].value,
  wind_kmph: .current_condition[0].windspeedKmph,
  wind_dir: .current_condition[0].winddir16Point,
  visibility_km: .current_condition[0].visibility,
  uv_index: .current_condition[0].uvIndex
}'

# 3-day forecast
curl -s "wttr.in/CITY?format=j1" | jq '.weather[] | {
  date: .date,
  max_c: .maxtempC,
  min_c: .mintempC,
  avg_c: .avgtempC,
  sun_hours: .sunHour,
  description: .hourly[4].weatherDesc[0].value
}'
```

## Moon phase

```bash
curl -s "wttr.in/Moon"
curl -s "wttr.in/Moon?format=j1" | jq '.weather[0].astronomy[0]'
```

## Open-Meteo API (alternative, more precise)

```bash
# Current weather by coordinates (no API key needed)
curl -s "https://api.open-meteo.com/v1/forecast?latitude=LAT&longitude=LON&current=temperature_2m,relative_humidity_2m,wind_speed_10m,weather_code" | jq '.current'

# 7-day forecast
curl -s "https://api.open-meteo.com/v1/forecast?latitude=LAT&longitude=LON&daily=temperature_2m_max,temperature_2m_min,precipitation_sum,weather_code&timezone=auto" | jq '.daily'

# Get coordinates from city name (geocoding)
curl -s "https://geocoding-api.open-meteo.com/v1/search?name=CITY&count=1" | jq '.results[0] | {name, latitude, longitude, country}'
```

## Tips

- Replace `CITY` with the city name. Use `+` for spaces: `New+York`.
- Use `lang=pt` for Portuguese, `lang=en` for English.
- wttr.in accepts airport codes (GRU, JFK), coordinates (48.8,2.3), and city names.
- The user's timezone and location may be in USER.md — use them as defaults.
- If the user doesn't specify a location, ask or use their default.
- Open-Meteo is better for precise forecasts and historical data.

## Triggers

weather, what's the weather, temperature, will it rain, forecast,
moon phase, previsão do tempo, vai chover, clima em, temperatura
