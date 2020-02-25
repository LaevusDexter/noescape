// This file has automatically been generated on Wed Feb 26 02:10:00 +05 2020.
// DO NOT EDIT.
package gif

import (
	"image"
	"image/gif"
	"io"
	_ "unsafe"
)

//go:linkname Decode image/gif.Decode
//go:noescape
func Decode(r io.Reader) (image.Image, error)

//go:linkname DecodeConfig image/gif.DecodeConfig
//go:noescape
func DecodeConfig(r io.Reader) (image.Config, error)

//go:linkname Encode image/gif.Encode
//go:noescape
func Encode(w io.Writer, m image.Image, o *gif.Options) error

//go:linkname EncodeAll image/gif.EncodeAll
//go:noescape
func EncodeAll(w io.Writer, g *gif.GIF) error

//go:linkname DecodeAll image/gif.DecodeAll
//go:noescape
func DecodeAll(r io.Reader) (*gif.GIF, error)
