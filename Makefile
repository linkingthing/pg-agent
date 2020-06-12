GOSRC = $(shell find . -type f -name '*.go')

build:
	CGO_ENABLED=0 GOOS=linux go build cmd/pg-agent/pg-agent.go

build-image:
	docker build -t linkingthing/pg-agent:latest .
	docker image prune -f

docker:
	docker build -t linkingthing/pg-agent:latest .
	docker image prune -f
	docker push linkingthing/pg-agent:latest

clean:
	rm -rf pg-agent

clean-image:
	docker rmi linkingthing/pg-agent:latest

.PHONY: build
