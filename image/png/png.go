// This file has automatically been generated on Wed Feb 26 02:10:01 +05 2020.
// DO NOT EDIT.
package png

import (
	"image"
	"image/png"
	"io"
	_ "unsafe"
)

//go:linkname Decode image/png.Decode
//go:noescape
func Decode(r io.Reader) (image.Image, error)

//go:linkname DecodeConfig image/png.DecodeConfig
//go:noescape
func DecodeConfig(r io.Reader) (image.Config, error)

//go:linkname Encode image/png.Encode
//go:noescape
func Encode(w io.Writer, m image.Image) error

//go:linkname encoderencode png.sub_encoderencode
func encoderencode(enc *png.Encoder, w io.Writer, m image.Image) error {
	return enc.Encode(w, m)
}

//go:linkname EncoderEncode png.sub_encoderencode
//go:noescape
func EncoderEncode(enc *png.Encoder, w io.Writer, m image.Image) error

//go:linkname formaterrorerror png.sub_formaterrorerror
func formaterrorerror(e png.FormatError) string {
	return e.Error()
}

//go:linkname FormatErrorError png.sub_formaterrorerror
//go:noescape
func FormatErrorError(e png.FormatError) string

//go:linkname unsupportederrorerror png.sub_unsupportederrorerror
func unsupportederrorerror(e png.UnsupportedError) string {
	return e.Error()
}

//go:linkname UnsupportedErrorError png.sub_unsupportederrorerror
//go:noescape
func UnsupportedErrorError(e png.UnsupportedError) string
