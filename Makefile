# メタ情報
NAME := apichecker
VERSION := $(shell git describe --tags --abbrev=0)
LDFLAGS := -X 'main.version=$(VERSION)'

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
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o build/$(NAME)

.PHONY: clean
clean:
	rm -rf build

