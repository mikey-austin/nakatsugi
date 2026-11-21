.PHONY: parser test

parser:
	go generate ./internal/...

test:
	go test ./internal/config \
		./internal/parser \
		./internal/health
