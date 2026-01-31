FROM oven/bun:1.3.6-debian AS builder

WORKDIR /app
COPY logs-server.ts .
RUN bun build --compile logs-server.ts --outfile server

FROM debian:13.3-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    nginx \
    docker.io \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY nginx.conf /etc/nginx/nginx.conf
COPY index.html /usr/share/nginx/html/index.html
COPY styles.css /usr/share/nginx/html/styles.css
COPY favicon.svg /usr/share/nginx/html/favicon.svg
COPY favicon-dark.svg /usr/share/nginx/html/favicon-dark.svg
COPY --from=builder /app/server /app/server

RUN echo '#!/bin/sh' > /start.sh && \
    echo 'nginx' >> /start.sh && \
    echo '/app/server' >> /start.sh && \
    chmod +x /start.sh

EXPOSE 8888

CMD ["/start.sh"]
