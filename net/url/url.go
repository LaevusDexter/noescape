// This file has automatically been generated on Wed Feb 26 15:50:46 +05 2020.
// DO NOT EDIT.
package url

import (
	"net/url"
	_ "unsafe"
)

//go:linkname errortemporary url.sub_errortemporary
func errortemporary(e *url.Error,) bool {
	return e.Temporary()
}

//go:linkname ErrorTemporary url.sub_errortemporary
//go:noescape
func ErrorTemporary(e *url.Error,) bool

//go:linkname urlquery url.sub_urlquery
func urlquery(u *url.URL,) url.Values {
	return u.Query()
}

//go:linkname URLQuery url.sub_urlquery
//go:noescape
func URLQuery(u *url.URL,) url.Values

//go:linkname valuesget url.sub_valuesget
func valuesget(v url.Values, key string) string {
	return v.Get(key)
}

//go:linkname ValuesGet url.sub_valuesget
//go:noescape
func ValuesGet(v url.Values, key string) string

//go:linkname QueryEscape net/url.QueryEscape
//go:noescape
func QueryEscape(s string) string

//go:linkname userinfopassword url.sub_userinfopassword
func userinfopassword(u *url.Userinfo,) (string, bool) {
	return u.Password()
}

//go:linkname UserinfoPassword url.sub_userinfopassword
//go:noescape
func UserinfoPassword(u *url.Userinfo,) (string, bool)

//go:linkname valuesencode url.sub_valuesencode
func valuesencode(v url.Values,) string {
	return v.Encode()
}

//go:linkname ValuesEncode url.sub_valuesencode
//go:noescape
func ValuesEncode(v url.Values,) string

//go:linkname urlhostname url.sub_urlhostname
func urlhostname(u *url.URL,) string {
	return u.Hostname()
}

//go:linkname URLHostname url.sub_urlhostname
//go:noescape
func URLHostname(u *url.URL,) string

//go:linkname urlresolvereference url.sub_urlresolvereference
func urlresolvereference(u *url.URL, ref *url.URL,) *url.URL {
	return u.ResolveReference(ref)
}

//go:linkname URLResolveReference url.sub_urlresolvereference
//go:noescape
func URLResolveReference(u *url.URL, ref *url.URL,) *url.URL

//go:linkname userinfostring url.sub_userinfostring
func userinfostring(u *url.Userinfo,) string {
	return u.String()
}

//go:linkname UserinfoString url.sub_userinfostring
//go:noescape
func UserinfoString(u *url.Userinfo,) string

//go:linkname Parse net/url.Parse
//go:noescape
func Parse(rawurl string) (*url.URL, error)

//go:linkname urlparse url.sub_urlparse
func urlparse(u *url.URL, ref string) (*url.URL, error) {
	return u.Parse(ref)
}

//go:linkname URLParse url.sub_urlparse
//go:noescape
func URLParse(u *url.URL, ref string) (*url.URL, error)

//go:linkname urlport url.sub_urlport
func urlport(u *url.URL,) string {
	return u.Port()
}

//go:linkname URLPort url.sub_urlport
//go:noescape
func URLPort(u *url.URL,) string

//go:linkname urlrequesturi url.sub_urlrequesturi
func urlrequesturi(u *url.URL,) string {
	return u.RequestURI()
}

//go:linkname URLRequestURI url.sub_urlrequesturi
//go:noescape
func URLRequestURI(u *url.URL,) string

//go:linkname ParseQuery net/url.ParseQuery
//go:noescape
func ParseQuery(query string) (url.Values, error)

//go:linkname PathEscape net/url.PathEscape
//go:noescape
func PathEscape(s string) string

//go:linkname PathUnescape net/url.PathUnescape
//go:noescape
func PathUnescape(s string) (string, error)

//go:linkname QueryUnescape net/url.QueryUnescape
//go:noescape
func QueryUnescape(s string) (string, error)

//go:linkname errortimeout url.sub_errortimeout
func errortimeout(e *url.Error,) bool {
	return e.Timeout()
}

//go:linkname ErrorTimeout url.sub_errortimeout
//go:noescape
func ErrorTimeout(e *url.Error,) bool

//go:linkname urlmarshalbinary url.sub_urlmarshalbinary
func urlmarshalbinary(u *url.URL,) ([]byte, error) {
	return u.MarshalBinary()
}

//go:linkname URLMarshalBinary url.sub_urlmarshalbinary
//go:noescape
func URLMarshalBinary(u *url.URL,) ([]byte, error)

//go:linkname urlunmarshalbinary url.sub_urlunmarshalbinary
func urlunmarshalbinary(u *url.URL, text []byte) error {
	return u.UnmarshalBinary(text)
}

//go:linkname URLUnmarshalBinary url.sub_urlunmarshalbinary
//go:noescape
func URLUnmarshalBinary(u *url.URL, text []byte) error

//go:linkname escapeerrorerror url.sub_escapeerrorerror
func escapeerrorerror(e url.EscapeError,) string {
	return e.Error()
}

//go:linkname EscapeErrorError url.sub_escapeerrorerror
//go:noescape
func EscapeErrorError(e url.EscapeError,) string

//go:linkname User net/url.User
//go:noescape
func User(username string) *url.Userinfo

//go:linkname UserPassword net/url.UserPassword
//go:noescape
func UserPassword(username, password string) *url.Userinfo

//go:linkname userinfousername url.sub_userinfousername
func userinfousername(u *url.Userinfo,) string {
	return u.Username()
}

//go:linkname UserinfoUsername url.sub_userinfousername
//go:noescape
func UserinfoUsername(u *url.Userinfo,) string

//go:linkname errorerror url.sub_errorerror
func errorerror(e *url.Error,) string {
	return e.Error()
}

//go:linkname ErrorError url.sub_errorerror
//go:noescape
func ErrorError(e *url.Error,) string

//go:linkname errorunwrap url.sub_errorunwrap
func errorunwrap(e *url.Error,) error {
	return e.Unwrap()
}

//go:linkname ErrorUnwrap url.sub_errorunwrap
//go:noescape
func ErrorUnwrap(e *url.Error,) error

//go:linkname ParseRequestURI net/url.ParseRequestURI
//go:noescape
func ParseRequestURI(rawurl string) (*url.URL, error)

//go:linkname urlstring url.sub_urlstring
func urlstring(u *url.URL,) string {
	return u.String()
}

//go:linkname URLString url.sub_urlstring
//go:noescape
func URLString(u *url.URL,) string

//go:linkname invalidhosterrorerror url.sub_invalidhosterrorerror
func invalidhosterrorerror(e url.InvalidHostError,) string {
	return e.Error()
}

//go:linkname InvalidHostErrorError url.sub_invalidhosterrorerror
//go:noescape
func InvalidHostErrorError(e url.InvalidHostError,) string

//go:linkname urlescapedpath url.sub_urlescapedpath
func urlescapedpath(u *url.URL,) string {
	return u.EscapedPath()
}

//go:linkname URLEscapedPath url.sub_urlescapedpath
//go:noescape
func URLEscapedPath(u *url.URL,) string

//go:linkname urlisabs url.sub_urlisabs
func urlisabs(u *url.URL,) bool {
	return u.IsAbs()
}

//go:linkname URLIsAbs url.sub_urlisabs
//go:noescape
func URLIsAbs(u *url.URL,) bool
