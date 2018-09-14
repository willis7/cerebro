NAME = cerebro
PWD := $(MKPATH:%/Makefile=%)

.DEFAULT_GOAL := test

clean:
	cd "$(PWD)"
	rm -rf vendor


build:
	go build -race -o $(NAME)

run:
	./$(NAME)

test:
	go test -race -v $(shell go list ./... | grep -v /vendor/)

coverage:
	go test -race -cover -v $(shell go list ./... | grep -v /vendor/)

vet:
	go vet $(shell go list ./... | grep -v /vendor/)

#staticcheck:
#	go get -u honnef.co/go/staticcheck/cmd/staticcheck
#	staticcheck $(shell go list ./... | grep -v /vendor/)

lint:
	golint $(shell go list ./... | grep -v /vendor/)

#simple:
#	go get -u honnef.co/go/simple/cmd/gosimple
#    gosimple $(go list ./... | grep -v "vendor")
