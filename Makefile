VERSION = 1.0.0
CUR_DIR = $(shell pwd)
WORKSPACE = /go/src/github.com/hanks/terraform-variables-generator
DEV_IMAGE = hanks/tfvargen-dev:1.0.0
OS = $(shell uname -s | tr '[:upper:]' '[:lower:]')

.PHONY: dev push build test debug install uninstall clean

default: test

dev:
	docker build -t $(DEV_IMAGE) .

push:
	docker push $(DEV_IMAGE)

build: test clean
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) -e "CGO_ENABLED=0" -e "GOARCH=amd64" -e "GOOS=linux" $(DEV_IMAGE) go build -o ./dist/bin/tfvargen_linux_amd64_$(VERSION) main.go
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) -e "CGO_ENABLED=0" -e "GOARCH=amd64" -e "GOOS=darwin" $(DEV_IMAGE) go build -o ./dist/bin/tfvargen_darwin_amd64_$(VERSION) main.go

test:
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) $(DEV_IMAGE) go test -v -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) $(DEV_IMAGE) go vet ./...
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) $(DEV_IMAGE) golint -set_exit_status $(go list ./...)

debug:
	docker run -it --rm --security-opt=seccomp:unconfined -v $(CUR_DIR):$(WORKSPACE) $(DEV_IMAGE) dlv debug main.go

install:
	cp ./dist/bin/tfvargen_$(OS)_amd64_$(VERSION) /usr/local/bin/tfvargen

uninstall:
	rm /usr/local/bin/tfvargen

clean:
	rm -rf ./dist
