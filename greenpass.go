// Package greenpass provides functions to decode QR-encoded EU digital
// COVID-19 certificates.
//
// See
// https://ec.europa.eu/health/sites/default/files/ehealth/docs/covid-certificate_json_specification_en.pdf
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
func Read(r io.Reader) (*CovidCertificate, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}
	return Decode(img)
}

// FromBytes reads an EU COVID-19 Green Pass from bytes.
func FromBytes(imagedata []byte) (*CovidCertificate, error) {
	return Read(bytes.NewReader(imagedata))
}

// Decode decodes an image into a structure containing the EU COVID-19
// Green Pass data.
func Decode(img image.Image) (*CovidCertificate, error) {

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

	// strip the leading "HC1:" string. HC1 stands for Health Certificate
	// version 1.
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
	var data map[int]cbor.RawMessage
	if err := cbor.Unmarshal(contentBytes, &data); err != nil {
		return nil, fmt.Errorf("failed to decode inner CBOR-2 data: %w", err)
	}
	// FIXME figure out what the `-260` key means (it contains the actual
	// JSON content of the digital covid certificate), and what the
	// remaining bits of data are precisely. One is the country code, the
	// other ones look like timestamps.
	// Also figure out why the data is further nested in a field with key 0x1.
	ccraw1, ok := data[-260]
	if !ok {
		return nil, fmt.Errorf("covid certificate JSON not found: key -260 is missing")
	}
	var data2 map[int]CovidCertificate
	if err := cbor.Unmarshal(ccraw1, &data2); err != nil {
		return nil, fmt.Errorf("failed to decode inner CBOR-2 data: %w", err)
	}
	cc, ok := data2[1]
	if !ok {
		return nil, fmt.Errorf("covid certificate JSON not found: key 1 is missing")
	}
	if err := cc.Validate(); err != nil {
		return nil, err
	}

	return &cc, nil
}
