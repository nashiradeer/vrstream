# Nashira Deer // VRStream

Self-host your own streaming service to play anything you want on VRChat.

## Installation

1. Install Docker Compose and run `docker-compose up -d` to start the services.
2. Use the following address to stream: `rtmp://your_server_ip:1935/live` with the stream key `admin?key=youshallnotpass`.
3. Use the following address to watch: `https://your_server_ip/admin`.

## Recommendations to OBS Studio

Use the following settings in OBS Studio:

- **Output Mode**: Advanced
- **Video Codec**: H.264
- **Audio Codec**: AAC
- **Resolution**: 1280x720
- **Rescale Output**: Lanczos
- **FPS**: 30
- **Rate Control**: CBR
- **Bitrate**: 3000 Kbps
- **Keyframe Interval**: 1 seconds
- **Profile**: baseline

### For Nvidia GPUs

- **Video Encoder**: NVIDIA NVENC H.264
- **Preset**: P6: Slower (Best Quality)
- **Tuning**: High Quality
- **Multipass Mode**: Two Pass (Quarter Resolution)
- **Look-ahead**: Enabled
- **Adaptive Quantization**: Enabled
- **B-frames**: 4

## License

This project is licensed under the [AGPL-3.0-or-later License](https://www.gnu.org/licenses/agpl-3.0.html).
