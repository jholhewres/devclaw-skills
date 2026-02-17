---
name: qr-code
version: 0.1.0
author: devclaw
description: "QR code generation and reading — create and decode QR codes"
category: utilities
tags: [qr, qr-code, barcode, generate]
requires:
  bins: [curl, qrencode]
---
# QR Code

Generate and read QR codes.

## Generate QR Codes

### Option 1: API (Quick)

```bash
# QRServer API (free)
curl -s "https://api.qrserver.com/v1/create-qr-code/?size=300x300&data=https://example.com" -o qr.png

# With customization
curl -s "https://api.qrserver.com/v1/create-qr-code/?size=400x400&color=000000&bgcolor=FFFFFF&data=Hello%20World" -o qr.png

# GoQR API
curl -s "https://api.qrserver.com/v1/create-qr-code/?size=300x300&data=TEXT" -o qr.png

# QR code for WiFi
WIFI_DATA="WIFI:T:WPA;S:NetworkName;P:Password;;"
curl -s "https://api.qrserver.com/v1/create-qr-code/?size=300x300&data=$(echo "$WIFI_DATA" | jq -sRr @uri)" -o wifi-qr.png
```

### Option 2: qrencode (Local)

```bash
# Install
# Ubuntu/Debian: sudo apt install qrencode
# macOS: brew install qrencode

# Basic QR code (terminal)
qrencode -t UTF8 "https://example.com"

# PNG output
qrencode -o qr.png "https://example.com"

# Custom size
qrencode -o qr.png -s 10 -m 2 "https://example.com"

# High error correction
qrencode -o qr.png -l H "https://example.com"

# SVG output
qrencode -o qr.svg -t SVG "https://example.com"

# ASCII output
qrencode -t ASCII "https://example.com"
```

### Option 3: Python (qrcode)

```python
# generate_qr.py
import qrcode
import sys

data = sys.argv[1] if len(sys.argv) > 1 else "https://example.com"
filename = sys.argv[2] if len(sys.argv) > 2 else "qr.png"

qr = qrcode.QRCode(
    version=1,
    error_correction=qrcode.constants.ERROR_CORRECT_L,
    box_size=10,
    border=4,
)
qr.add_data(data)
qr.make(fit=True)

img = qr.make_image(fill_color="black", back_color="white")
img.save(filename)
print(f"QR code saved to {filename}")
```

```bash
pip install qrcode[pil]
python generate_qr.py "https://example.com" qr.png
```

## Read/Decode QR Codes

### Option 1: Online API

```bash
# Decode QR code image
curl -s -X POST "https://api.qrserver.com/v1/read-qr-code/" \
  -F "file=@qr.png" | jq '.[0].symbol[0].data'
```

### Option 2: zbar (Local)

```bash
# Install
# Ubuntu/Debian: sudo apt install zbar-tools
# macOS: brew install zbar

# Decode QR code from image
zbarimg qr.png

# Decode from camera (Linux)
zbarcam

# Quiet output (just data)
zbarimg --raw qr.png
```

### Option 3: Python (pyzbar)

```python
# read_qr.py
from pyzbar.pyzbar import decode
from PIL import Image
import sys

image_path = sys.argv[1] if len(sys.argv) > 1 else "qr.png"
img = Image.open(image_path)

results = decode(img)
for result in results:
    print(result.data.decode('utf-8'))
```

```bash
pip install pyzbar pillow
python read_qr.py qr.png
```

## QR Code Types

```bash
# URL
qrencode -o url.png "https://example.com"

# Email
qrencode -o email.png "mailto:hello@example.com?subject=Hello&body=Hi"

# Phone
qrencode -o phone.png "tel:+5511999999999"

# SMS
qrencode -o sms.png "sms:+5511999999999?body=Hello"

# WiFi
qrencode -o wifi.png "WIFI:T:WPA;S:MyNetwork;P:MyPassword;;"

# vCard
VCARD="BEGIN:VCARD\nVERSION:3.0\nFN:John Doe\nTEL:+5511999999999\nEMAIL:john@example.com\nEND:VCARD"
echo -e "$VCARD" | qrencode -o contact.png -r -

# Bitcoin address
qrencode -o btc.png "bitcoin:bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh"
```

## Tips

- Use error correction level L/M/Q/H for redundancy
- Higher error correction = larger QR codes
- Test QR codes with multiple readers
- Keep QR codes simple for better scanning
- Minimum recommended size: 200x200 pixels

## Triggers

qr code, qr, generate qr, create qr code, barcode, qr generator,
scan qr, read qr
