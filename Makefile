.PHONY: all
all: build test

NAME = cerebro
DOCKERHUB_USR = willis7
PWD := $(MKPATH:%/Makefile=%)

### Colour Definitions
END_COLOR=\x1b[0m
GREEN_COLOR=\x1b[32;01m
RED_COLOR=\x1b[31;01m
YELLOW_COLOR=\x1b[33;01m


end:
	@echo "$(YELLOW_COLOR)🤟🏼 🤟🏼 🤟🏼$(END_COLOR)"

clean:
	@echo "$(GREEN_COLOR)Cleaning unwanted files $(END_COLOR)"
	cd "$(PWD)"
	rm -rf vendor
	rm -rf coverage.txt
	rm -rf coverage.html
	rm -rf bin/

init:	go.check
	@echo "$(GREEN_COLOR)Initialising dep for the first time $(END_COLOR)"
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/golang/lint/golint

update:	go.check
	@echo "$(GREEN_COLOR)Running dep ensure $(END_COLOR)"
	dep ensure

fmt:	go.check
	@echo "$(GREEN_COLOR)Running fmt $(END_COLOR)"
	go fmt $(shell go list ./... | grep -v /vendor/)

vet:	go.check
	@echo "$(GREEN_COLOR)Running vet $(END_COLOR)"
	go vet $(shell go list ./... | grep -v /vendor/)

lint:	golint.check
	golint $(shell go list ./... | grep -v /vendor/)

test:	go.check
	@echo "$(GREEN_COLOR)Running tests for all packages $(END_COLOR)"
	go test ./... -v -p=5 -race -covermode=atomic -timeout=30s

compile:	go.check
	@echo "$(GREEN_COLOR)Compiling linux and mac binaries in ./bin $(END_COLOR)"
	mkdir -p bin/
	go build  -race -o bin/$(NAME)
	CGO_ENABLED=0 GOOS=linux go build -o bin/$(NAME)_linux

coverage:	go.check
	@echo "$(GREEN_COLOR)Calculating test coverage across packages $(END_COLOR)"
	@echo 'mode: atomic' > coverage.txt && echo '' > coverage.tmp && go list ./... | xargs -n1 -I{} sh -c 'go test -p=5 -race -covermode=atomic -coverprofile=coverage.tmp -timeout=30s {} && tail -n +2 coverage.tmp >> coverage.txt'
	go tool cover -html=coverage.txt -o coverage.html
	@rm coverage.txt
	@rm coverage.tmp
	@echo "$(YELLOW_COLOR)Run open ./coverage.html to view coverage $(END_COLOR)"

install:	go.check
	@echo "$(GREEN_COLOR)Installing all binaries $(END_COLOR)"
	go install ./...


build: fmt vet lint coverage install end

build_fresh: clean init update fmt vet lint coverage compile install end

build_docker:	docker.check
	@echo "$(GREEN_COLOR)Building a docker image $(END_COLOR)"
	docker build -t $(DOCKERHUB_USR)/$(NAME) .

run_docker:	docker.check
	@echo "$(GREEN_COLOR)Starting docker image $(END_COLOR)"
	docker run --rm $(DOCKERHUB_USR)/$(NAME)


# .check targets just tests for a command to be available on your PATH.
%.check:
	@which $*
