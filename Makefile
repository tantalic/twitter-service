build:
	go build

docker:
	docker run --rm -v "$(GOPATH)":/go -w /go/src/tantalic.com/twitter-service blang/golang-alpine go build -v
	docker build -t twitter-service .
	rm twitter-service 

