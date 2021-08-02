package main

import (
	"bytes"
	"fmt"
	"syscall/js"

	"github.com/insomniacslk/greenpass"
)

func jsRead() js.Func {
	fn := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "need exactly one argument in call to greenpassRead"
		}
		buf := bytes.NewBuffer([]byte(args[0].String()))
		gp, err := greenpass.Read(buf)
		if err != nil {
			return fmt.Errorf("failed to read QR code: %w", err).Error()
		}
		return gp.SummaryAsHTML()
	})
	return fn
}

func main() {
	fmt.Println("Loaded greenpass WASM module")
	js.Global().Set("greenpassRead", jsRead())

	// keep the code running or it will fail with "Go program has already exited"
	select {}
}
