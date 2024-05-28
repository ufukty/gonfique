.PHONY: build

VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -ldflags "-X 'main.Version=$(VERSION)'"

build:
	@echo "Version $(VERSION)..."
	mkdir -p "build/$(VERSION)"
	GOOS=darwin  GOARCH=amd64 go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-darwin-amd64  .
	GOOS=darwin  GOARCH=arm64 go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-darwin-arm64  .
	GOOS=linux   GOARCH=amd64 go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-linux-amd64   .
	GOOS=linux   GOARCH=386   go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-linux-386     .
	GOOS=linux   GOARCH=arm   go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-linux-arm     .
	GOOS=linux   GOARCH=arm64 go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-linux-arm64   .
	GOOS=freebsd GOARCH=amd64 go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-freebsd-amd64 .
	GOOS=freebsd GOARCH=386   go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-freebsd-386   .
	GOOS=freebsd GOARCH=arm   go build -trimpath $(LDFLAGS) -o build/$(VERSION)/gonfique-$(VERSION)-freebsd-arm   .

.PHONY: install

install:
	go build $(LDFLAGS) -o ~/bin/gonfique  .