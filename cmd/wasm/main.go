//go:build js && wasm
// +build js,wasm

package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"syscall/js"

	"github.com/ufukty/gonfique/internal/files/input/encoders"
	"github.com/ufukty/gonfique/internal/generates"
)

var Version string

//export Convert
func Convert(this js.Value, args []js.Value) (string, error) {
	if len(args) < 3 {
		return "", fmt.Errorf("Error: 3 arguments required: inputContent, inputMode, configContent")
	}

	var (
		input  = args[0].String()
		enc    = args[1].String()
		config = args[2].String()
	)

	enc2, err := encoders.FromString(enc)
	if err != nil {
		return "", fmt.Errorf("validating language: %w", err)
	}

	i := strings.NewReader(input)
	var c io.Reader
	if config != "" {
		c = strings.NewReader(config)
	}
	o := bytes.NewBufferString("")

	err = generates.FromReaders(i, enc2, c, o, false)
	if err != nil {
		return "", fmt.Errorf("generation: %w", err)
	}

	return o.String(), nil
}

func lrp2[R1 any](f func(this js.Value, args []js.Value) (R1, error)) func(this js.Value, args []js.Value) any {
	return func(this js.Value, args []js.Value) any {
		r1, e := f(this, args)
		return js.ValueOf([]any{r1, e.Error()})
	}
}

func main() {
	// Make the Convert function available in JS
	js.Global().Set("Convert", js.FuncOf(lrp2(Convert)))
	select {}
}
