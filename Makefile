install:
	go mod tidy
	go build -o friendly-potato

run:
	go run main.go

docker-build:
	docker build -t friendly-potato:v1 .
