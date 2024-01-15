.PHONY: vendor clean

version=latest
image=hello-cafe/go/api

build: vendor build-api clean

vendor:
	go mod tidy -v && go mod vendor -v

clean:
	rm -r bin/api

build-api:
	GOOS=linux GOARCH=amd64 go build -o bin/api ./cmd
	DOCKER_BUILDKIT=1 docker build -t $(image):$(version) -f ./build/api.Dockerfile .

run-mysql:
	docker-compose up -d mysql

run:
	docker run -it -d --name hello-cafe -p 8000:8000 $(image):$(version)
