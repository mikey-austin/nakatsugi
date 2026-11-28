.PHONY: parser test

generate:
	go generate ./internal/...

vet:
	go vet ./internal/config \
		./internal/health \
		./internal/ipc

test:
	go test ./internal/config \
		./internal/parser \
		./internal/health \
		./internal/ipc
