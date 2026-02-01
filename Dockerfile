FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY logs-server.go .
RUN go build -ldflags="-s -w" -o server logs-server.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server
COPY index.html /usr/share/nginx/html/index.html
COPY styles.css /usr/share/nginx/html/styles.css
COPY script.js /usr/share/nginx/html/script.js
COPY favicon.svg /usr/share/nginx/html/favicon.svg
COPY favicon-dark.svg /usr/share/nginx/html/favicon-dark.svg

EXPOSE 8888

CMD ["/app/server"]
