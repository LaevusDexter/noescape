// This file has automatically been generated on Wed Feb 26 15:50:45 +05 2020.
// DO NOT EDIT.
package mail

import (
	"io"
	"net/mail"
	"time"
	_ "unsafe"
)

//go:linkname headeraddresslist mail.sub_headeraddresslist
func headeraddresslist(h mail.Header, key string) ([]*mail.Address, error) {
	return h.AddressList(key)
}

//go:linkname HeaderAddressList mail.sub_headeraddresslist
//go:noescape
func HeaderAddressList(h mail.Header, key string) ([]*mail.Address, error)

//go:linkname headerdate mail.sub_headerdate
func headerdate(h mail.Header,) (time.Time, error) {
	return h.Date()
}

//go:linkname HeaderDate mail.sub_headerdate
//go:noescape
func HeaderDate(h mail.Header,) (time.Time, error)

//go:linkname ReadMessage net/mail.ReadMessage
//go:noescape
func ReadMessage(r io.Reader) (*mail.Message, error)

//go:linkname ParseDate net/mail.ParseDate
//go:noescape
func ParseDate(date string) (time.Time, error)

//go:linkname ParseAddress net/mail.ParseAddress
//go:noescape
func ParseAddress(address string) (*mail.Address, error)

//go:linkname ParseAddressList net/mail.ParseAddressList
//go:noescape
func ParseAddressList(list string) ([]*mail.Address, error)

//go:linkname addressstring mail.sub_addressstring
func addressstring(a *mail.Address,) string {
	return a.String()
}

//go:linkname AddressString mail.sub_addressstring
//go:noescape
func AddressString(a *mail.Address,) string

//go:linkname addressparserparse mail.sub_addressparserparse
func addressparserparse(p *mail.AddressParser, address string) (*mail.Address, error) {
	return p.Parse(address)
}

//go:linkname AddressParserParse mail.sub_addressparserparse
//go:noescape
func AddressParserParse(p *mail.AddressParser, address string) (*mail.Address, error)

//go:linkname addressparserparselist mail.sub_addressparserparselist
func addressparserparselist(p *mail.AddressParser, list string) ([]*mail.Address, error) {
	return p.ParseList(list)
}

//go:linkname AddressParserParseList mail.sub_addressparserparselist
//go:noescape
func AddressParserParseList(p *mail.AddressParser, list string) ([]*mail.Address, error)

//go:linkname headerget mail.sub_headerget
func headerget(h mail.Header, key string) string {
	return h.Get(key)
}

//go:linkname HeaderGet mail.sub_headerget
//go:noescape
func HeaderGet(h mail.Header, key string) string
