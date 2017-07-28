IMAGE_NAME ?= hub.c.163.com/bingohuang/go-playground-sandbox
TAG ?= 1.8-bingo
CONTAINER_NAME  ?= go-playground-sandbox

docker: Dockerfile
	docker build -t $(IMAGE_NAME) .

test: docker
	go test
	docker run --rm $(IMAGE_NAME) test

run:
	docker run -d -p 8080:8080 --name $(CONTAINER_NAME) -t $(IMAGE_NAME)

push:
	docker push $(IMAGE_NAME)
