VERSION = 1.0.1
CUR_DIR = $(shell pwd)
WORKSPACE = /go/src/github.com/hanks/terraform-variables-generator
DEV_IMAGE = hanks/tfvargen-dev:1.1.0
OS = $(shell uname -s | tr '[:upper:]' '[:lower:]')

.PHONY: dev run push build test debug install uninstall clean

default: test

dev:
	docker build -t $(DEV_IMAGE) .

run:
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) $(DEV_IMAGE) go run main.go cmd

push:
	docker push $(DEV_IMAGE)

build: test clean
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) -e "CGO_ENABLED=0" -e "GOARCH=amd64" -e "GOOS=linux" $(DEV_IMAGE) go build -o ./dist/bin/tfvargen_linux_amd64_$(VERSION) main.go
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) -e "CGO_ENABLED=0" -e "GOARCH=amd64" -e "GOOS=darwin" $(DEV_IMAGE) go build -o ./dist/bin/tfvargen_darwin_amd64_$(VERSION) main.go

test:
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) $(DEV_IMAGE) sh -c 'go vet $$(go list ./...)'
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) $(DEV_IMAGE) sh -c 'golint -set_exit_status $$(go list ./...)'
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) $(DEV_IMAGE) sh -c 'go test -v -covermode=count -coverprofile=coverage.out $$(go list ./... | grep -v /configs | grep -v /version)'
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) $(DEV_IMAGE) go tool cover -html=coverage.out -o coverage.html

coveralls:
	docker run -it --rm -v $(CUR_DIR):$(WORKSPACE) $(DEV_IMAGE) goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)

debug:
	docker run -it --rm --security-opt=seccomp:unconfined -v $(CUR_DIR):$(WORKSPACE) $(DEV_IMAGE) dlv debug main.go

install:
	cp ./dist/bin/tfvargen_$(OS)_amd64_$(VERSION) /usr/local/bin/tfvargen

uninstall:
	rm /usr/local/bin/tfvargen

clean:
	rm -rf ./dist
