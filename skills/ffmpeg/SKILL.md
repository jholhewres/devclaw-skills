---
name: ffmpeg
description: "Video and audio processing with ffmpeg"
---
# FFmpeg

Use the **bash** tool with ffmpeg for media processing.

## Info
```bash
ffprobe -v quiet -print_format json -show_format -show_streams <input>
```

## Convert
```bash
ffmpeg -i input.mp4 -c:v libx264 -crf 23 -preset medium output.mp4
ffmpeg -i input.mp4 -vn -acodec libmp3lame -q:a 2 output.mp3
ffmpeg -i input.webm -c:v libx264 output.mp4
```

## Trim & Cut
```bash
ffmpeg -i input.mp4 -ss 00:01:00 -t 00:00:30 -c copy output.mp4
```

## Resize
```bash
ffmpeg -i input.mp4 -vf "scale=1280:720" output.mp4
ffmpeg -i input.mp4 -vf "scale=-1:720" output.mp4   # maintain aspect ratio
```

## Audio
```bash
ffmpeg -i input.mp4 -vn -acodec copy audio.aac
ffmpeg -i audio.wav -acodec libmp3lame -b:a 192k output.mp3
```

## GIF
```bash
ffmpeg -i input.mp4 -vf "fps=10,scale=480:-1:flags=lanczos" -c:v gif output.gif
```

## Tips
- Use -c copy for fast operations without re-encoding
- Use -crf for quality (lower = better, 18-28 is good range)
- Always check output with ffprobe
