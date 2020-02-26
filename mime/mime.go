// This file has automatically been generated on Wed Feb 26 15:50:42 +05 2020.
// DO NOT EDIT.
package mime

import (
	"mime"
	_ "unsafe"
)

//go:linkname worddecoderdecode mime.sub_worddecoderdecode
func worddecoderdecode(d *mime.WordDecoder, word string) (string, error) {
	return d.Decode(word)
}

//go:linkname WordDecoderDecode mime.sub_worddecoderdecode
//go:noescape
func WordDecoderDecode(d *mime.WordDecoder, word string) (string, error)

//go:linkname worddecoderdecodeheader mime.sub_worddecoderdecodeheader
func worddecoderdecodeheader(d *mime.WordDecoder, header string) (string, error) {
	return d.DecodeHeader(header)
}

//go:linkname WordDecoderDecodeHeader mime.sub_worddecoderdecodeheader
//go:noescape
func WordDecoderDecodeHeader(d *mime.WordDecoder, header string) (string, error)

//go:linkname wordencoderencode mime.sub_wordencoderencode
func wordencoderencode(e mime.WordEncoder, charset, s string) string {
	return e.Encode(charset, s)
}

//go:linkname WordEncoderEncode mime.sub_wordencoderencode
//go:noescape
func WordEncoderEncode(e mime.WordEncoder, charset, s string) string

//go:linkname AddExtensionType mime.AddExtensionType
//go:noescape
func AddExtensionType(ext, typ string) error

//go:linkname ExtensionsByType mime.ExtensionsByType
//go:noescape
func ExtensionsByType(typ string) ([]string, error)

//go:linkname FormatMediaType mime.FormatMediaType
//go:noescape
func FormatMediaType(t string, param map[string]string) string

//go:linkname ParseMediaType mime.ParseMediaType
//go:noescape
func ParseMediaType(v string) (string, map[string]string, error)

//go:linkname TypeByExtension mime.TypeByExtension
//go:noescape
func TypeByExtension(ext string) string
