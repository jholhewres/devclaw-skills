---
name: video-tools
version: 0.1.0
author: goclaw
description: "Video manipulation — convert, compress, trim, and edit videos with ffmpeg"
category: media
tags: [video, ffmpeg, convert, compress, trim, audio]
requires:
  bins: [ffmpeg, ffprobe]
---
# Video Tools

Manipulate videos using ffmpeg.

## Setup

```bash
# Ubuntu/Debian
sudo apt install ffmpeg

# macOS
brew install ffmpeg

# Check installation
ffmpeg -version
```

## Get Video Info

```bash
# Basic info
ffprobe -i video.mp4

# Duration only
ffprobe -i video.mp4 -show_entries format=duration -v quiet -of csv="p=0"

# Resolution
ffprobe -i video.mp4 -show_entries stream=width,height -v quiet -of csv="p=0"

# JSON output
ffprobe -i video.mp4 -print_format json -show_format -show_streams
```

## Convert Formats

```bash
# MP4 to WebM
ffmpeg -i input.mp4 -c:v libvpx -c:a libvorbis output.webm

# MP4 to MOV
ffmpeg -i input.mp4 -c:v prores -c:a pcm_s16le output.mov

# Any to MP4 (H.264)
ffmpeg -i input.avi -c:v libx264 -c:a aac output.mp4

# MKV to MP4
ffmpeg -i input.mkv -c:v copy -c:a copy output.mp4

# GIF to MP4
ffmpeg -i input.gif -c:v libx264 -pix_fmt yuv420p output.mp4
```

## Compress Video

```bash
# Reduce file size (CRF: 0-51, higher = smaller)
ffmpeg -i input.mp4 -c:v libx264 -crf 28 -c:a aac compressed.mp4

# CRF values: 18 (high quality) to 28 (smaller size)
ffmpeg -i input.mp4 -c:v libx264 -crf 23 -preset slow compressed.mp4

# Preset options: ultrafast, superfast, veryfast, faster, fast, medium, slow, slower, veryslow

# Scale down for compression
ffmpeg -i input.mp4 -vf "scale=iw/2:ih/2" -c:v libx264 -crf 23 smaller.mp4

# Limit bitrate
ffmpeg -i input.mp4 -b:v 1M -b:a 128k limited.mp4
```

## Resize Video

```bash
# Scale to width (keep aspect ratio)
ffmpeg -i input.mp4 -vf "scale=1280:-1" output.mp4

# Scale to height
ffmpeg -i input.mp4 -vf "scale=-1:720" output.mp4

# Exact dimensions
ffmpeg -i input.mp4 -vf "scale=1280x720" output.mp4

# Half size
ffmpeg -i input.mp4 -vf "scale=iw/2:ih/2" output.mp4
```

## Trim Video

```bash
# Trim from 10s to 30s
ffmpeg -i input.mp4 -ss 00:00:10 -to 00:00:30 -c copy output.mp4

# Trim from 10s for 20 seconds duration
ffmpeg -i input.mp4 -ss 00:00:10 -t 00:00:20 -c copy output.mp4

# Fast trim (may be less accurate)
ffmpeg -ss 10 -i input.mp4 -t 20 -c copy output.mp4
```

## Extract Audio

```bash
# Extract audio to MP3
ffmpeg -i input.mp4 -vn -c:a libmp3lame -q:a 2 audio.mp3

# Extract audio to AAC
ffmpeg -i input.mp4 -vn -c:a aac -b:a 192k audio.aac

# Extract audio without re-encoding
ffmpeg -i input.mp4 -vn -c:a copy audio.aac
```

## Add Audio to Video

```bash
# Add audio track
ffmpeg -i video.mp4 -i audio.mp3 -c:v copy -c:a aac -map 0:v:0 -map 1:a:0 output.mp4

# Replace audio
ffmpeg -i video.mp4 -i audio.mp3 -c:v copy -c:a aac -map 0:v:0 -map 1:a:0 -shortest output.mp4
```

## Create GIF

```bash
# Video to GIF
ffmpeg -i input.mp4 -vf "fps=10,scale=320:-1" output.gif

# High quality GIF with palette
ffmpeg -i input.mp4 -vf "fps=15,scale=480:-1:flags=lanczos,split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse" output.gif

# Trim and convert to GIF
ffmpeg -ss 5 -t 3 -i input.mp4 -vf "fps=10,scale=320:-1" clip.gif
```

## Combine Videos

```bash
# Concatenate videos
ffmpeg -f concat -safe 0 -i filelist.txt -c copy output.mp4

# filelist.txt content:
# file 'video1.mp4'
# file 'video2.mp4'
# file 'video3.mp4'

# Or inline
ffmpeg -i "concat:video1.mp4|video2.mp4" -c copy output.mp4
```

## Add Watermark

```bash
# Add image watermark (bottom right)
ffmpeg -i input.mp4 -i watermark.png -filter_complex "overlay=W-w-10:H-h-10" output.mp4

# Add text watermark
ffmpeg -i input.mp4 -vf "drawtext=text='Watermark':fontcolor=white:fontsize=24:x=W-w-10:y=H-h-10" output.mp4
```

## Tips

- Use `-c copy` to avoid re-encoding (much faster)
- Use `-ss` before `-i` for fast seeking
- CRF 23 is good default for quality/size balance
- Use `-preset slow` for better compression
- Check codec support for target platform

## Triggers

video, ffmpeg, convert video, compress video, trim video,
video editing, extract audio, video to gif
