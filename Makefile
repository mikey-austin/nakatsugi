.PHONY: parser test

parser:
	go generate ./...

test:
	go test ./config ./parser
