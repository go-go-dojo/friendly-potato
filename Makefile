#to run all tests don't forget to export the CF_TOKEN and the EXTERNAL_IP environment variables

install:
	go mod tidy
	go build -o friendly-potato

run:
	go run main.go

docker-build:
	docker build -t friendly-potato:v1 .
