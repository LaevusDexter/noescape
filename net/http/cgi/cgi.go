// This file has automatically been generated on Wed Feb 26 02:10:06 +05 2020.
// DO NOT EDIT.
package cgi

import (
	"net/http"
	_ "net/http/cgi"
	_ "unsafe"
)

//go:linkname Request net/http/cgi.Request
//go:noescape
func Request() (*http.Request, error)

//go:linkname RequestFromMap net/http/cgi.RequestFromMap
//go:noescape
func RequestFromMap(params map[string]string) (*http.Request, error)

//go:linkname Serve net/http/cgi.Serve
//go:noescape
func Serve(handler http.Handler) error
