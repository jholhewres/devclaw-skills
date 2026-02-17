---
name: calculator
version: 0.2.0
author: devclaw
description: "Mathematical calculations, conversions, and computations"
category: builtin
tags: [math, calculator, conversion, computation]
---
# Calculator

You can perform mathematical calculations and unit conversions.

## Calculations

```bash
# Basic arithmetic
echo "scale=10; 3.14159 * (5^2)" | bc -l

# Complex expressions
echo "scale=6; sqrt(144) + e(1)" | bc -l

# Percentages
echo "scale=2; 1500 * 15 / 100" | bc -l

# Python for complex math (if available)
python3 -c "
import math
result = math.sqrt(2) * math.pi
print(f'{result:.6f}')
"
```

## Unit conversions

```bash
# Temperature: Celsius ↔ Fahrenheit
python3 -c "c=30; print(f'{c}°C = {c*9/5+32:.1f}°F')"
python3 -c "f=86; print(f'{f}°F = {(f-32)*5/9:.1f}°C')"

# Distance
python3 -c "km=10; print(f'{km} km = {km*0.621371:.2f} miles')"
python3 -c "mi=6.2; print(f'{mi} miles = {mi*1.60934:.2f} km')"

# Weight
python3 -c "kg=75; print(f'{kg} kg = {kg*2.20462:.1f} lbs')"

# Data storage
python3 -c "gb=500; print(f'{gb} GB = {gb/1024:.2f} TB = {gb*1024:.0f} MB')"

# Currency (use web-search for current rates)
# exchange rates change — always search for current rate first
```

## Date & time calculations

```bash
# Days between dates
python3 -c "
from datetime import datetime
d1 = datetime(2026, 1, 1)
d2 = datetime(2026, 12, 31)
print(f'{(d2-d1).days} days')
"

# Add days to a date
python3 -c "
from datetime import datetime, timedelta
d = datetime(2026, 2, 14) + timedelta(days=30)
print(d.strftime('%Y-%m-%d (%A)'))
"
```

## Tips

- Use `bc -l` for quick arithmetic (supports scale, sqrt, e, etc.).
- Use `python3 -c` for more complex calculations with `math` module.
- For currency conversions, always search for the latest exchange rate first.
- Show your work when doing multi-step calculations.
- Round to a reasonable number of decimal places for the context.

## Triggers

calculate, what is, convert, how much is, calcular, converter, quanto é
