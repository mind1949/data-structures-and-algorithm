fmt:
	gofmt -w -s .
	go mod tidy
test:
	 go test -v ./...