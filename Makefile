.PHONY: build test

ifdef VERSION
    VERSION=${VERSION}
else
    VERSION="v0.0.1"
endif

build:
	GOOS=darwin go build -ldflags "-X github.com/dellkeji/bcs-create-chart/pkg/version.Version=${VERSION}"  -o bin/darwin/amd64/bcs-create main.go

test:
	go test -v ./... -cover
