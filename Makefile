clean:
	go mod tidy

test:
	go test ./... -test.v

fmt:
	gofmt -w .
