all:
	$(MAKE) -C internal

.PHONY: build

VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -ldflags "-X 'github.com/ufukty/gonfique/cmd/gonfique/commands/version.Version=$(VERSION)'"
LDFLAGS_WASM := -ldflags "-X 'github.com/ufukty/gonfique/cmd/wasm.Version=$(VERSION)'"

build:
	@echo "Version $(VERSION)..."
	
	mkdir -p "build/gonfique/$(VERSION)"
	GOOS=darwin  GOARCH=amd64 go build -trimpath $(LDFLAGS) -o build/gonfique/$(VERSION)/gonfique-$(VERSION)-darwin-amd64  ./cmd/gonfique
	GOOS=darwin  GOARCH=arm64 go build -trimpath $(LDFLAGS) -o build/gonfique/$(VERSION)/gonfique-$(VERSION)-darwin-arm64  ./cmd/gonfique
	GOOS=linux   GOARCH=amd64 go build -trimpath $(LDFLAGS) -o build/gonfique/$(VERSION)/gonfique-$(VERSION)-linux-amd64   ./cmd/gonfique
	GOOS=linux   GOARCH=386   go build -trimpath $(LDFLAGS) -o build/gonfique/$(VERSION)/gonfique-$(VERSION)-linux-386     ./cmd/gonfique
	GOOS=linux   GOARCH=arm   go build -trimpath $(LDFLAGS) -o build/gonfique/$(VERSION)/gonfique-$(VERSION)-linux-arm     ./cmd/gonfique
	GOOS=linux   GOARCH=arm64 go build -trimpath $(LDFLAGS) -o build/gonfique/$(VERSION)/gonfique-$(VERSION)-linux-arm64   ./cmd/gonfique
	GOOS=freebsd GOARCH=amd64 go build -trimpath $(LDFLAGS) -o build/gonfique/$(VERSION)/gonfique-$(VERSION)-freebsd-amd64 ./cmd/gonfique
	GOOS=freebsd GOARCH=386   go build -trimpath $(LDFLAGS) -o build/gonfique/$(VERSION)/gonfique-$(VERSION)-freebsd-386   ./cmd/gonfique
	GOOS=freebsd GOARCH=arm   go build -trimpath $(LDFLAGS) -o build/gonfique/$(VERSION)/gonfique-$(VERSION)-freebsd-arm   ./cmd/gonfique
	
	mkdir -p "build/wasm/$(VERSION)"
	GOOS=js GOARCH=wasm go build -trimpath $(LDFLAGS_WASM) -o build/wasm/$(VERSION)/gonfique-$(VERSION)-js-wasm ./cmd/wasm

.PHONY: install

install:
	go build $(LDFLAGS) -o ~/bin/gonfique ./cmd/gonfique

README.toc.md: README.md
	pandoc -s --toc --toc-depth=6 --wrap=preserve README.md -o README.toc.md
	gsed --in-place 's/{.*}//g' README.toc.md