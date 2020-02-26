// This file has automatically been generated on Wed Feb 26 15:50:45 +05 2020.
// DO NOT EDIT.
package jsonrpc

import (
	"io"
	"net/rpc"
	_ "net/rpc/jsonrpc"
	_ "unsafe"
)

//go:linkname Dial net/rpc/jsonrpc.Dial
//go:noescape
func Dial(network, address string) (*rpc.Client, error)

//go:linkname NewClient net/rpc/jsonrpc.NewClient
//go:noescape
func NewClient(conn io.ReadWriteCloser) *rpc.Client

//go:linkname NewClientCodec net/rpc/jsonrpc.NewClientCodec
//go:noescape
func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec

//go:linkname NewServerCodec net/rpc/jsonrpc.NewServerCodec
//go:noescape
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec
