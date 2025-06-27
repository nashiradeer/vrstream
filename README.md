# Nashira Deer // VRStream

Self-host your own streaming service to play anything you want on VRChat.

## Installation

This project has been created to be used in Docker Swarm mode, compiling directly inside the production environment, you can do this by running the following command: `docker stack deploy -c docker-compose.yml vrstream`.

This project hasn't been tested in Docker Compose mode, Kubernetes or in bare metal environments.

## Configuration

This project has HTTPS enabled by default using Cloudflare Origin CA and Cloudflare proxy. If don't want to use Cloudflare or HTTPS you will need to manually edit the `nginx.conf` and `docker-compose.yml` files.

1. Create a secret named `cloudflare-cert` with your Cloudflare Origin CA certificate.
2. Create a secret named `cloudflare-key` with your Cloudflare Origin CA private key.
3. Create a file named `vrstream.yml` with the following content:

```yaml
listen: :5000
users:
  - username: "admin"
    password: "your_password_here"
```

4. Use the following address to stream: `rtmp://your_server_ip:1935/live` with the stream key `admin?key=your_password_here`.
5. Use the following address to watch: `https://your_server_ip/admin`.

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
