// This file has automatically been generated on Wed Feb 26 02:09:53 +05 2020.
// DO NOT EDIT.
package json

import (
	"bytes"
	"encoding/json"
	"io"
	_ "unsafe"
)

//go:linkname MarshalIndent encoding/json.MarshalIndent
//go:noescape
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

//go:linkname Unmarshal encoding/json.Unmarshal
//go:noescape
func Unmarshal(data []byte, v interface{}) error

//go:linkname marshalererrorerror json.sub_marshalererrorerror
func marshalererrorerror(e *json.MarshalerError) string {
	return e.Error()
}

//go:linkname MarshalerErrorError json.sub_marshalererrorerror
//go:noescape
func MarshalerErrorError(e *json.MarshalerError) string

//go:linkname rawmessagemarshaljson json.sub_rawmessagemarshaljson
func rawmessagemarshaljson(m json.RawMessage) ([]byte, error) {
	return m.MarshalJSON()
}

//go:linkname RawMessageMarshalJSON json.sub_rawmessagemarshaljson
//go:noescape
func RawMessageMarshalJSON(m json.RawMessage) ([]byte, error)

//go:linkname unmarshalfielderrorerror json.sub_unmarshalfielderrorerror
func unmarshalfielderrorerror(e *json.UnmarshalFieldError) string {
	return e.Error()
}

//go:linkname UnmarshalFieldErrorError json.sub_unmarshalfielderrorerror
//go:noescape
func UnmarshalFieldErrorError(e *json.UnmarshalFieldError) string

//go:linkname Marshal encoding/json.Marshal
//go:noescape
func Marshal(v interface{}) ([]byte, error)

//go:linkname decoderbuffered json.sub_decoderbuffered
func decoderbuffered(dec *json.Decoder) io.Reader {
	return dec.Buffered()
}

//go:linkname DecoderBuffered json.sub_decoderbuffered
//go:noescape
func DecoderBuffered(dec *json.Decoder) io.Reader

//go:linkname invalidutf8errorerror json.sub_invalidutf8errorerror
func invalidutf8errorerror(e *json.InvalidUTF8Error) string {
	return e.Error()
}

//go:linkname InvalidUTF8ErrorError json.sub_invalidutf8errorerror
//go:noescape
func InvalidUTF8ErrorError(e *json.InvalidUTF8Error) string

//go:linkname rawmessageunmarshaljson json.sub_rawmessageunmarshaljson
func rawmessageunmarshaljson(m *json.RawMessage, data []byte) error {
	return m.UnmarshalJSON(data)
}

//go:linkname RawMessageUnmarshalJSON json.sub_rawmessageunmarshaljson
//go:noescape
func RawMessageUnmarshalJSON(m *json.RawMessage, data []byte) error

//go:linkname unsupportedvalueerrorerror json.sub_unsupportedvalueerrorerror
func unsupportedvalueerrorerror(e *json.UnsupportedValueError) string {
	return e.Error()
}

//go:linkname UnsupportedValueErrorError json.sub_unsupportedvalueerrorerror
//go:noescape
func UnsupportedValueErrorError(e *json.UnsupportedValueError) string

//go:linkname Compact encoding/json.Compact
//go:noescape
func Compact(dst *bytes.Buffer, src []byte) error

//go:linkname Indent encoding/json.Indent
//go:noescape
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error

//go:linkname decodertoken json.sub_decodertoken
func decodertoken(dec *json.Decoder) (json.Token, error) {
	return dec.Token()
}

//go:linkname DecoderToken json.sub_decodertoken
//go:noescape
func DecoderToken(dec *json.Decoder) (json.Token, error)

//go:linkname NewEncoder encoding/json.NewEncoder
//go:noescape
func NewEncoder(w io.Writer) *json.Encoder

//go:linkname decoderinputoffset json.sub_decoderinputoffset
func decoderinputoffset(dec *json.Decoder) int64 {
	return dec.InputOffset()
}

//go:linkname DecoderInputOffset json.sub_decoderinputoffset
//go:noescape
func DecoderInputOffset(dec *json.Decoder) int64

//go:linkname unmarshaltypeerrorerror json.sub_unmarshaltypeerrorerror
func unmarshaltypeerrorerror(e *json.UnmarshalTypeError) string {
	return e.Error()
}

//go:linkname UnmarshalTypeErrorError json.sub_unmarshaltypeerrorerror
//go:noescape
func UnmarshalTypeErrorError(e *json.UnmarshalTypeError) string

//go:linkname unsupportedtypeerrorerror json.sub_unsupportedtypeerrorerror
func unsupportedtypeerrorerror(e *json.UnsupportedTypeError) string {
	return e.Error()
}

//go:linkname UnsupportedTypeErrorError json.sub_unsupportedtypeerrorerror
//go:noescape
func UnsupportedTypeErrorError(e *json.UnsupportedTypeError) string

//go:linkname NewDecoder encoding/json.NewDecoder
//go:noescape
func NewDecoder(r io.Reader) *json.Decoder

//go:linkname syntaxerrorerror json.sub_syntaxerrorerror
func syntaxerrorerror(e *json.SyntaxError) string {
	return e.Error()
}

//go:linkname SyntaxErrorError json.sub_syntaxerrorerror
//go:noescape
func SyntaxErrorError(e *json.SyntaxError) string

//go:linkname encoderencode json.sub_encoderencode
func encoderencode(enc *json.Encoder, v interface{}) error {
	return enc.Encode(v)
}

//go:linkname EncoderEncode json.sub_encoderencode
//go:noescape
func EncoderEncode(enc *json.Encoder, v interface{}) error

//go:linkname numberstring json.sub_numberstring
func numberstring(n json.Number) string {
	return n.String()
}

//go:linkname NumberString json.sub_numberstring
//go:noescape
func NumberString(n json.Number) string

//go:linkname decoderdecode json.sub_decoderdecode
func decoderdecode(dec *json.Decoder, v interface{}) error {
	return dec.Decode(v)
}

//go:linkname DecoderDecode json.sub_decoderdecode
//go:noescape
func DecoderDecode(dec *json.Decoder, v interface{}) error

//go:linkname decodermore json.sub_decodermore
func decodermore(dec *json.Decoder) bool {
	return dec.More()
}

//go:linkname DecoderMore json.sub_decodermore
//go:noescape
func DecoderMore(dec *json.Decoder) bool

//go:linkname marshalererrorunwrap json.sub_marshalererrorunwrap
func marshalererrorunwrap(e *json.MarshalerError) error {
	return e.Unwrap()
}

//go:linkname MarshalerErrorUnwrap json.sub_marshalererrorunwrap
//go:noescape
func MarshalerErrorUnwrap(e *json.MarshalerError) error

//go:linkname Valid encoding/json.Valid
//go:noescape
func Valid(data []byte) bool

//go:linkname delimstring json.sub_delimstring
func delimstring(d json.Delim) string {
	return d.String()
}

//go:linkname DelimString json.sub_delimstring
//go:noescape
func DelimString(d json.Delim) string

//go:linkname invalidunmarshalerrorerror json.sub_invalidunmarshalerrorerror
func invalidunmarshalerrorerror(e *json.InvalidUnmarshalError) string {
	return e.Error()
}

//go:linkname InvalidUnmarshalErrorError json.sub_invalidunmarshalerrorerror
//go:noescape
func InvalidUnmarshalErrorError(e *json.InvalidUnmarshalError) string
