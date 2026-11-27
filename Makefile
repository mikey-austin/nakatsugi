.PHONY: parser test

generate:
	go generate ./internal/...

test:
	go test ./internal/config \
		./internal/parser \
		./internal/health \
		./internal/ipc
