IMAGE_NAME ?= hub.c.163.com/bingohuang/go-playground-webapp
TAG ?= 1.8-bingo
CONTAINER_NAME  ?= go-playground-webapp
TAG ?= latest

require-glide:
	@if [ -z "$$(which glide)" ]; then \
		echo 'Missing "glide" command. See https://github.com/Masterminds/glide' && \
		exit 1; \
	fi

require-uglifyjs:
	@if [ -z "$$(which uglifyjs)" ]; then \
		echo 'Missing "uglifyjs" command. See https://github.com/mishoo/UglifyJS' && \
		exit 1; \
	fi

docker-build: require-glide require-uglifyjs
	glide install && \
	make -C static && \
	GOOS=linux GOARCH=amd64 go build -o app_linux_amd64 && \
	docker build -t $(IMAGE_NAME) .

docker-run:
	(docker stop $(CONTAINER_NAME) || exit 0) && \
	(docker rm $(CONTAINER_NAME) || exit 0) && \
	docker run -d -p 127.0.0.1:3000:3000 --name $(CONTAINER_NAME) -t $(IMAGE_NAME)

docker-push: docker-build
	docker tag $(IMAGE_NAME) $(IMAGE_NAME):$(TAG) && \
	docker push $(IMAGE_NAME):$(TAG)
