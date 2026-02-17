---
name: ocr
version: 0.1.0
author: devclaw
description: "OCR — extract text from images and PDFs using various OCR engines"
category: data
tags: [ocr, text recognition, image to text, extraction]
requires:
  bins: [tesseract]
---
# OCR

Extract text from images and scanned documents.

## Option 1: Tesseract (Local)

```bash
# Setup
# Ubuntu/Debian
sudo apt install tesseract-ocr tesseract-ocr-eng tesseract-ocr-por

# macOS
brew install tesseract tesseract-lang

# Basic OCR
tesseract image.png output
cat output.txt

# Portuguese
tesseract image.png output -l por

# Multiple languages
tesseract image.png output -l eng+por

# Output to stdout
tesseract image.png stdout

# PDF output (searchable)
tesseract image.png output pdf
```

## Option 2: Tesseract with Preprocessing

```bash
# Enhance image for better OCR
convert input.png -colorspace Gray -contrast-stretch 0 -resize 300% processed.png
tesseract processed.png output

# Denoise and sharpen
convert input.png -despeckle -sharpen 0x1 -colorspace Gray clean.png
tesseract clean.png output

# Threshold for clearer text
convert input.png -threshold 50% binary.png
tesseract binary.png output
```

## Option 3: Python + pytesseract

```python
# ocr.py
import pytesseract
from PIL import Image
import sys

image_path = sys.argv[1]
image = Image.open(image_path)

# Basic OCR
text = pytesseract.image_to_string(image)
print(text)

# With language
text = pytesseract.image_to_string(image, lang='por')

# Get bounding boxes
data = pytesseract.image_to_data(image, output_type=pytesseract.Output.DICT)
for i, word in enumerate(data['text']):
    if word.strip():
        print(f"{word}: ({data['left'][i]}, {data['top'][i]})")

# Get structured data
print(pytesseract.image_to_string(image, config='--psm 6'))
```

```bash
pip install pytesseract pillow
python ocr.py image.png
```

## Option 4: OCR APIs

```bash
# Google Cloud Vision
curl -s -X POST "https://vision.googleapis.com/v1/images:annotate?key=$GOOGLE_VISION_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "requests": [{
      "image": {"content": "'$(base64 -w0 image.png)'"},
      "features": [{"type": "TEXT_DETECTION"}]
    }]
  }' | jq '.responses[0].fullTextAnnotation.text'

# AWS Textract (requires AWS CLI)
aws textract detect-document-text \
  --document '{"S3Object":{"Bucket":"bucket","Name":"image.png"}}'

# OCR.space (free tier)
curl -s -X POST "https://api.ocr.space/parse/image" \
  -H "apikey:$OCR_SPACE_KEY" \
  -F "file=@image.png" | jq '.ParsedResults[0].ParsedText'
```

## PDF OCR

```bash
# Convert PDF to images first
pdftoppm input.pdf page -png

# OCR each page
for f in page-*.png; do
  tesseract "$f" "${f%.png}" -l eng+por
done

# Combine results
cat page-*.txt > full_text.txt

# Or use OCRmyPDF (adds text layer to PDF)
ocrmypdf input.pdf searchable.pdf
```

## Tesseract Page Segmentation Modes (PSM)

```bash
# PSM options:
# 0 = Orientation and script detection
# 3 = Fully automatic (default)
# 6 = Single uniform block of text
# 11 = Sparse text with specific ordering
# 12 = Sparse text with OSD

tesseract image.png output --psm 6
```

## Tips

- Preprocess images for better results (contrast, resize, denoise)
- Use appropriate PSM mode for your content type
- Install language packs: `tesseract-ocr-por` for Portuguese
- 300 DPI minimum for good OCR results
- Use `--oem 1` for LSTM neural net engine

## Triggers

ocr, text recognition, extract text, image to text,
read image, scan text, ocr image
