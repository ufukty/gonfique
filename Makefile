all:
	$(MAKE) -C internal

build-wasm:
	mkdir -p "build/gonfique-wasm"
	GOOS=js GOARCH=wasm go build -trimpath $(LDFLAGS) -o build/gonfique-wasm/$(VERSION).wasm ./cmd/wasm

build: build-wasm

README.toc.md: README.md
	pandoc -s --toc --toc-depth=6 --wrap=preserve README.md -o README.toc.md
	gsed --in-place 's/{.*}//g' README.toc.md