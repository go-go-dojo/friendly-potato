install:
	go mod tidy
	go build -o friendly-potato

run:
	go run main.go

test:
	go test -v -coverprofile c.out ./...

docker-build:
	docker build -t friendly-potato:v1 .
