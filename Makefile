# メタ情報
NAME := apichecker
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)'


.PHONY: imports
imports:
	goimports -w .

.PHONY: lint
lint:
	golint .

.PHONY: vet
vet:
	go vet .

.PHONY: test
test:lint vet imports
	go test .

.PHONY: build
build: test
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o build/$(NAME)

.PHONY: clean
clean:
	rm -rf build

