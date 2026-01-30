FROM node:alpine

RUN apk add --no-cache nginx docker-cli

COPY nginx.conf /etc/nginx/nginx.conf
COPY index.html /usr/share/nginx/html/index.html
COPY styles.css /usr/share/nginx/html/styles.css
COPY favicon.svg /usr/share/nginx/html/favicon.svg
COPY favicon-dark.svg /usr/share/nginx/html/favicon-dark.svg
COPY logs-server.js /app/logs-server.js

RUN echo '#!/bin/sh' > /start.sh && \
    echo 'nginx' >> /start.sh && \
    echo 'cd /app && node logs-server.js' >> /start.sh && \
    chmod +x /start.sh

EXPOSE 8888

CMD ["/start.sh"]
