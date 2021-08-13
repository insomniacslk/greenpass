package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"syscall/js"

	"github.com/insomniacslk/greenpass"
)

// max size of a QR code that will be parsed, in bytes.
const maxSize = 10000

func jsFromBytes() js.Func {
	fn := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "need exactly one argument in call to greenpassFromBytes"
		}
		// max size of an image that will be decoded
		size := args[0].Length()
		if size > maxSize {
			return fmt.Errorf("image too large: maximum allowed is %d bytes", maxSize)
		}
		buf := make([]byte, size)
		js.CopyBytesToGo(buf, args[0])
		gp, err := greenpass.FromBytes(buf)
		if err != nil {
			return fmt.Errorf("failed to read QR code: %w", err).Error()
		}
		return gp.SummaryAsHTML()
	})
	return fn
}

func main() {
	fmt.Println("Loaded greenpass WASM module")
	js.Global().Set("greenpassFromBytes", jsFromBytes())

	// keep the code running or it will fail with "Go program has already exited"
	select {}
}
