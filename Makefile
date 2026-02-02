lint:
	~/go/bin/golangci-lint run

dev:
	@test -f ~/go/bin/air || (echo "Installing air..." && go install github.com/air-verse/air@latest)
	~/go/bin/air

build:
	go build -ldflags="-s -w" -o server dashboard-server.go

docker:
	docker build -t snowflake-dashboard:latest .

test:
	go test -v ./...

release:
	gh workflow run release.yml

.PHONY: lint dev build docker test release
