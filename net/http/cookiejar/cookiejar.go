// This file has automatically been generated on Wed Feb 26 15:50:44 +05 2020.
// DO NOT EDIT.
package cookiejar

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	_ "unsafe"
)

//go:linkname jarcookies cookiejar.sub_jarcookies
func jarcookies(j *cookiejar.Jar, u *url.URL) []*http.Cookie {
	return j.Cookies(u)
}

//go:linkname JarCookies cookiejar.sub_jarcookies
//go:noescape
func JarCookies(j *cookiejar.Jar, u *url.URL) []*http.Cookie

//go:linkname New net/http/cookiejar.New
//go:noescape
func New(o *cookiejar.Options,) (*cookiejar.Jar, error)
