---
name: media-config
description: "Configure audio transcription and image/video understanding for channels"
---
# Media Configuration

Configure how DevClaw processes images, videos, and audio received from messaging channels (WhatsApp, Telegram, Discord, WebUI).

Settings live in `config.yaml` under the `media:` section and can also be changed via WebUI → Configuração.

## Vision (Image/Video Understanding)

Controls how images and video frames are described before being added to conversation context.

```yaml
media:
  vision_enabled: true
  vision_model: ""          # empty = use main chat model
  vision_detail: "auto"     # auto | low | high
  max_image_size: 20971520  # 20MB
```

### Available Vision Models

| Provider  | Model                     | Notes                                |
|-----------|---------------------------|--------------------------------------|
| Z.AI      | glm-4.6v                  | Flagship, 128K, native tool use      |
| Z.AI      | glm-4.6v-flashx           | Lightweight, affordable              |
| Z.AI      | glm-4.6v-flash            | Free tier                            |
| Z.AI      | glm-4.5v                  | 106B MOE, thinking mode              |
| OpenAI    | gpt-4o                    | Best quality                         |
| OpenAI    | gpt-4o-mini               | Fast, cheap                          |
| OpenAI    | gpt-4.1                   | Latest                               |
| OpenAI    | gpt-4.1-mini              | Balanced                             |
| Anthropic | claude-sonnet-4-20250514  | Claude Sonnet 4                      |
| Anthropic | claude-opus-4-20250514    | Claude Opus 4                        |
| Anthropic | claude-haiku-3-5-20241022 | Fast, cheap                          |
| Google    | gemini-3-pro              | Multimodal flagship                  |
| Google    | gemini-3-flash            | Fast multimodal                      |
| Google    | gemini-2.5-pro            | Strong reasoning                     |

If `vision_model` is empty, the main `model` from config is used. Set a dedicated vision model when:
- The main model doesn't support images (e.g. text-only)
- You want a cheaper model for image understanding
- You want the best vision quality regardless of chat model

### Vision Detail

- **auto**: Let the API decide based on image size
- **low**: Faster, fewer tokens (~85 tokens/image)
- **high**: Detailed analysis, more tokens (~1590 tokens/1092×1092 image)

## Audio Transcription

Controls how voice messages and audio files are converted to text.

```yaml
media:
  transcription_enabled: true
  transcription_model: "whisper-1"
  transcription_base_url: "https://api.openai.com/v1"
  transcription_api_key: ""    # empty = use main API key
  max_audio_size: 26214400     # 25MB
```

### Available Transcription Models

| Provider | Model                   | Base URL                          | Notes                              |
|----------|-------------------------|-----------------------------------|------------------------------------|
| Z.AI     | glm-asr-2512            | https://api.z.ai/api/paas/v4     | Multilingual, CER 0.07, max 25MB  |
| OpenAI   | whisper-1               | https://api.openai.com/v1        | Legacy, widely compatible          |
| OpenAI   | gpt-4o-transcribe       | https://api.openai.com/v1        | Best quality, logprobs support     |
| OpenAI   | gpt-4o-mini-transcribe  | https://api.openai.com/v1        | Lighter, fast                      |
| Groq     | whisper-large-v3        | https://api.groq.com/openai/v1   | 189x realtime speed, $0.11/hr     |
| Groq     | whisper-large-v3-turbo  | https://api.groq.com/openai/v1   | 216x speed, $0.04/hr              |

### Choosing a Transcription Provider

- **Z.AI GLM-ASR-2512**: Best if already using Z.AI as main provider. Low CER, supports Chinese/English/dialects.
- **OpenAI GPT-4o Transcribe**: Best quality, supports diarization variant.
- **Groq Whisper**: Fastest (189–216x realtime), cheapest, OpenAI-compatible endpoint.
- **OpenAI Whisper-1**: Reliable fallback, broadest format support (SRT, VTT, verbose JSON).

### Using a Different Transcription Provider

When the main LLM provider doesn't support audio transcription (e.g. Anthropic, xAI), set `transcription_base_url` and `transcription_api_key`:

```yaml
# Main provider is Anthropic, transcription via Groq
api:
  provider: anthropic
  api_key: ${DEVCLAW_API_KEY}

media:
  transcription_enabled: true
  transcription_model: whisper-large-v3
  transcription_base_url: https://api.groq.com/openai/v1
  transcription_api_key: ${DEVCLAW_GROQ_API_KEY}
```

Store the separate API key in the vault:
```
vault_save groq_api_key gsk_xxxx
```

## Quick Setup Examples

### Z.AI Full Stack (Vision + Audio)
```yaml
media:
  vision_enabled: true
  vision_model: glm-4.6v
  vision_detail: auto
  transcription_enabled: true
  transcription_model: glm-asr-2512
  transcription_base_url: https://api.z.ai/api/paas/v4
```

### OpenAI Full Stack
```yaml
media:
  vision_enabled: true
  vision_model: gpt-4o
  vision_detail: auto
  transcription_enabled: true
  transcription_model: gpt-4o-transcribe
  transcription_base_url: https://api.openai.com/v1
```

### Budget Setup (Cheap Vision + Fast Audio)
```yaml
media:
  vision_enabled: true
  vision_model: gpt-4o-mini
  vision_detail: low
  transcription_enabled: true
  transcription_model: whisper-large-v3-turbo
  transcription_base_url: https://api.groq.com/openai/v1
  transcription_api_key: ${DEVCLAW_GROQ_API_KEY}
```
