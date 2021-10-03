build:
	go build -o bin/pokcli main.go
setup:
	go mod download
test:
	go test -v ./...
