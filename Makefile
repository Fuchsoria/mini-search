test:
	go test -v .

bench:
	go test -benchmem -bench .

.PHONY: test bench