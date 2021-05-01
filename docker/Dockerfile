FROM alpine:3.13

RUN set -eux; \
	apkArch="$(apk --print-arch)"; \
	case "$apkArch" in \
		armhf) arch='armv6' ;; \
		aarch64) arch='arm64' ;; \
		x86_64) arch='amd64' ;; \
		*) echo >&2 "error: unsupported architecture: $apkArch"; exit 1 ;; \
	esac; \
	wget -O /tmp/servekit.tar.gz "https://github.com/cjaewon/servekit/releases/download/v2.0.0/servekit_2.0.0_linux_$arch.tar.gz"; \
	tar xzvf /tmp/servekit.tar.gz -C /usr/local/bin servekit; \
	rm -f /tmp/servekit.tar.gz; \
	chmod +x /usr/local/bin/servekit; 

CMD ["servekit"]

# Meta
LABEL org.opencontainers.image.vendor="Cjaewon" \
	org.opencontainers.image.url="https://github.com/cjaewon/servekit" \
	org.opencontainers.image.title="Servekit" \
	org.opencontainers.image.description="Micro HTTP server for serving static files." \
	org.opencontainers.image.version="v2.0.0" \
	org.opencontainers.image.documentation="https://github.com/cjaewon/servekit/blob/main/README.md"