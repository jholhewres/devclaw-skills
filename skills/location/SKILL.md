---
name: location
version: 0.1.0
author: devclaw
description: "Location and geocoding — addresses, coordinates, maps, and timezone"
category: utilities
tags: [location, geocoding, maps, coordinates, timezone]
requires:
  bins: [curl, jq]
---
# Location

Get location data, geocoding, coordinates, and maps.

## Geocoding (Address → Coordinates)

### Nominatim (OpenStreetMap - Free)

```bash
# Search address
curl -s "https://nominatim.openstreetmap.org/search?q=1600+Pennsylvania+Avenue,Washington+DC&format=json" | jq '.[0] | {lat, lon, display_name}'

# With limit
curl -s "https://nominatim.openstreetmap.org/search?q=Paris,France&format=json&limit=1" | jq '.[0]'

# Reverse geocoding (coordinates → address)
curl -s "https://nominatim.openstreetmap.org/reverse?lat=48.8584&lon=2.2945&format=json" | jq '.display_name'
```

### Google Geocoding API

```bash
# Setup
export GOOGLE_MAPS_KEY="xxx"

# Geocode address
curl -s "https://maps.googleapis.com/maps/api/geocode/json?address=1600+Pennsylvania+Avenue,Washington+DC&key=$GOOGLE_MAPS_KEY" | jq '.results[0].geometry.location'

# Reverse geocode
curl -s "https://maps.googleapis.com/maps/api/geocode/json?latlng=48.8584,2.2945&key=$GOOGLE_MAPS_KEY" | jq '.results[0].formatted_address'
```

## IP Geolocation

```bash
# Get location from IP (ip-api.com - free)
curl -s "http://ip-api.com/json/" | jq '{country, city, lat, lon, timezone, isp}'

# Specific IP
curl -s "http://ip-api.com/json/8.8.8.8" | jq '.'

# ipinfo.io
curl -s "https://ipinfo.io/json" | jq '.'

# With token
curl -s "https://ipinfo.io/json?token=$IPINFO_TOKEN" | jq '.'
```

## Distance Calculation

```bash
# Haversine formula (great-circle distance)
python3 -c "
from math import radians, sin, cos, sqrt, atan2

def haversine(lat1, lon1, lat2, lon2):
    R = 6371  # Earth radius in km
    dlat = radians(lat2 - lat1)
    dlon = radians(lon2 - lon1)
    a = sin(dlat/2)**2 + cos(radians(lat1)) * cos(radians(lat2)) * sin(dlon/2)**2
    c = 2 * atan2(sqrt(a), sqrt(1-a))
    return R * c

# New York to Los Angeles
distance = haversine(40.7128, -74.0060, 34.0522, -118.2437)
print(f'{distance:.1f} km')
"
```

## Timezone

```bash
# Get timezone from coordinates
curl -s "https://timezoneapi.io/api/ip/?token=$TIMEZONE_TOKEN" | jq '.data.timezone'

# Google Timezone API
curl -s "https://maps.googleapis.com/maps/api/timezone/json?location=48.8584,2.2945&timestamp=$(date +%s)&key=$GOOGLE_MAPS_KEY" | jq '.'

# Current time in timezone
TZ='America/New_York' date
TZ='Europe/London' date
TZ='Asia/Tokyo' date
```

## Static Maps

### Google Static Maps

```bash
# Generate map image
curl -s "https://maps.googleapis.com/maps/api/staticmap?center=48.8584,2.2945&zoom=15&size=600x400&markers=48.8584,2.2945&key=$GOOGLE_MAPS_KEY" -o map.png

# With multiple markers
curl -s "https://maps.googleapis.com/maps/api/staticmap?size=600x400&markers=color:red|48.8584,2.2945&markers=color:blue|48.8606,2.3376&key=$GOOGLE_MAPS_KEY" -o map.png
```

### OpenStreetMap Static

```bash
# Static map via OSM
curl -s "https://staticmap.openstreetmap.de/staticmap.php?center=48.8584,2.2945&zoom=15&size=600x400&markers=48.8584,2.2945,red-pushpin" -o map.png
```

## Places Search

```bash
# Google Places API
curl -s "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=48.8584,2.2945&radius=1000&type=restaurant&key=$GOOGLE_MAPS_KEY" | jq '.results[:5] | .[] | {name, rating, vicinity}'

# Text search
curl -s "https://maps.googleapis.com/maps/api/place/textsearch/json?query=restaurants+in+Paris&key=$GOOGLE_MAPS_KEY" | jq '.results[:5]'
```

## Elevation

```bash
# Google Elevation API
curl -s "https://maps.googleapis.com/maps/api/elevation/json?locations=48.8584,2.2945&key=$GOOGLE_MAPS_KEY" | jq '.results[0].elevation'

# Open-Elevation (free)
curl -s "https://api.open-elevation.com/api/v1/lookup?locations=48.8584,2.2945" | jq '.results[0].elevation'
```

## Address Validation

```bash
# Validate and normalize address
curl -s "https://address-validation.googleapis.com/v1:validateAddress?key=$GOOGLE_MAPS_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "address": {
      "regionCode": "US",
      "addressLines": ["1600 Pennsylvania Avenue", "Washington DC"]
    }
  }' | jq '.result.geocode.location'
```

## Tips

- Nominatim: Free but rate limited (1 req/sec)
- Google Maps: Requires API key, paid after free tier
- Use `Accept-Language` header for localized results
- Cache geocoding results to save API calls
- Handle cases where geocoding fails

## Triggers

location, geocoding, address, coordinates, maps, latitude, longitude,
find location, ip location, timezone
