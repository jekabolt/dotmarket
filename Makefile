# REPOSITORY=jekabolt
# REGISTRY=jekabolt
IMAGE_NAME=dotmarket
VERSION=latest

build:
	go build -o ./bin/$(IMAGE_NAME) ./cmd/

run: build
	./bin/$(IMAGE_NAME)

# image:
# 	docker build -t $(REPOSITORY)/${IMAGE_NAME}:$(VERSION) -f ./Dockerfile . 
# 	docker tag $(REPOSITORY)/${IMAGE_NAME}:$(VERSION) $(REGISTRY)/${IMAGE_NAME}:$(VERSION)
