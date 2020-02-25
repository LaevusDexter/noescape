// This file has automatically been generated on Wed Feb 26 02:09:53 +05 2020.
// DO NOT EDIT.
package xml

import (
	"encoding/xml"
	"io"
	_ "unsafe"
)

//go:linkname procinstcopy xml.sub_procinstcopy
func procinstcopy(p xml.ProcInst) xml.ProcInst {
	return p.Copy()
}

//go:linkname ProcInstCopy xml.sub_procinstcopy
//go:noescape
func ProcInstCopy(p xml.ProcInst) xml.ProcInst

//go:linkname decoderinputoffset xml.sub_decoderinputoffset
func decoderinputoffset(d *xml.Decoder) int64 {
	return d.InputOffset()
}

//go:linkname DecoderInputOffset xml.sub_decoderinputoffset
//go:noescape
func DecoderInputOffset(d *xml.Decoder) int64

//go:linkname MarshalIndent encoding/xml.MarshalIndent
//go:noescape
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

//go:linkname directivecopy xml.sub_directivecopy
func directivecopy(d xml.Directive) xml.Directive {
	return d.Copy()
}

//go:linkname DirectiveCopy xml.sub_directivecopy
//go:noescape
func DirectiveCopy(d xml.Directive) xml.Directive

//go:linkname startelementend xml.sub_startelementend
func startelementend(e xml.StartElement) xml.EndElement {
	return e.End()
}

//go:linkname StartElementEnd xml.sub_startelementend
//go:noescape
func StartElementEnd(e xml.StartElement) xml.EndElement

//go:linkname decoderdecodeelement xml.sub_decoderdecodeelement
func decoderdecodeelement(d *xml.Decoder, v interface{}, start *xml.StartElement) error {
	return d.DecodeElement(v, start)
}

//go:linkname DecoderDecodeElement xml.sub_decoderdecodeelement
//go:noescape
func DecoderDecodeElement(d *xml.Decoder, v interface{}, start *xml.StartElement) error

//go:linkname decodertoken xml.sub_decodertoken
func decodertoken(d *xml.Decoder) (xml.Token, error) {
	return d.Token()
}

//go:linkname DecoderToken xml.sub_decodertoken
//go:noescape
func DecoderToken(d *xml.Decoder) (xml.Token, error)

//go:linkname encoderencode xml.sub_encoderencode
func encoderencode(enc *xml.Encoder, v interface{}) error {
	return enc.Encode(v)
}

//go:linkname EncoderEncode xml.sub_encoderencode
//go:noescape
func EncoderEncode(enc *xml.Encoder, v interface{}) error

//go:linkname CopyToken encoding/xml.CopyToken
//go:noescape
func CopyToken(t xml.Token) xml.Token

//go:linkname encoderencodeelement xml.sub_encoderencodeelement
func encoderencodeelement(enc *xml.Encoder, v interface{}, start xml.StartElement) error {
	return enc.EncodeElement(v, start)
}

//go:linkname EncoderEncodeElement xml.sub_encoderencodeelement
//go:noescape
func EncoderEncodeElement(enc *xml.Encoder, v interface{}, start xml.StartElement) error

//go:linkname startelementcopy xml.sub_startelementcopy
func startelementcopy(e xml.StartElement) xml.StartElement {
	return e.Copy()
}

//go:linkname StartElementCopy xml.sub_startelementcopy
//go:noescape
func StartElementCopy(e xml.StartElement) xml.StartElement

//go:linkname syntaxerrorerror xml.sub_syntaxerrorerror
func syntaxerrorerror(e *xml.SyntaxError) string {
	return e.Error()
}

//go:linkname SyntaxErrorError xml.sub_syntaxerrorerror
//go:noescape
func SyntaxErrorError(e *xml.SyntaxError) string

//go:linkname unsupportedtypeerrorerror xml.sub_unsupportedtypeerrorerror
func unsupportedtypeerrorerror(e *xml.UnsupportedTypeError) string {
	return e.Error()
}

