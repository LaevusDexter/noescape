// This file has automatically been generated on Wed Feb 26 15:50:44 +05 2020.
// DO NOT EDIT.
package httptest

import (
	"crypto/x509"
	"io"
	"net/http"
	"net/http/httptest"
	_ "unsafe"
)

//go:linkname NewTLSServer net/http/httptest.NewTLSServer
//go:noescape
func NewTLSServer(handler http.Handler) *httptest.Server

//go:linkname NewUnstartedServer net/http/httptest.NewUnstartedServer
//go:noescape
func NewUnstartedServer(handler http.Handler) *httptest.Server

//go:linkname NewRequest net/http/httptest.NewRequest
//go:noescape
func NewRequest(method, target string, body io.Reader) *http.Request

//go:linkname NewRecorder net/http/httptest.NewRecorder
//go:noescape
func NewRecorder() *httptest.ResponseRecorder

//go:linkname responserecorderheader httptest.sub_responserecorderheader
func responserecorderheader(rw *httptest.ResponseRecorder,) http.Header {
	return rw.Header()
}

//go:linkname ResponseRecorderHeader httptest.sub_responserecorderheader
//go:noescape
func ResponseRecorderHeader(rw *httptest.ResponseRecorder,) http.Header

//go:linkname responserecorderresult httptest.sub_responserecorderresult
func responserecorderresult(rw *httptest.ResponseRecorder,) *http.Response {
	return rw.Result()
}

//go:linkname ResponseRecorderResult httptest.sub_responserecorderresult
//go:noescape
func ResponseRecorderResult(rw *httptest.ResponseRecorder,) *http.Response

//go:linkname responserecorderwrite httptest.sub_responserecorderwrite
func responserecorderwrite(rw *httptest.ResponseRecorder, buf []byte) (int, error) {
	return rw.Write(buf)
}

//go:linkname ResponseRecorderWrite httptest.sub_responserecorderwrite
//go:noescape
func ResponseRecorderWrite(rw *httptest.ResponseRecorder, buf []byte) (int, error)

//go:linkname NewServer net/http/httptest.NewServer
//go:noescape
func NewServer(handler http.Handler) *httptest.Server

//go:linkname serverclient httptest.sub_serverclient
func serverclient(s *httptest.Server,) *http.Client {
	return s.Client()
}

//go:linkname ServerClient httptest.sub_serverclient
//go:noescape
func ServerClient(s *httptest.Server,) *http.Client

//go:linkname responserecorderwritestring httptest.sub_responserecorderwritestring
func responserecorderwritestring(rw *httptest.ResponseRecorder, str string) (int, error) {
	return rw.WriteString(str)
}

//go:linkname ResponseRecorderWriteString httptest.sub_responserecorderwritestring
//go:noescape
func ResponseRecorderWriteString(rw *httptest.ResponseRecorder, str string) (int, error)

//go:linkname servercertificate httptest.sub_servercertificate
func servercertificate(s *httptest.Server,) *x509.Certificate {
	return s.Certificate()
}

//go:linkname ServerCertificate httptest.sub_servercertificate
//go:noescape
func ServerCertificate(s *httptest.Server,) *x509.Certificate
