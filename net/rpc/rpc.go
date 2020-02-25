// This file has automatically been generated on Wed Feb 26 02:10:08 +05 2020.
// DO NOT EDIT.
package rpc

import (
	"io"
	"net/rpc"
	_ "unsafe"
)

//go:linkname serverserverequest rpc.sub_serverserverequest
func serverserverequest(server *rpc.Server, codec rpc.ServerCodec) error {
	return server.ServeRequest(codec)
}

//go:linkname ServerServeRequest rpc.sub_serverserverequest
//go:noescape
func ServerServeRequest(server *rpc.Server, codec rpc.ServerCodec) error

//go:linkname Dial net/rpc.Dial
//go:noescape
func Dial(network, address string) (*rpc.Client, error)

//go:linkname NewClient net/rpc.NewClient
//go:noescape
func NewClient(conn io.ReadWriteCloser) *rpc.Client

//go:linkname NewClientWithCodec net/rpc.NewClientWithCodec
//go:noescape
func NewClientWithCodec(codec rpc.ClientCodec) *rpc.Client

//go:linkname clientgo rpc.sub_clientgo
func clientgo(client *rpc.Client, serviceMethod string, args interface{}, reply interface{}, done chan *rpc.Call) *rpc.Call {
	return client.Go(serviceMethod, args, reply, done)
}

//go:linkname ClientGo rpc.sub_clientgo
//go:noescape
func ClientGo(client *rpc.Client, serviceMethod string, args interface{}, reply interface{}, done chan *rpc.Call) *rpc.Call

//go:linkname DialHTTP net/rpc.DialHTTP
//go:noescape
func DialHTTP(network, address string) (*rpc.Client, error)

//go:linkname serverregistername rpc.sub_serverregistername
func serverregistername(server *rpc.Server, name string, rcvr interface{}) error {
	return server.RegisterName(name, rcvr)
}

//go:linkname ServerRegisterName rpc.sub_serverregistername
//go:noescape
func ServerRegisterName(server *rpc.Server, name string, rcvr interface{}) error

//go:linkname servererrorerror rpc.sub_servererrorerror
func servererrorerror(e rpc.ServerError) string {
	return e.Error()
}

//go:linkname ServerErrorError rpc.sub_servererrorerror
//go:noescape
func ServerErrorError(e rpc.ServerError) string

//go:linkname RegisterName net/rpc.RegisterName
//go:noescape
func RegisterName(name string, rcvr interface{}) error

//go:linkname DialHTTPPath net/rpc.DialHTTPPath
//go:noescape
func DialHTTPPath(network, address, path string) (*rpc.Client, error)

//go:linkname NewServer net/rpc.NewServer
//go:noescape
func NewServer() *rpc.Server

//go:linkname serverregister rpc.sub_serverregister
func serverregister(server *rpc.Server, rcvr interface{}) error {
	return server.Register(rcvr)
}

//go:linkname ServerRegister rpc.sub_serverregister
//go:noescape
func ServerRegister(server *rpc.Server, rcvr interface{}) error

//go:linkname Register net/rpc.Register
//go:noescape
func Register(rcvr interface{}) error

//go:linkname ServeRequest net/rpc.ServeRequest
//go:noescape
func ServeRequest(codec rpc.ServerCodec) error

//go:linkname clientcall rpc.sub_clientcall
func clientcall(client *rpc.Client, serviceMethod string, args interface{}, reply interface{}) error {
	return client.Call(serviceMethod, args, reply)
}

//go:linkname ClientCall rpc.sub_clientcall
//go:noescape
func ClientCall(client *rpc.Client, serviceMethod string, args interface{}, reply interface{}) error

//go:linkname clientclose rpc.sub_clientclose
func clientclose(client *rpc.Client) error {
	return client.Close()
}

//go:linkname ClientClose rpc.sub_clientclose
//go:noescape
func ClientClose(client *rpc.Client) error
