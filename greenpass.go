package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
	"strings"

	base45 "github.com/adrianrudnik/base45-go"
	"github.com/fxamacker/cbor/v2"
	"github.com/kr/pretty"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func decode(r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return fmt.Errorf("failed to initialize QR object from image: %w", err)
	}

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return fmt.Errorf("failed to decode QR code: %w", err)
	}

	// strip the leading "HC1:" string
	b45encoded := strings.TrimPrefix(result.GetText(), "HC1:")
	decoded, err := base45.Decode([]byte(b45encoded))
	if err != nil {
		return fmt.Errorf("base45 decoding failed: %w", err)
	}

	compressed, err := zlib.NewReader(bytes.NewBuffer(decoded))
	if err != nil {
		return fmt.Errorf("failed to initialize zlib reader: %w", err)
	}
	var uncompressed bytes.Buffer
	if _, err := io.Copy(&uncompressed, compressed); err != nil {
		return fmt.Errorf("failed to decompress data: %w", err)
	}
	// decode CBOR-2
	var cbordata cbor.Tag
	if err := cbor.Unmarshal(uncompressed.Bytes(), &cbordata); err != nil {
		return fmt.Errorf("failed to decode CBOR-2 data: %w", err)
	}
	content, ok := cbordata.Content.([]interface{})
	if !ok {
		return fmt.Errorf("CBOR-2 content is not an array")
	}
	contentBytes, ok := content[2].([]byte)
	if !ok {
		return fmt.Errorf("CBOR-2 content of item 2 is not a byte array")
	}
	// decode the third item, which contains the green pass data
	var data map[interface{}]interface{}
	if err := cbor.Unmarshal(contentBytes, &data); err != nil {
		return fmt.Errorf("failed to decode inner CBOR-2 data: %w", err)
	}
	pretty.Print(data)
	return nil
}

func main() {
	// open and decode image file
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file name>\n", os.Args[0])
		os.Exit(1)
	}
	fname := os.Args[1]
	fd, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Failed to open %s: %v", fname, err)
	}
	if err := decode(fd); err != nil {
		log.Fatal(err)
	}
}
