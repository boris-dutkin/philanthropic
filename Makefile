dev:
	ENV=dev go run server.go

prod:
	ENV=dev go run server.go

build:
	go build -o bin server.go