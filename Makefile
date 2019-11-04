DOCKER_REPO=geekodour
DOCKER_IMAGE_NAME=ghcs
DOCKER_IMAGE_TAG=0.0.2

build:
	go build -o ghcs -a -tags netgo .
docker-build: build
	docker build -t "$(DOCKER_REPO)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .
docker-push:
	docker push "$(DOCKER_REPO)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)"
clean:
	rm -rf ghcs