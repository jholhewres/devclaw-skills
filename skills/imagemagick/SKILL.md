---
name: imagemagick
description: "Image processing with ImageMagick (convert, magick)"
---
# ImageMagick

Use the **bash** tool with magick (v7) or convert (v6) for image processing.

## Info
```bash
magick identify -verbose <image>
magick identify -format "%wx%h\n" <image>
```

## Resize
```bash
magick <input> -resize 800x600 <output>
magick <input> -resize 50% <output>
magick <input> -resize 800x600^ -gravity center -extent 800x600 <output>  # crop to fit
```

## Convert Format
```bash
magick input.png output.jpg
magick input.svg -density 300 output.png
magick input.webp output.png
```

## Optimize
```bash
magick input.jpg -quality 85 -strip output.jpg
magick input.png -strip -colors 256 output.png
```

## Batch
```bash
magick mogrify -resize 800x600 -quality 85 *.jpg
```

## Compose
```bash
magick base.png overlay.png -gravity southeast -composite output.png
magick -size 800x400 xc:white -font Helvetica -pointsize 48 -gravity center -annotate 0 "Text" output.png
```

## Tips
- Use magick (v7) instead of convert (v6) when available
- Use -strip to remove metadata (reduces file size)
- Use mogrify for in-place batch operations
