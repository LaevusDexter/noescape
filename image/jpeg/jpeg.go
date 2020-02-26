// This file has automatically been generated on Wed Feb 26 15:50:38 +05 2020.
// DO NOT EDIT.
package jpeg

import (
	"image"
	"image/jpeg"
	"io"
	_ "unsafe"
)

//go:linkname Decode image/jpeg.Decode
//go:noescape
func Decode(r io.Reader) (image.Image, error)

//go:linkname DecodeConfig image/jpeg.DecodeConfig
//go:noescape
func DecodeConfig(r io.Reader) (image.Config, error)

//go:linkname Encode image/jpeg.Encode
//go:noescape
func Encode(w io.Writer, m image.Image, o *jpeg.Options,) error

//go:linkname formaterrorerror jpeg.sub_formaterrorerror
func formaterrorerror(e jpeg.FormatError,) string {
	return e.Error()
}

//go:linkname FormatErrorError jpeg.sub_formaterrorerror
//go:noescape
func FormatErrorError(e jpeg.FormatError,) string

//go:linkname unsupportederrorerror jpeg.sub_unsupportederrorerror
func unsupportederrorerror(e jpeg.UnsupportedError,) string {
	return e.Error()
}

//go:linkname UnsupportedErrorError jpeg.sub_unsupportederrorerror
//go:noescape
func UnsupportedErrorError(e jpeg.UnsupportedError,) string
