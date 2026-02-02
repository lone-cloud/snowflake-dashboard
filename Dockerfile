FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY logs-server.go .
RUN go build -ldflags="-s -w" -o server logs-server.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server
COPY index.html /tmp/index.html
COPY styles.css /app/static/styles.css
COPY script.js /app/static/script.js
COPY favicon.svg /app/static/favicon.svg
COPY favicon-dark.svg /app/static/favicon-dark.svg
COPY VERSION /tmp/VERSION
RUN VERSION=$(cat /tmp/VERSION) && \
    sed "s/?v=VERSION/?v=$VERSION/g" /tmp/index.html > /app/static/index.html

EXPOSE 8888

CMD ["/app/server"]
