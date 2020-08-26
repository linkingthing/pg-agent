GOSRC = $(shell find . -type f -name '*.go')

IMAGE_VERSION=v1.0.2

build:
	CGO_ENABLED=0 GOOS=linux go build cmd/pg-agent/pg-agent.go

build-image:
	docker build -t linkingthing/pg-agent:${IMAGE_VERSION} .
	docker image prune -f

docker:
	docker build -t linkingthing/pg-agent:${IMAGE_VERSION} .
	docker image prune -f
	docker push linkingthing/pg-agent:${IMAGE_VERSION}

clean:
	rm -rf pg-agent

clean-image:
	docker rmi linkingthing/pg-agent:${IMAGE_VERSION}

.PHONY: build
