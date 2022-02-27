dev:
	nodemon --exec go run main.go --signal SIGTERM

build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .