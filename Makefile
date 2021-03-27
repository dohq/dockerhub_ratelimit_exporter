IMAGE_NAME=$(REPOSITORY):$(VERSION)
LATEST=$(REPOSITORY):latest
CACHE=$(REPOSITORY):cache-latest

container_build:
	docker build --cache-from=$(CACHE) \
		--build-arg=BUILDKIT_INLINE_CACHE=1 \
		-t $(IMAGE_NAME) \
		-t $(LATEST) .

container_push:
	docker push $(IMAGE_NAME)
	docker push $(CACHE)

# If necessary, add --target
container_cache:
	docker build \
		--target=builder \
		--build-arg=BUILDKIT_INLINE_CACHE=1 \
		-t $(CACHE) .
	docker push $(CACHE)

build:
	go build
test:
	go test -v -race ./...
run:
	go run .
help:
	go run *.go -h
