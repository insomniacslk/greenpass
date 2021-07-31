package greenpass

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"image"
	"io"
	"strings"

	// load JPEG and PNG decoders
	_ "image/jpeg"
	_ "image/png"

	base45 "github.com/adrianrudnik/base45-go"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

// Read reads an io.Reader's data into a structure containing the EU COVID-19
// Green Pass data.
func Read(r io.Reader) (interface{}, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}
	return Decode(img)
}

// Decode decodes an image into a structure containing the EU COVID-19
// Green Pass data.
func Decode(img image.Image) (interface{}, error) {

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize QR object from image: %w", err)
	}

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decode QR code: %w", err)
	}

	// strip the leading "HC1:" string
	b45encoded := strings.TrimPrefix(result.GetText(), "HC1:")
	decoded, err := base45.Decode([]byte(b45encoded))
	if err != nil {
		return nil, fmt.Errorf("base45 decoding failed: %w", err)
	}

	compressed, err := zlib.NewReader(bytes.NewBuffer(decoded))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize zlib reader: %w", err)
	}
	var uncompressed bytes.Buffer
	if _, err := io.Copy(&uncompressed, compressed); err != nil {
		return nil, fmt.Errorf("failed to decompress data: %w", err)
	}
	// decode CBOR-2
	var cbordata cbor.Tag
	if err := cbor.Unmarshal(uncompressed.Bytes(), &cbordata); err != nil {
		return nil, fmt.Errorf("failed to decode CBOR-2 data: %w", err)
	}
	content, ok := cbordata.Content.([]interface{})
	if !ok {
		return nil, fmt.Errorf("CBOR-2 content is not an array")
	}
	contentBytes, ok := content[2].([]byte)
	if !ok {
		return nil, fmt.Errorf("CBOR-2 content of item 2 is not a byte array")
	}
	// decode the third item, which contains the green pass data
	var data map[interface{}]interface{}
	if err := cbor.Unmarshal(contentBytes, &data); err != nil {
		return nil, fmt.Errorf("failed to decode inner CBOR-2 data: %w", err)
	}
	return data, nil
}
