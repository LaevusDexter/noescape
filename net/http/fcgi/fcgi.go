// This file has automatically been generated on Wed Feb 26 15:50:44 +05 2020.
// DO NOT EDIT.
package fcgi

import (
	"net"
	"net/http"
	_ "net/http/fcgi"
	_ "unsafe"
)

//go:linkname ProcessEnv net/http/fcgi.ProcessEnv
//go:noescape
func ProcessEnv(r *http.Request) map[string]string

//go:linkname Serve net/http/fcgi.Serve
//go:noescape
func Serve(l net.Listener, handler http.Handler) error
