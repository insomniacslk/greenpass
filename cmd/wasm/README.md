# wasm

WebAssembly wrapper for the greenpass library.

## Build

```
GOOS=js GOARCH=wasm go build -o greenpass.wasm
```

## Use

The output file `greenpass.wasm` needs to be integrated in your web application.
You need:
* `greenpass.wasm`, obviously
* HTML code that will load `greenpass.wasm`
* Go's `wasm_exec.js`

The HTML code has to load the Go WASM stub `wasm_exec.js`. You can find it in
`$(go env GOROOT)/misc/wasm/wasm_exec.js`.

Then you can use `greenpassRead` from JavaScript.

An example of HTML can be found at [index.html.example](index.html.example).


## Links

Some links that I found useful:
* https://github.com/golang/go/wiki/WebAssembly
* https://marianogappa.github.io/software/2020/04/01/webassembly-tinygo-cheesse/
* https://medium.com/swlh/getting-started-with-webassembly-and-go-by-building-an-image-to-ascii-converter-dea10bdf71f6
