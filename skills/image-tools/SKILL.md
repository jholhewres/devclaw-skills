---
name: image-tools
version: 0.1.0
author: devclaw
description: "Image manipulation — resize, convert, compress, and edit images"
category: media
tags: [image, resize, convert, compress, imagemagick]
requires:
  bins: [convert, identify]
---
# Image Tools

Manipulate images using ImageMagick and other tools.

## Setup

```bash
# Ubuntu/Debian
sudo apt install imagemagick

# macOS
brew install imagemagick

# Check installation
convert --version
```

## Resize Images

```bash
# Resize to specific dimensions
convert input.png -resize 800x600 output.png

# Resize maintaining aspect ratio
convert input.png -resize 800x output.png
convert input.png -resize x600 output.png

# Resize percentage
convert input.png -resize 50% output.png

# Force exact size (may distort)
convert input.png -resize 800x600! output.png

# Fill area and crop excess
convert input.png -resize 800x600^ -gravity center -extent 800x600 output.png

# Fit within area (no crop)
convert input.png -resize 800x600\> output.png
```

## Convert Formats

```bash
# PNG to JPG
convert input.png output.jpg

# JPG to PNG
convert input.jpg output.png

# To WebP
convert input.png -quality 85 output.webp

# To GIF
convert input.png output.gif

# Multiple images to PDF
convert *.png output.pdf

# PDF to images
convert input.pdf output.png

# Batch convert
for f in *.png; do convert "$f" "${f%.png}.jpg"; done
```

## Compress Images

```bash
# Compress JPEG (quality 1-100)
convert input.jpg -quality 85 compressed.jpg

# Strip metadata and compress
convert input.jpg -strip -quality 85 compressed.jpg

# Compress PNG
convert input.png -strip -quality 85 compressed.png

# Optimize for web
convert input.png -strip -interlace Plane -quality 85 web.jpg
```

## Crop Images

```bash
# Crop to dimensions (width x height + x_offset + y_offset)
convert input.png -crop 400x300+100+50 output.png

# Crop from center
convert input.png -gravity center -crop 400x300+0+0 output.png

# Remove borders
convert input.png -trim output.png

# Crop percentage
convert input.png -crop 50%x50% output.png
```

## Rotate & Flip

```bash
# Rotate 90 degrees
convert input.png -rotate 90 output.png

# Rotate 180
convert input.png -rotate 180 output.png

# Flip vertical
convert input.png -flip output.png

# Flip horizontal (mirror)
convert input.png -flop output.png

# Auto-rotate based on EXIF
convert input.jpg -auto-orient output.jpg
```

## Add Effects

```bash
# Blur
convert input.png -blur 0x8 output.png

# Sharpen
convert input.png -sharpen 0x2 output.png

# Grayscale
convert input.png -colorspace Gray output.png

# Sepia
convert input.png -sepia-tone 80% output.png

# Add border
convert input.png -bordercolor white -border 10 output.png

# Rounded corners
convert input.png -matte -fill none -draw "roundrectangle 0,0,w,h,15,15" rounded.png
```

## Add Watermark

```bash
# Text watermark
convert input.png -pointsize 36 -fill white -gravity southeast \
  -annotate +10+10 "Watermark" watermarked.png

# Image watermark
composite -dissolve 50% -gravity southeast watermark.png input.png output.png

# Tiled watermark
convert input.png -fill white -font Arial -pointsize 30 \
  -draw "gravity center rotate -45 text 0,0 'WATERMARK'" tiled.png
```

## Get Image Info

```bash
# Basic info
identify input.png

# Detailed info
identify -verbose input.png

# JSON output
identify -format '{"width": %w, "height": %h, "format": "%m", "size": "%b"}\n' input.png

# Get dimensions only
identify -format "%wx%h" input.png
```

## Batch Processing

```bash
# Resize all images
mogrify -resize 800x *.png

# Convert all to JPG
mogrify -format jpg *.png

# Compress all
mogrify -strip -quality 85 *.jpg

# ⚠️ Warning: mogrify modifies files in place!
```

## Tips

- Use `-strip` to remove metadata and reduce size
- Use `-quality 85` for good JPEG compression
- Use `mogrify` for batch processing (modifies in place!)
- Use `-interlace Plane` for progressive loading
- Check dimensions with `identify` before processing

## Triggers

image resize, resize image, convert image, compress image,
image tools, imagemagick, image manipulation
