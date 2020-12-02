install:
	go mod tidy
	go build -o friendly-potato

run:
	go run main.go
go-report:
	$(info =================== RUNNING GOREPORT CARD ENV ===================)
	go get github.com/gojp/goreportcard/cmd/goreportcard-cli
	export PATH=${PATH}:$$(go env GOPATH)/bin/; goreportcard-cli
docker-build:
	docker build -t friendly-potato:v1 .
