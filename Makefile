PROJECTNAME := $(shell basename "$(PWD)")
GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin
GOFILES := $(wildcard *.go)

install:
	go mod download && go mod verify
	
build :
	@-go build -o ${GOBIN}/${PROJECTNAME} ./cmd/app/${GOFILES}

run: build
	@-${GOBIN}/./${PROJECTNAME}

test: 
	go test -v -race ./...

docker-run: 
	docker compose up -d