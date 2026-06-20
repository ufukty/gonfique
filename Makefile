build-wasm:
	mkdir -p "build/gonfique-wasm"
	GOOS=js GOARCH=wasm go build -o build/gonfique-wasm/$(VERSION).wasm ./cmd/wasm
