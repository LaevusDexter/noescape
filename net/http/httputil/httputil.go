// This file has automatically been generated on Wed Feb 26 15:50:44 +05 2020.
// DO NOT EDIT.
package httputil

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	_ "unsafe"
)

//go:linkname clientconnpending httputil.sub_clientconnpending
func clientconnpending(cc *httputil.ClientConn,) int {
	return cc.Pending()
}

//go:linkname ClientConnPending httputil.sub_clientconnpending
//go:noescape
func ClientConnPending(cc *httputil.ClientConn,) int

//go:linkname serverconnclose httputil.sub_serverconnclose
func serverconnclose(sc *httputil.ServerConn,) error {
	return sc.Close()
}

//go:linkname ServerConnClose httputil.sub_serverconnclose
//go:noescape
func ServerConnClose(sc *httputil.ServerConn,) error

//go:linkname serverconnhijack httputil.sub_serverconnhijack
func serverconnhijack(sc *httputil.ServerConn,) (net.Conn, *bufio.Reader) {
	return sc.Hijack()
}

//go:linkname ServerConnHijack httputil.sub_serverconnhijack
//go:noescape
func ServerConnHijack(sc *httputil.ServerConn,) (net.Conn, *bufio.Reader)

//go:linkname NewChunkedReader net/http/httputil.NewChunkedReader
//go:noescape
func NewChunkedReader(r io.Reader) io.Reader

//go:linkname NewProxyClientConn net/http/httputil.NewProxyClientConn
//go:noescape
func NewProxyClientConn(c net.Conn, r *bufio.Reader) *httputil.ClientConn

//go:linkname clientconndo httputil.sub_clientconndo
func clientconndo(cc *httputil.ClientConn, req *http.Request) (*http.Response, error) {
	return cc.Do(req)
}

//go:linkname ClientConnDo httputil.sub_clientconndo
//go:noescape
func ClientConnDo(cc *httputil.ClientConn, req *http.Request) (*http.Response, error)

//go:linkname clientconnhijack httputil.sub_clientconnhijack
func clientconnhijack(cc *httputil.ClientConn,) (net.Conn, *bufio.Reader) {
	return cc.Hijack()
}

//go:linkname ClientConnHijack httputil.sub_clientconnhijack
//go:noescape
func ClientConnHijack(cc *httputil.ClientConn,) (net.Conn, *bufio.Reader)

//go:linkname NewServerConn net/http/httputil.NewServerConn
//go:noescape
func NewServerConn(c net.Conn, r *bufio.Reader) *httputil.ServerConn

//go:linkname serverconnpending httputil.sub_serverconnpending
func serverconnpending(sc *httputil.ServerConn,) int {
	return sc.Pending()
}

//go:linkname ServerConnPending httputil.sub_serverconnpending
//go:noescape
func ServerConnPending(sc *httputil.ServerConn,) int

//go:linkname DumpResponse net/http/httputil.DumpResponse
//go:noescape
func DumpResponse(resp *http.Response, body bool) ([]byte, error)

//go:linkname NewChunkedWriter net/http/httputil.NewChunkedWriter
//go:noescape
func NewChunkedWriter(w io.Writer) io.WriteCloser

//go:linkname NewClientConn net/http/httputil.NewClientConn
//go:noescape
func NewClientConn(c net.Conn, r *bufio.Reader) *httputil.ClientConn

//go:linkname serverconnread httputil.sub_serverconnread
func serverconnread(sc *httputil.ServerConn,) (*http.Request, error) {
	return sc.Read()
}

//go:linkname ServerConnRead httputil.sub_serverconnread
//go:noescape
func ServerConnRead(sc *httputil.ServerConn,) (*http.Request, error)

//go:linkname serverconnwrite httputil.sub_serverconnwrite
func serverconnwrite(sc *httputil.ServerConn, req *http.Request, resp *http.Response) error {
	return sc.Write(req, resp)
}

//go:linkname ServerConnWrite httputil.sub_serverconnwrite
//go:noescape
func ServerConnWrite(sc *httputil.ServerConn, req *http.Request, resp *http.Response) error

//go:linkname DumpRequestOut net/http/httputil.DumpRequestOut
//go:noescape
func DumpRequestOut(req *http.Request, body bool) ([]byte, error)

//go:linkname clientconnread httputil.sub_clientconnread
func clientconnread(cc *httputil.ClientConn, req *http.Request) (*http.Response, error) {
	return cc.Read(req)
}

//go:linkname ClientConnRead httputil.sub_clientconnread
//go:noescape
func ClientConnRead(cc *httputil.ClientConn, req *http.Request) (*http.Response, error)

//go:linkname NewSingleHostReverseProxy net/http/httputil.NewSingleHostReverseProxy
//go:noescape
func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy

//go:linkname DumpRequest net/http/httputil.DumpRequest
//go:noescape
func DumpRequest(req *http.Request, body bool) ([]byte, error)

//go:linkname clientconnclose httputil.sub_clientconnclose
func clientconnclose(cc *httputil.ClientConn,) error {
	return cc.Close()
}

//go:linkname ClientConnClose httputil.sub_clientconnclose
//go:noescape
func ClientConnClose(cc *httputil.ClientConn,) error

//go:linkname clientconnwrite httputil.sub_clientconnwrite
func clientconnwrite(cc *httputil.ClientConn, req *http.Request) error {
	return cc.Write(req)
}

//go:linkname ClientConnWrite httputil.sub_clientconnwrite
//go:noescape
func ClientConnWrite(cc *httputil.ClientConn, req *http.Request) error
