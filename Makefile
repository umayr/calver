.PHONY: default s clean fmt pretest lint lint-test list vet build test all
SHELL := /bin/bash
BINARY=calver

VERSION=
BUILD_TIME=`date +%FT%T%z`

BRANCH=`git rev-parse --abbrev-ref HEAD`
COMMIT=`git rev-parse --short HEAD`

LDFLAGS="-X ${BINARY}.version=${VERSION} -X ${BINARY}.buildtime=${BUILD_TIME} -X ${BINARY}.branch=${BRANCH} -X ${BINARY}.commit=${COMMIT}"

default: build

.PHONY: clean
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: pretest
pretest:
	@gofmt -d $$(find . -type f -name '*.go' -not -path "./vendor/*") 2>&1 | read; [ $$? == 1 ]

.PHONY: vet
vet:
	@go vet

.PHONY: test
test: pretest vet lint
	@go test -v -p=1

.PHONY: build
build: clean test
	@go build -x -ldflags ${LDFLAGS} -o bin/${BINARY} github.com/umayr/${BINARY}/cmd/${BINARY}

.PHONY: fmt
fmt:
	@gofmt -w $$(find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: lint
lint:
	@go get -v golang.org/x/lint/golint
	@golint ./... | grep -v vendor/ | true
