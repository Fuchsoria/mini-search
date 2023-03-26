test:
	go test -v .

bench:
	go test -benchmem -bench .

dev:
	go run ./server/cmd/server -config ./config.json

.PHONY: test bench