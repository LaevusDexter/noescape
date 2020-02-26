// This file has automatically been generated on Wed Feb 26 15:50:28 +05 2020.
// DO NOT EDIT.
package base32

import (
	"encoding/base32"
	"io"
	_ "unsafe"
)

//go:linkname NewDecoder encoding/base32.NewDecoder
//go:noescape
func NewDecoder(enc *base32.Encoding, r io.Reader) io.Reader

//go:linkname NewEncoder encoding/base32.NewEncoder
//go:noescape
func NewEncoder(enc *base32.Encoding, w io.Writer) io.WriteCloser

//go:linkname encodingencodetostring base32.sub_encodingencodetostring
func encodingencodetostring(enc *base32.Encoding, src []byte) string {
	return enc.EncodeToString(src)
}

//go:linkname EncodingEncodeToString base32.sub_encodingencodetostring
//go:noescape
func EncodingEncodeToString(enc *base32.Encoding, src []byte) string

//go:linkname encodingencodedlen base32.sub_encodingencodedlen
func encodingencodedlen(enc *base32.Encoding, n int) int {
	return enc.EncodedLen(n)
}

//go:linkname EncodingEncodedLen base32.sub_encodingencodedlen
//go:noescape
func EncodingEncodedLen(enc *base32.Encoding, n int) int

//go:linkname encodingwithpadding base32.sub_encodingwithpadding
func encodingwithpadding(enc base32.Encoding, padding rune) *base32.Encoding {
	return enc.WithPadding(padding)
}

//go:linkname EncodingWithPadding base32.sub_encodingwithpadding
//go:noescape
func EncodingWithPadding(enc base32.Encoding, padding rune) *base32.Encoding

//go:linkname corruptinputerrorerror base32.sub_corruptinputerrorerror
func corruptinputerrorerror(e base32.CorruptInputError,) string {
	return e.Error()
}

//go:linkname CorruptInputErrorError base32.sub_corruptinputerrorerror
//go:noescape
func CorruptInputErrorError(e base32.CorruptInputError,) string

//go:linkname NewEncoding encoding/base32.NewEncoding
//go:noescape
func NewEncoding(encoder string) *base32.Encoding

//go:linkname encodingdecode base32.sub_encodingdecode
func encodingdecode(enc *base32.Encoding, dst, src []byte) (int, error) {
	return enc.Decode(dst, src)
}

//go:linkname EncodingDecode base32.sub_encodingdecode
//go:noescape
func EncodingDecode(enc *base32.Encoding, dst, src []byte) (int, error)

//go:linkname encodingdecodestring base32.sub_encodingdecodestring
func encodingdecodestring(enc *base32.Encoding, s string) ([]byte, error) {
	return enc.DecodeString(s)
}

//go:linkname EncodingDecodeString base32.sub_encodingdecodestring
//go:noescape
func EncodingDecodeString(enc *base32.Encoding, s string) ([]byte, error)

//go:linkname encodingdecodedlen base32.sub_encodingdecodedlen
func encodingdecodedlen(enc *base32.Encoding, n int) int {
	return enc.DecodedLen(n)
}

//go:linkname EncodingDecodedLen base32.sub_encodingdecodedlen
//go:noescape
func EncodingDecodedLen(enc *base32.Encoding, n int) int
