FROM debian:bookworm-slim

RUN apt update -y && \
    apt install -y nginx libnginx-mod-rtmp

CMD [ "/usr/sbin/nginx", "-g", "daemon off;" ]