//go:linkname UnsupportedTypeErrorError xml.sub_unsupportedtypeerrorerror
//go:noescape
func UnsupportedTypeErrorError(e *xml.UnsupportedTypeError) string

//go:linkname Unmarshal encoding/xml.Unmarshal
//go:noescape
func Unmarshal(data []byte, v interface{}) error

//go:linkname commentcopy xml.sub_commentcopy
func commentcopy(c xml.Comment) xml.Comment {
	return c.Copy()
}

//go:linkname CommentCopy xml.sub_commentcopy
//go:noescape
func CommentCopy(c xml.Comment) xml.Comment

//go:linkname decoderdecode xml.sub_decoderdecode
func decoderdecode(d *xml.Decoder, v interface{}) error {
	return d.Decode(v)
}

//go:linkname DecoderDecode xml.sub_decoderdecode
//go:noescape
func DecoderDecode(d *xml.Decoder, v interface{}) error

//go:linkname decoderskip xml.sub_decoderskip
func decoderskip(d *xml.Decoder) error {
	return d.Skip()
}

//go:linkname DecoderSkip xml.sub_decoderskip
//go:noescape
func DecoderSkip(d *xml.Decoder) error

//go:linkname chardatacopy xml.sub_chardatacopy
func chardatacopy(c xml.CharData) xml.CharData {
	return c.Copy()
}

//go:linkname CharDataCopy xml.sub_chardatacopy
//go:noescape
func CharDataCopy(c xml.CharData) xml.CharData

//go:linkname encoderencodetoken xml.sub_encoderencodetoken
func encoderencodetoken(enc *xml.Encoder, t xml.Token) error {
	return enc.EncodeToken(t)
}

//go:linkname EncoderEncodeToken xml.sub_encoderencodetoken
//go:noescape
func EncoderEncodeToken(enc *xml.Encoder, t xml.Token) error

//go:linkname encoderflush xml.sub_encoderflush
func encoderflush(enc *xml.Encoder) error {
	return enc.Flush()
}

//go:linkname EncoderFlush xml.sub_encoderflush
//go:noescape
func EncoderFlush(enc *xml.Encoder) error

//go:linkname NewDecoder encoding/xml.NewDecoder
//go:noescape
func NewDecoder(r io.Reader) *xml.Decoder

//go:linkname Marshal encoding/xml.Marshal
//go:noescape
func Marshal(v interface{}) ([]byte, error)

//go:linkname decoderrawtoken xml.sub_decoderrawtoken
func decoderrawtoken(d *xml.Decoder) (xml.Token, error) {
	return d.RawToken()
}

//go:linkname DecoderRawToken xml.sub_decoderrawtoken
//go:noescape
func DecoderRawToken(d *xml.Decoder) (xml.Token, error)

//go:linkname unmarshalerrorerror xml.sub_unmarshalerrorerror
func unmarshalerrorerror(e xml.UnmarshalError) string {
	return e.Error()
}

//go:linkname UnmarshalErrorError xml.sub_unmarshalerrorerror
//go:noescape
func UnmarshalErrorError(e xml.UnmarshalError) string

//go:linkname EscapeText encoding/xml.EscapeText
//go:noescape
func EscapeText(w io.Writer, s []byte) error

//go:linkname NewTokenDecoder encoding/xml.NewTokenDecoder
//go:noescape
func NewTokenDecoder(t xml.TokenReader) *xml.Decoder

//go:linkname NewEncoder encoding/xml.NewEncoder
//go:noescape
func NewEncoder(w io.Writer) *xml.Encoder

//go:linkname tagpatherrorerror xml.sub_tagpatherrorerror
func tagpatherrorerror(e *xml.TagPathError) string {
	return e.Error()
}

//go:linkname TagPathErrorError xml.sub_tagpatherrorerror
//go:noescape
func TagPathErrorError(e *xml.TagPathError) string
