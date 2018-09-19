# Cerebro
[![Build Status](https://travis-ci.org/willis7/cerebro.svg?branch=master)](https://travis-ci.org/willis7/cerebro)


[![codecov](https://codecov.io/gh/willis7/cerebro/branch/master/graph/badge.svg)](https://codecov.io/gh/willis7/cerebro)



Cerebro is a build and release monitor. You load it with the environments and deployments you have and it will allow you to query and display the current state. Eventually...


### Prerequisites

Here are the things that you would require before you get started

1. [Install git](https://www.atlassian.com/git/tutorials/install-git)
1. [Install golang](https://golang.org/doc/install)
1. [Install docker](https://docs.docker.com/install/#supported-platforms), we use it both for deployment and development

## Running the tests

If you would like to run the automated tests for the complete package, run this

```bash
make coverage
open ./coverage.html
```

### And coding style tests

We use the default golang coding conventions. Run the following to test for those

```bash
make fmt
make vet
make lint
```

## Built With

* [DEP](https://github.com/golang/dep) - For dependency management
* [PQ](github.com/lib/pq) - SQL driver for postgres
* [SQLX](github.com/jmoiron/sqlx) - For connecting to postgres
* [TESTIFY](github.com/stretchr/testify) - For asserting tests
* [GO-SQLMOCK](github.com/DATA-DOG/go-sqlmock) - For mocking postgres
