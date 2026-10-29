# ebook2audiobook (E2A)

[![Docker Build](https://github.com/DrewThomasson/ebook2audiobook/actions/workflows/Docker-Build.yml/badge.svg)](https://github.com/DrewThomasson/ebook2audiobook/actions/workflows/Docker-Build.yml)

Convert eBooks to audiobooks with chapters and metadata. Runs on CPU or GPU. Supports voice cloning and 1,158 languages.

> [!IMPORTANT]
> **This tool is intended for use with non-DRM, legally acquired eBooks only.**
> The authors are not responsible for any misuse of this software or any resulting legal consequences.
> Use this tool responsibly and in accordance with all applicable laws.

[![Discord](https://dcbadge.limes.pink/api/server/https://discord.gg/63Tv3F65k6)](https://discord.gg/63Tv3F65k6)

## Features

- **TTS engines:** XTTSv2, Bark, Fairseq, VITS, Tacotron2, Tortoise, GlowTTS, YourTTS
- **Voice cloning** from any audio file (1 to 5 minutes, noise is cleaned automatically)
- **1,158 languages** ([full list](https://dl.fbaipublicfiles.com/mms/tts/all-tts-languages.html))
- **OCR scanning** for image-based pages (PDF, JPG, BMP, PNG, TIFF)
- **SML tags** for pauses, breaks, and mid-text voice switching
- **Custom models** (XTTSv2, VITS, Fairseq, Piper, others on request)
- **Fine-tuned preset models** trained by the E2A team

## Requirements

| Resource | Minimum | Recommended |
| -------- | ------- | ----------- |
| RAM | 2 GB | 8 GB |
| VRAM | 1 GB | 4 GB |

Supported hardware: CPU, CUDA, ROCm, XPU (Intel, AMD, ARM), MPS (Apple Silicon), Jetson.

Modern TTS engines run slowly on CPU. Use lighter engines (YourTTS, Tacotron2) for CPU-only systems.

Windows Docker users must enable virtualization.

## Supported formats

**Input.** `.epub`, `.pdf`, `.mobi`, `.txt`, `.html`, `.rtf`, `.chm`, `.lit`, `.pdb`, `.fb2`, `.odt`, `.cbr`, `.cbz`, `.prc`, `.lrf`, `.pml`, `.snb`, `.cbc`, `.rb`, `.tcr`. Also supports `.doc`, `.docx`, `.azw`, `.azw3`, `.tiff`, `.tif`, `.png`, `.jpg`, `.jpeg`, `.bmp`, `.zip`.

`.epub` and `.mobi` produce the best chapter detection.

**Output.** `.m4b`, `.m4a`, `.mp4`, `.webm`, `.mov`, `.mp3`, `.flac`, `.wav`, `.ogg`, `.aac` (mono or stereo)

Change the default output format in `lib/conf.py`.

## Supported languages

| Arabic (ar) | Chinese (zh) | English (en) | Spanish (es) |
|:-----------:|:------------:|:------------:|:------------:|
| French (fr) | German (de) | Italian (it) | Portuguese (pt) |
| Polish (pl) | Turkish (tr) | Russian (ru) | Dutch (nl) |
| Czech (cs) | Japanese (ja) | Hindi (hi) | Bengali (bn) |
| Hungarian (hu) | Korean (ko) | Vietnamese (vi) | Swedish (sv) |
| Persian (fa) | Yoruba (yo) | Swahili (sw) | Indonesian (id) |
| Slovak (sk) | Croatian (hr) | Tamil (ta) | Danish (da) |

[+1,130 languages and dialects](https://dl.fbaipublicfiles.com/mms/tts/all-tts-languages.html)

## Install

### Clone and run

```bash
git clone https://github.com/DrewThomasson/ebook2audiobook.git
cd ebook2audiobook
```

**Linux/macOS:**

```bash
./ebook2audiobook.command
```

macOS installs missing dependencies through Homebrew.

**Mac launcher.** Double-click `Mac Ebook2Audiobook Launcher.command`.

**Windows:**

```bash
ebook2audiobook.cmd
```

Or double-click `ebook2audiobook.cmd`. Windows installs missing dependencies through Scoop (no admin required).

### Open the web interface

The terminal prints a URL. Open `http://localhost:7860/` in your browser.

For a public link:

```bash
./ebook2audiobook.command --share    # Linux/macOS
ebook2audiobook.cmd --share          # Windows
python app.py --share                # any OS
```

> If you stop and restart the script, refresh the web interface to reconnect.

## Headless mode

### Basic usage

**Linux/macOS:**

```bash
./ebook2audiobook.command --headless --ebook <path_to_ebook> --voice <path_to_voice> --language <language_code>
```

**Windows:**

```bash
ebook2audiobook.cmd --headless --ebook <path_to_ebook> --voice <path_to_voice> --language <language_code>
```

- `--ebook`: Path to the eBook file
- `--voice`: Path to voice cloning audio (optional)
- `--language`: ISO-639-3 code (for example, `ita` for Italian, `eng` for English). ISO-639-1 two-letter codes also work. Defaults to `eng`.

### Custom model

Provide a `.zip` file with the required model files. For XTTSv2: `config.json`, `model.pth`, `vocab.json`, and `ref.wav`.

```bash
./ebook2audiobook.command --headless --ebook <path> --language <lang> --custom_model <model.zip>
```

The `ref.wav` inside the zip is always used as the voice for conversion. See `lib/models.py` for required files per engine.

### All parameters

```bash
./ebook2audiobook.command --help    # Linux/macOS
ebook2audiobook.cmd --help          # Windows
python app.py --help                # any OS
```

<details>
<summary>Full help output</summary>

```
usage: app.py [-h] [--session SESSION] [--share] [--headless] [--ebook EBOOK] [--ebooks_dir EBOOKS_DIR]
              [--language LANGUAGE] [--voice VOICE] [--voice_map VOICE_MAP] [--device {CPU,CUDA,MPS,ROCM,XPU,JETSON}]
              [--tts_engine {XTTS,BARK,VITS,FAIRSEQ,TACOTRON,YOURTTS,...}]
              [--custom_model CUSTOM_MODEL] [--fine_tuned FINE_TUNED] [--output_format OUTPUT_FORMAT]
              [--output_channel OUTPUT_CHANNEL] [--temperature TEMPERATURE] [--length_penalty LENGTH_PENALTY]
              [--num_beams NUM_BEAMS] [--repetition_penalty REPETITION_PENALTY] [--top_k TOP_K] [--top_p TOP_P]
              [--speed SPEED] [--enable_text_splitting] [--text_temp TEXT_TEMP] [--waveform_temp WAVEFORM_TEMP]
              [--output_dir OUTPUT_DIR] [--version]

Convert eBooks to audiobooks using a text-to-speech model.

options:
  -h, --help                Show help and exit
  --session SESSION         Resume a previous session (reuses custom models and voices)
  --share                   Enable a public Gradio link (GUI mode only)
  --headless                Run without the GUI
  --ebook EBOOK             Path to a single eBook file
  --ebooks_dir EBOOKS_DIR   Path to a directory of eBooks (cannot combine with --ebook)
  --text TEXT                Raw text for conversion (cannot combine with --ebook or --ebooks_dir)
  --language LANGUAGE        Language code (ISO-639-3). Default set in ./lib/lang.py
  --translate ISO3           Translate the eBook before synthesis. Uses argostranslate.
  --voice VOICE              Voice cloning audio file. Uses default voice if omitted.
  --voice_map VOICE_MAP      JSON file mapping eBook paths to voice files (--ebooks_dir only)
  --device DEVICE            Processor: CPU, CUDA, MPS, ROCM, XPU, or JETSON
  --tts_engine ENGINE        TTS engine: XTTS, BARK, VITS, FAIRSEQ, TACOTRON, or YOURTTS
  --custom_model PATH        Path to a custom model .zip
  --fine_tuned PATH          Path to a fine-tuned model
  --output_format FORMAT     Output audio format (default: m4b)
  --output_channel CHANNEL   Mono or stereo (default: mono)
  --temperature FLOAT        Model temperature (XTTSv2 only)
  --length_penalty FLOAT     Autoregressive decoder length penalty (XTTSv2 only)
  --num_beams INT            Beam search width (XTTSv2 only)
  --repetition_penalty FLOAT Repetition penalty (XTTSv2 only)
  --top_k INT                Top-k sampling (XTTSv2 only)
  --top_p FLOAT              Top-p sampling (XTTSv2 only)
  --speed FLOAT              Speech speed factor (XTTSv2 only)
  --enable_text_splitting    Enable TTS text splitting (XTTSv2 only)
  --text_temp FLOAT          Text temperature (Bark only)
  --waveform_temp FLOAT      Waveform temperature (Bark only)
  --output_dir PATH          Output directory
  --abs_url URL              Audiobookshelf server URL
  --abs_api_token TOKEN      Audiobookshelf API token
  --abs_library ID           Audiobookshelf library ID
  --enable_interlude         Add MusicGen music between chapters (settings in ./lib/conf_interlude.py)
  --version                  Show version and exit
```

</details>

## SML tags

Control pauses and voice switching inline in your text:

- `[break]`: silence (0.3 to 0.6 seconds, random)
- `[pause]`: silence (1.0 to 1.6 seconds, random)
- `[pause:N]`: fixed pause of N seconds
- `[voice:/path/to/file]...[/voice]`: switch voice for the enclosed text

See [E2A-SML](https://github.com/DrewThomasson/E2A-SML) for automatic SML tag insertion.

## Docker

### Build

```bash
# Linux/macOS
./ebook2audiobook.command --script_mode build_docker

# Windows
ebook2audiobook.cmd --script_mode build_docker

# Docker Compose
./ebook2audiobook.command --script_mode build_docker --docker_mode compose

# Podman Compose
./ebook2audiobook.command --script_mode build_docker --docker_mode podman
```

### Run

All `docker run` commands mount five volumes for persistent storage:

```bash
-v "./ebooks:/app/ebooks" \
-v "./audiobooks:/app/audiobooks" \
-v "./models:/app/models" \
-v "./voices:/app/voices" \
-v "./tmp:/app/tmp"
```

**CPU (GUI):**

```bash
docker run -v ... --rm -it -p 7860:7860 athomasson2/ebook2audiobook:cpu
```

**CUDA (GUI):**

```bash
docker run -v ... --gpus all --rm -it -p 7860:7860 athomasson2/ebook2audiobook:cu128
```

Replace `cu128` with your CUDA version tag (cu118, cu122, cu124, cu126).

**ROCm (GUI):**

```bash
docker run -v ... --device=/dev/kfd --device=/dev/dri --rm -it -p 7860:7860 athomasson2/ebook2audiobook:rocm6.4
```

**XPU (GUI):**

```bash
docker run -v ... --device=/dev/dri --rm -it -p 7860:7860 athomasson2/ebook2audiobook:xpu
```

**Jetson (GUI):**

```bash
docker run -v ... --runtime nvidia --rm -it -p 7860:7860 athomasson2/ebook2audiobook:jetson60
```

For headless mode, add `--headless --ebook "/app/ebooks/myfile.pdf"` and any other flags after the image name.

For an eBook outside the mounted volumes, add another `-v` mount:

```bash
docker run -v ... -v "/path/to/ebooks:/app/another_folder" --rm -it -p 7860:7860 athomasson2/ebook2audiobook:cpu --headless --ebook "/app/another_folder/myfile.pdf"
```

**Docker Compose (CUDA 12.8 example):**

```bash
# GUI
DEVICE_TAG=cu128 docker compose --profile gpu up --no-log-prefix

# Headless
DEVICE_TAG=cu128 docker compose --profile gpu run --rm ebook2audiobook --headless --ebook "/app/ebooks/myfile.pdf"
```

**Podman Compose:**

```bash
DEVICE_TAG=cu128 podman-compose -f podman-compose.yml --profile gpu up
```

MPS is not available in Docker. Use CPU instead.

## Run remotely

- [Hugging Face Space](https://huggingface.co/spaces/drewThomasson/ebook2audiobook)
- [Google Colab](https://colab.research.google.com/github/DrewThomasson/ebook2audiobook/blob/main/Notebooks/colab_ebook2audiobook.ipynb)
- [Kaggle](https://github.com/Rihcus/ebook2audiobookXTTS/blob/main/Notebooks/kaggle-ebook2audiobook.ipynb)

## Voice cloning

Upload any audio file in a supported format. 1 to 5 minutes of audio is ideal. E2A cleans background noise automatically.

The built-in voice list is mostly English. Contact the team to add voices in other languages.

## Fine-tuned models

### Train your own

Use [Universal TTS Finetune](https://github.com/DrewThomasson/Universal_TTS_Finetune).

### De-noise training data

Use [DeepFilterNet](https://github.com/Rikorose/DeepFilterNet) or the [Hugging Face Space](https://huggingface.co/spaces/drewThomasson/DeepFilterNet2_no_limit).

### Pre-trained model collection

Browse models on [Hugging Face](https://huggingface.co/drewThomasson/fineTunedTTSModels/tree/main).

XTTSv2 custom models require a reference audio clip (`ref.wav`) in the zip.

## Customization

Edit `lib/conf.py` to change default settings. Back up your modified `conf.py` before updating E2A, then restore it after the update. The same applies to `models.py`.

To submit a custom model as an official preset, contact the team.

## Older versions

Download releases from the [releases page](https://github.com/DrewThomasson/ebook2audiobook/releases).

```bash
git checkout tags/v25.7.7
```

## Troubleshooting

**GPU not detected.** See the [GPU Issues wiki page](https://github.com/DrewThomasson/ebook2audiobook/wiki/GPU-ISSUES).

**CPU is slow.** GPU conversion runs at near real-time speed. CPU is significantly slower, especially with modern TTS engines. See [this discussion](https://github.com/DrewThomasson/ebook2audiobook/discussions/19#discussioncomment-10879846) for alternatives.

**Dependency issues.** Use the Docker image. It is self-contained. Add `--help` to the Docker run command for options.

**Truncated audio.** [Open an issue](https://github.com/DrewThomasson/ebook2audiobook/issues). The team needs user input to improve sentence splitting for different languages.

> Before opening an issue, search the [open and closed issues](https://github.com/DrewThomasson/ebook2audiobook/issues?q=is%3Aissue) to check if your problem is already reported.

> EPUB files have no standard structure for chapters and paragraphs. Remove any text you do not want converted before running the tool.

## Roadmap

See [docs/ROADMAP.md](docs/ROADMAP.md) for planned features, TTS engine support, and platform compatibility status.

## Contributing

### Code style

- No blank lines between code (only between functions and classes)
- Single quotes for all keys except `dict()` and JSON
- 4-space indentation, no tabs
- Type annotations on all function arguments and return values
- No space between the argument and its type annotation

```python
def get_user(user_id:int, users:list[dict])->Optional[dict]:
    for user in users:
        if user['id'] == user_id:
            return user
    return None
```

### Hardware donations

The team accepts hardware for testing:

- NVIDIA cards with CUDA 11.8+
- Intel XPU cards
- AMD ROCm cards with ROCm 5.7+

Contact [@DrewThomasson](https://github.com/DrewThomasson).

## Sponsorship

If E2A helps your work, consider [sponsoring on Ko-fi](https://ko-fi.com/athomasson2).

## Special Thanks

To all financial and code contributors. Every contribution improves E2A.
