dev:
	ENV=dev go run ./cmd/server.go

prod:
	ENV=prod go run ./cmd/server.go

build:
	go build -o bin ./cmd/server.go