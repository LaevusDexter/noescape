// This file has automatically been generated on Wed Feb 26 02:09:52 +05 2020.
// DO NOT EDIT.
package gob

import (
	"encoding/gob"
	"io"
	"reflect"
	_ "unsafe"
)

//go:linkname encoderencode gob.sub_encoderencode
func encoderencode(enc *gob.Encoder, e interface{}) error {
	return enc.Encode(e)
}

//go:linkname EncoderEncode gob.sub_encoderencode
//go:noescape
func EncoderEncode(enc *gob.Encoder, e interface{}) error

//go:linkname encoderencodevalue gob.sub_encoderencodevalue
func encoderencodevalue(enc *gob.Encoder, value reflect.Value) error {
	return enc.EncodeValue(value)
}

//go:linkname EncoderEncodeValue gob.sub_encoderencodevalue
//go:noescape
func EncoderEncodeValue(enc *gob.Encoder, value reflect.Value) error

//go:linkname NewDecoder encoding/gob.NewDecoder
//go:noescape
func NewDecoder(r io.Reader) *gob.Decoder

//go:linkname decoderdecode gob.sub_decoderdecode
func decoderdecode(dec *gob.Decoder, e interface{}) error {
	return dec.Decode(e)
}

//go:linkname DecoderDecode gob.sub_decoderdecode
//go:noescape
func DecoderDecode(dec *gob.Decoder, e interface{}) error

//go:linkname decoderdecodevalue gob.sub_decoderdecodevalue
func decoderdecodevalue(dec *gob.Decoder, v reflect.Value) error {
	return dec.DecodeValue(v)
}

//go:linkname DecoderDecodeValue gob.sub_decoderdecodevalue
//go:noescape
func DecoderDecodeValue(dec *gob.Decoder, v reflect.Value) error

//go:linkname NewEncoder encoding/gob.NewEncoder
//go:noescape
func NewEncoder(w io.Writer) *gob.Encoder
