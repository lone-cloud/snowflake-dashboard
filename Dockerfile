FROM oven/bun:1.3.7-debian

RUN apt-get update && apt-get install -y nginx docker.io && rm -rf /var/lib/apt/lists/*

COPY nginx.conf /etc/nginx/nginx.conf
COPY index.html /usr/share/nginx/html/index.html
COPY styles.css /usr/share/nginx/html/styles.css
COPY favicon.svg /usr/share/nginx/html/favicon.svg
COPY logs-server.js /app/logs-server.js

RUN echo '#!/bin/sh' > /start.sh && \
    echo 'nginx' >> /start.sh && \
    echo 'cd /app && bun logs-server.js' >> /start.sh && \
    chmod +x /start.sh

EXPOSE 8888

CMD ["/start.sh"]
