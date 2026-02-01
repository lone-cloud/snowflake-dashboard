lint:
	~/go/bin/golangci-lint run

dev:
	go run dev-server.go

build:
	go build -ldflags="-s -w" -o server logs-server.go

docker:
	docker build -t snowflake-dashboard:latest .

test:
	go test -v ./...

release:
	gh workflow run release.yml

.PHONY: lint dev build docker test release
