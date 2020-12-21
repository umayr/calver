.PHONY: default s clean fmt pretest lint lint-test list vet build test all
SHELL := /bin/bash
BINARY=calver

VERSION=0.1.2
BUILD_TIME=`date +%FT%T%z`

BRANCH=`git rev-parse --abbrev-ref HEAD`
COMMIT=`git rev-parse --short HEAD`

PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64 arm

LDFLAGS="-X ${BINARY}.version=${VERSION} -X ${BINARY}.buildtime=${BUILD_TIME} -X ${BINARY}.branch=${BRANCH} -X ${BINARY}.commit=${COMMIT}"

default: build

.PHONY: clean
clean:
	@rm -rf ./bin

.PHONY: pretest
pretest:
	@gofmt -d $$(find . -type f -name '*.go' -not -path "./vendor/*") 2>&1 | read; [ $$? == 1 ]

.PHONY: vet
vet:
	@go vet

.PHONY: lint
lint:
	@go get -v golang.org/x/lint/golint
	@golint ./... | grep -v vendor/ | true

.PHONY: test
test: pretest vet lint
	@go test -v -p=1

.PHONY: fmt
fmt:
	@gofmt -w $$(find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: build
build: clean test
	@go build -x -ldflags ${LDFLAGS} -o bin/${BINARY} github.com/umayr/${BINARY}/cmd/${BINARY}

.PHONY: all
all: clean test
	$(foreach GOOS, $(PLATFORMS),\
		$(foreach GOARCH, $(ARCHITECTURES),\
			$(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -x -ldflags ${LDFLAGS} -o bin/${BINARY}-${GOOS}-${GOARCH}/${BINARY} github.com/umayr/${BINARY}/cmd/${BINARY})\
		)\
	)

.PHONY: tarball
tarball: all
	@for folder in $(shell ls ./bin); do tar -C ./bin/$${folder} -cvzf ./bin/$${folder}.tar.gz ./${BINARY}; done
