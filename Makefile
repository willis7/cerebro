NAME = cerebro
DOCKERHUB_USR = willis7
PWD := $(MKPATH:%/Makefile=%)

.DEFAULT_GOAL := test

clean:
	cd "$(PWD)"
	rm -rf vendor


build:	docker.check
	docker build -t $(DOCKERHUB_USR)/$(NAME) .

run:	docker.check
	docker run --rm $(DOCKERHUB_USR)/$(NAME)

test:	go.check
	go test -race -v $(shell go list ./... | grep -v /vendor/)

coverage:	go.check
	go test -race -cover -v $(shell go list ./... | grep -v /vendor/)

vet:	go.check
	go vet $(shell go list ./... | grep -v /vendor/)


lint:	golint.check
	golint $(shell go list ./... | grep -v /vendor/)


# .check targets just tests for a command to be available on your PATH.
%.check:
	@which $*
