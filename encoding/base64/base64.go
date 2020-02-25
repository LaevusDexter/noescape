// This file has automatically been generated on Wed Feb 26 02:09:51 +05 2020.
// DO NOT EDIT.
package base64

import (
	"encoding/base64"
	"io"
	_ "unsafe"
)

//go:linkname encodingstrict base64.sub_encodingstrict
func encodingstrict(enc base64.Encoding) *base64.Encoding {
	return enc.Strict()
}

//go:linkname EncodingStrict base64.sub_encodingstrict
//go:noescape
func EncodingStrict(enc base64.Encoding) *base64.Encoding

//go:linkname NewEncoder encoding/base64.NewEncoder
//go:noescape
func NewEncoder(enc *base64.Encoding, w io.Writer) io.WriteCloser

//go:linkname corruptinputerrorerror base64.sub_corruptinputerrorerror
func corruptinputerrorerror(e base64.CorruptInputError) string {
	return e.Error()
}

//go:linkname CorruptInputErrorError base64.sub_corruptinputerrorerror
//go:noescape
func CorruptInputErrorError(e base64.CorruptInputError) string

//go:linkname encodingdecodestring base64.sub_encodingdecodestring
func encodingdecodestring(enc *base64.Encoding, s string) ([]byte, error) {
	return enc.DecodeString(s)
}

//go:linkname EncodingDecodeString base64.sub_encodingdecodestring
//go:noescape
func EncodingDecodeString(enc *base64.Encoding, s string) ([]byte, error)

//go:linkname encodingdecodedlen base64.sub_encodingdecodedlen
func encodingdecodedlen(enc *base64.Encoding, n int) int {
	return enc.DecodedLen(n)
}

//go:linkname EncodingDecodedLen base64.sub_encodingdecodedlen
//go:noescape
func EncodingDecodedLen(enc *base64.Encoding, n int) int

//go:linkname encodingencodetostring base64.sub_encodingencodetostring
func encodingencodetostring(enc *base64.Encoding, src []byte) string {
	return enc.EncodeToString(src)
}

//go:linkname EncodingEncodeToString base64.sub_encodingencodetostring
//go:noescape
func EncodingEncodeToString(enc *base64.Encoding, src []byte) string

//go:linkname NewDecoder encoding/base64.NewDecoder
//go:noescape
func NewDecoder(enc *base64.Encoding, r io.Reader) io.Reader

//go:linkname NewEncoding encoding/base64.NewEncoding
//go:noescape
func NewEncoding(encoder string) *base64.Encoding

//go:linkname encodingdecode base64.sub_encodingdecode
func encodingdecode(enc *base64.Encoding, dst, src []byte) (int, error) {
	return enc.Decode(dst, src)
}

//go:linkname EncodingDecode base64.sub_encodingdecode
//go:noescape
func EncodingDecode(enc *base64.Encoding, dst, src []byte) (int, error)

//go:linkname encodingencodedlen base64.sub_encodingencodedlen
func encodingencodedlen(enc *base64.Encoding, n int) int {
	return enc.EncodedLen(n)
}

//go:linkname EncodingEncodedLen base64.sub_encodingencodedlen
//go:noescape
func EncodingEncodedLen(enc *base64.Encoding, n int) int

//go:linkname encodingwithpadding base64.sub_encodingwithpadding
func encodingwithpadding(enc base64.Encoding, padding rune) *base64.Encoding {
	return enc.WithPadding(padding)
}

//go:linkname EncodingWithPadding base64.sub_encodingwithpadding
//go:noescape
func EncodingWithPadding(enc base64.Encoding, padding rune) *base64.Encoding
