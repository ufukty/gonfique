.PHONY: build

VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -ldflags "-X 'github.com/ufukty/gonfique/cmd/main/commands/version.Version=$(VERSION)'"

build:
	@echo "Version $(VERSION)..."
	mkdir -p "build/$(VERSION)"
	GOOS=darwin  GOARCH=amd64 go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-darwin-amd64  ./cmd/main
	GOOS=darwin  GOARCH=arm64 go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-darwin-arm64  ./cmd/main
	GOOS=linux   GOARCH=amd64 go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-linux-amd64   ./cmd/main
	GOOS=linux   GOARCH=386   go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-linux-386     ./cmd/main
	GOOS=linux   GOARCH=arm   go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-linux-arm     ./cmd/main
	GOOS=linux   GOARCH=arm64 go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-linux-arm64   ./cmd/main
	GOOS=freebsd GOARCH=amd64 go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-freebsd-amd64 ./cmd/main
	GOOS=freebsd GOARCH=386   go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-freebsd-386   ./cmd/main
	GOOS=freebsd GOARCH=arm   go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-freebsd-arm   ./cmd/main

.PHONY: install

install:
	go build $(LDFLAGS) -o ~/bin/gonfique  ./cmd/main