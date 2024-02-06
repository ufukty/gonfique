.PHONY: build

VERSION := $(shell git describe --tags --always --dirty)

.PHONY: update-examples
update-examples:
	@for example in examples/*; do \
        if [ -d "$$example" ]; then \
            echo "Making in $$example"; \
            $(MAKE) -C $$example output.go; \
        fi \
    done


build:
	@echo "Version $(VERSION)..."
	mkdir -p "build/$(VERSION)"
	GOOS=darwin GOARCH=amd64 go build -o build/$(VERSION)/gonfique-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o build/$(VERSION)/gonfique-darwin-arm64 .
	GOOS=linux GOARCH=amd64 go build -o build/$(VERSION)/gonfique-linux-amd64 .
	GOOS=linux GOARCH=386 go build -o build/$(VERSION)/gonfique-linux-386 .
	GOOS=linux GOARCH=arm go build -o build/$(VERSION)/gonfique-linux-arm .
	GOOS=linux GOARCH=arm64 go build -o build/$(VERSION)/gonfique-linux-arm64 .
	GOOS=freebsd GOARCH=amd64 go build -o build/$(VERSION)/gonfique-freebsd-amd64 .
	GOOS=freebsd GOARCH=386 go build -o build/$(VERSION)/gonfique-freebsd-386 .
	GOOS=freebsd GOARCH=arm go build -o build/$(VERSION)/gonfique-freebsd-arm .
