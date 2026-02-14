---
name: pdf-tools
version: 0.1.0
author: goclaw
description: "PDF manipulation — merge, split, compress, convert, and modify PDFs"
category: documents
tags: [pdf, merge, split, compress, convert, document]
requires:
  bins: [qpdf, pdftk, ghostscript]
---
# PDF Tools

Manipulate PDF files using command-line tools.

## Setup

```bash
# Ubuntu/Debian
sudo apt install qpdf pdftk ghostscript poppler-utils

# macOS
brew install qpdf pdftk ghostscript poppler
```

## Merge PDFs

```bash
# Merge multiple PDFs
qpdf --empty --keep-temporary-files page1.pdf page2.pdf page3.pdf -- merged.pdf

# With pdftk
pdftk page1.pdf page2.pdf page3.pdf cat output merged.pdf

# Merge all PDFs in directory
pdftk *.pdf cat output all_merged.pdf
```

## Split PDFs

```bash
# Extract pages 1-5
qpdf input.pdf --pages . 1-5 -- output.pdf

# Extract single page
pdftk input.pdf cat 3 output page3.pdf

# Split into individual pages
pdftk input.pdf burst output page_%d.pdf

# Extract range
qpdf input.pdf --pages . 5-10 -- extracted.pdf
```

## Compress PDF

```bash
# Compress with Ghostscript (best quality)
gs -sDEVICE=pdfwrite -dCompatibilityLevel=1.4 -dPDFSETTINGS=/ebook \
   -dNOPAUSE -dQUIET -dBATCH -sOutputFile=compressed.pdf input.pdf

# PDF settings options:
# /screen   - 72 dpi, smallest size, lowest quality
# /ebook    - 150 dpi, medium size, good quality
# /printer  - 300 dpi, larger size, high quality
# /prepress - 300 dpi, highest quality

# Quick compress with qpdf
qpdf --compress-streams=y --recompress-flate input.pdf compressed.pdf
```

## Rotate PDF

```bash
# Rotate all pages 90 degrees clockwise
pdftk input.pdf cat 1-endright output rotated.pdf

# Rotate specific pages
qpdf input.pdf --rotate=+90:1,3,5 output.pdf

# Rotation options: +90, +180, +270, -90
```

## Add Watermark

```bash
# Add text watermark with pdftk
pdftk input.pdf background watermark.pdf output watermarked.pdf

# Create watermark PDF first
echo "DRAFT" | enscript -B -f Courier-Bold100 -o - | ps2pdf - watermark.pdf
```

## Password Protect

```bash
# Add password (owner + user)
qpdf --encrypt user_pass owner_pass 256 -- input.pdf protected.pdf

# With pdftk
pdftk input.pdf output protected.pdf encrypt_128bit owner_pw owner_pass user_pw user_pass
```

## Remove Password

```bash
# Remove password (requires password)
qpdf --password=yourpass --decrypt protected.pdf unlocked.pdf

# With pdftk
pdftk protected.pdf input_pw yourpass output unlocked.pdf
```

## PDF to Images

```bash
# Convert PDF to images (each page)
pdftoppm input.pdf output -png

# Specific page
pdftoppm -f 1 -l 1 input.pdf output -png

# Set resolution (DPI)
pdftoppm -rx 300 -ry 300 input.pdf output -png
```

## Images to PDF

```bash
# Convert images to PDF
convert image1.png image2.png output.pdf

# Or with img2pdf (better quality)
img2pdf image1.png image2.png -o output.pdf
```

## Extract Text

```bash
# Extract text from PDF
pdftotext input.pdf output.txt

# Keep layout
pdftotext -layout input.pdf output.txt

# Extract from specific pages
pdftotext -f 1 -l 5 input.pdf output.txt
```

## Remove Metadata

```bash
# Strip all metadata
qpdf --linearize input.pdf clean.pdf
exiftool -all:all= clean.pdf

# Or with Ghostscript
gs -sDEVICE=pdfwrite -dNOPAUSE -dQUIET -dBATCH \
   -sOutputFile=clean.pdf input.pdf
```

## Tips

- Use `qpdf` for most operations (faster, more reliable)
- Use `pdftk` for complex page manipulations
- Use Ghostscript for compression
- Always backup before modifying important PDFs

## Triggers

pdf, merge pdf, split pdf, compress pdf, pdf tools, pdf merge,
password pdf, watermark pdf
