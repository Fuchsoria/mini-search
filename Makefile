BIN := "./bin/search"
BIN_WIN_64 := "./bin/search_win_x64.exe"
BIN_MAC_64 := "./bin/search_mac_x64"
BIN_MAC_ARM := "./bin/search_mac_arm"
BIN_LINUX_64 := "./bin/search_linux_x64"
BIN_LINUX_ARM := "./bin/search_linux_arm"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

test:
	go test -v .

bench:
	go test -benchmem -bench .

dev:
	go run ./server/cmd/server -config ./config.json

build:
	CGO_ENABLED=0 go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./server/cmd/server

build-win-64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -o $(BIN_WIN_64) -ldflags "$(LDFLAGS)" ./server/cmd/server

build-mac-64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o $(BIN_MAC_64) -ldflags "$(LDFLAGS)" ./server/cmd/server

build-mac-arm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -o $(BIN_MAC_ARM) -ldflags "$(LDFLAGS)" ./server/cmd/server

build-linux-64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o $(BIN_LINUX_64) -ldflags "$(LDFLAGS)" ./server/cmd/server

build-linux-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -o $(BIN_LINUX_ARM) -ldflags "$(LDFLAGS)" ./server/cmd/server

.PHONY: test bench