---
name: audio-tts
version: 0.1.0
author: devclaw
description: "Text-to-Speech and audio tools — convert text to speech, manipulate audio"
category: media
tags: [audio, tts, speech, voice, text-to-speech]
requires:
  bins: [curl, jq]
---
# Audio TTS

Convert text to speech and manipulate audio files.

## Option 1: OpenAI TTS (Recommended)

```bash
# Setup
export OPENAI_API_KEY="sk-xxx"

# Text to speech
curl -s -X POST "https://api.openai.com/v1/audio/speech" \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "tts-1",
    "input": "Hello, this is a text to speech example.",
    "voice": "alloy"
  }' -o output.mp3

# Voices: alloy, echo, fable, onyx, nova, shimmer
# Models: tts-1 (faster), tts-1-hd (higher quality)
```

## Option 2: ElevenLabs TTS

```bash
# Setup
export ELEVENLABS_API_KEY="xxx"

# Text to speech
curl -s -X POST "https://api.elevenlabs.io/v1/text-to-speech/VOICE_ID" \
  -H "xi-api-key: $ELEVENLABS_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "text": "Hello from DevClaw!",
    "model_id": "eleven_monolingual_v1",
    "voice_settings": {"stability": 0.5, "similarity_boost": 0.5}
  }' -o output.mp3

# List voices
curl -s "https://api.elevenlabs.io/v1/voices" \
  -H "xi-api-key: $ELEVENLABS_API_KEY" | jq '.voices[]'
```

## Option 3: Google Cloud TTS

```bash
# Setup
export GOOGLE_TTS_KEY="xxx"

# Synthesize speech
curl -s -X POST "https://texttospeech.googleapis.com/v1/text:synthesize?key=$GOOGLE_TTS_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "input": {"text": "Hello world"},
    "voice": {"languageCode": "en-US", "name": "en-US-Standard-A"},
    "audioConfig": {"audioEncoding": "MP3"}
  }' | jq -r '.audioContent' | base64 -d > output.mp3
```

## Option 4: Amazon Polly

```bash
# Requires AWS CLI setup
aws polly synthesize-speech \
  --output-format mp3 \
  --voice-id Joanna \
  --text "Hello from DevClaw" \
  output.mp3
```

## Option 5: Free TTS (say command)

```bash
# macOS built-in
say "Hello world" -o output.aiff

# Convert to MP3
ffmpeg -i output.aiff output.mp3

# Linux (espeak)
espeak "Hello world" -w output.wav
```

## Audio Manipulation (ffmpeg)

```bash
# Change speed
ffmpeg -i input.mp3 -filter:a "atempo=1.5" output.mp3

# Change volume
ffmpeg -i input.mp3 -filter:a "volume=2.0" output.mp3

# Trim audio
ffmpeg -i input.mp3 -ss 00:00:10 -t 30 output.mp3

# Convert format
ffmpeg -i input.wav -c:a libmp3lame -q:a 2 output.mp3

# Merge audio files
ffmpeg -i "concat:file1.mp3|file2.mp3" -c copy output.mp3

# Extract audio from video
ffmpeg -i video.mp4 -vn -c:a libmp3lame audio.mp3
```

## Speech-to-Text (Transcription)

```bash
# OpenAI Whisper
curl -s -X POST "https://api.openai.com/v1/audio/transcriptions" \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -F "file=@audio.mp3" \
  -F "model=whisper-1" | jq '.text'

# With timestamp
curl -s -X POST "https://api.openai.com/v1/audio/transcriptions" \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -F "file=@audio.mp3" \
  -F "model=whisper-1" \
  -F "response_format=verbose_json" | jq '.segments'
```

## Tips

- OpenAI TTS: Good quality, simple API
- ElevenLabs: Best quality, many voice options
- Use `tts-1-hd` for production quality
- Audio formats: MP3, WAV, OGG, AAC supported
- For long texts, split into chunks

## Triggers

tts, text to speech, speech synthesis, voice, audio,
convert text to speech, read text aloud
