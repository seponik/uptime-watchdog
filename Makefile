run:
	go run cmd/uwdog/main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o uwdog cmd/uwdog/main.go

test: 
	go test -cover ./...

fmt:
	go fmt ./...