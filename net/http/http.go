// This file has automatically been generated on Wed Feb 26 02:10:06 +05 2020.
// DO NOT EDIT.
package http

import (
	"bufio"
	"context"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"time"
	_ "unsafe"
)

//go:linkname ParseTime net/http.ParseTime
//go:noescape
func ParseTime(text string) (time.Time, error)

//go:linkname ProxyFromEnvironment net/http.ProxyFromEnvironment
//go:noescape
func ProxyFromEnvironment(req *http.Request) (*url.URL, error)

//go:linkname requestreferer http.sub_requestreferer
func requestreferer(r *http.Request) string {
	return r.Referer()
}

//go:linkname RequestReferer http.sub_requestreferer
//go:noescape
func RequestReferer(r *http.Request) string

//go:linkname responselocation http.sub_responselocation
func responselocation(r *http.Response) (*url.URL, error) {
	return r.Location()
}

//go:linkname ResponseLocation http.sub_responselocation
//go:noescape
func ResponseLocation(r *http.Response) (*url.URL, error)

//go:linkname responsewrite http.sub_responsewrite
func responsewrite(r *http.Response, w io.Writer) error {
	return r.Write(w)
}

//go:linkname ResponseWrite http.sub_responsewrite
//go:noescape
func ResponseWrite(r *http.Response, w io.Writer) error

//go:linkname connstatestring http.sub_connstatestring
func connstatestring(c http.ConnState) string {
	return c.String()
}

//go:linkname ConnStateString http.sub_connstatestring
//go:noescape
func ConnStateString(c http.ConnState) string

//go:linkname requestwrite http.sub_requestwrite
func requestwrite(r *http.Request, w io.Writer) error {
	return r.Write(w)
}

//go:linkname RequestWrite http.sub_requestwrite
//go:noescape
func RequestWrite(r *http.Request, w io.Writer) error

//go:linkname ReadResponse net/http.ReadResponse
//go:noescape
func ReadResponse(r *bufio.Reader, req *http.Request) (*http.Response, error)

//go:linkname serverlistenandservetls http.sub_serverlistenandservetls
func serverlistenandservetls(srv *http.Server, certFile, keyFile string) error {
	return srv.ListenAndServeTLS(certFile, keyFile)
}

//go:linkname ServerListenAndServeTLS http.sub_serverlistenandservetls
//go:noescape
func ServerListenAndServeTLS(srv *http.Server, certFile, keyFile string) error

//go:linkname headerget http.sub_headerget
func headerget(h http.Header, key string) string {
	return h.Get(key)
}

//go:linkname HeaderGet http.sub_headerget
//go:noescape
func HeaderGet(h http.Header, key string) string

//go:linkname Head net/http.Head
//go:noescape
func Head(url string) (*http.Response, error)

//go:linkname MaxBytesReader net/http.MaxBytesReader
//go:noescape
func MaxBytesReader(w http.ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser

//go:linkname Serve net/http.Serve
//go:noescape
func Serve(l net.Listener, handler http.Handler) error

//go:linkname clienthead http.sub_clienthead
func clienthead(c *http.Client, url string) (*http.Response, error) {
	return c.Head(url)
}

//go:linkname ClientHead http.sub_clienthead
//go:noescape
func ClientHead(c *http.Client, url string) (*http.Response, error)

//go:linkname requestcookie http.sub_requestcookie
func requestcookie(r *http.Request, name string) (*http.Cookie, error) {
	return r.Cookie(name)
}

//go:linkname RequestCookie http.sub_requestcookie
//go:noescape
func RequestCookie(r *http.Request, name string) (*http.Cookie, error)

//go:linkname requestprotoatleast http.sub_requestprotoatleast
func requestprotoatleast(r *http.Request, major, minor int) bool {
	return r.ProtoAtLeast(major, minor)
}

//go:linkname RequestProtoAtLeast http.sub_requestprotoatleast
//go:noescape
func RequestProtoAtLeast(r *http.Request, major, minor int) bool

//go:linkname ServeTLS net/http.ServeTLS
//go:noescape
func ServeTLS(l net.Listener, handler http.Handler, certFile, keyFile string) error

//go:linkname clientget http.sub_clientget
func clientget(c *http.Client, url string) (*http.Response, error) {
	return c.Get(url)
}

//go:linkname ClientGet http.sub_clientget
//go:noescape
func ClientGet(c *http.Client, url string) (*http.Response, error)

//go:linkname clientpost http.sub_clientpost
func clientpost(c *http.Client, url, contentType string, body io.Reader) (*http.Response, error) {
	return c.Post(url, contentType, body)
}

//go:linkname ClientPost http.sub_clientpost
//go:noescape
func ClientPost(c *http.Client, url, contentType string, body io.Reader) (*http.Response, error)

//go:linkname headerclone http.sub_headerclone
func headerclone(h http.Header) http.Header {
	return h.Clone()
}

//go:linkname HeaderClone http.sub_headerclone
//go:noescape
func HeaderClone(h http.Header) http.Header

//go:linkname requestbasicauth http.sub_requestbasicauth
func requestbasicauth(r *http.Request) (string, bool) {
	return r.BasicAuth()
}

//go:linkname RequestBasicAuth http.sub_requestbasicauth
//go:noescape
func RequestBasicAuth(r *http.Request) (string, bool)

//go:linkname StatusText net/http.StatusText
//go:noescape
func StatusText(code int) string

//go:linkname requestcontext http.sub_requestcontext
func requestcontext(r *http.Request) context.Context {
	return r.Context()
}

//go:linkname RequestContext http.sub_requestcontext
//go:noescape
func RequestContext(r *http.Request) context.Context

//go:linkname requestpostformvalue http.sub_requestpostformvalue
func requestpostformvalue(r *http.Request, key string) string {
	return r.PostFormValue(key)
}

//go:linkname RequestPostFormValue http.sub_requestpostformvalue
//go:noescape
func RequestPostFormValue(r *http.Request, key string) string

//go:linkname serverlistenandserve http.sub_serverlistenandserve
func serverlistenandserve(srv *http.Server) error {
	return srv.ListenAndServe()
}

//go:linkname ServerListenAndServe http.sub_serverlistenandserve
//go:noescape
func ServerListenAndServe(srv *http.Server) error

//go:linkname transportclone http.sub_transportclone
func transportclone(t *http.Transport) *http.Transport {
	return t.Clone()
}

//go:linkname TransportClone http.sub_transportclone
//go:noescape
func TransportClone(t *http.Transport) *http.Transport

//go:linkname headervalues http.sub_headervalues
func headervalues(h http.Header, key string) []string {
	return h.Values(key)
}

//go:linkname HeaderValues http.sub_headervalues
//go:noescape
func HeaderValues(h http.Header, key string) []string

//go:linkname PostForm net/http.PostForm
//go:noescape
func PostForm(url string, data url.Values) (*http.Response, error)

//go:linkname NotFoundHandler net/http.NotFoundHandler
//go:noescape
func NotFoundHandler() http.Handler

//go:linkname headerwrite http.sub_headerwrite
func headerwrite(h http.Header, w io.Writer) error {
	return h.Write(w)
}

//go:linkname HeaderWrite http.sub_headerwrite
//go:noescape
func HeaderWrite(h http.Header, w io.Writer) error

//go:linkname protocolerrorerror http.sub_protocolerrorerror
func protocolerrorerror(pe *http.ProtocolError) string {
	return pe.Error()
}

//go:linkname ProtocolErrorError http.sub_protocolerrorerror
//go:noescape
func ProtocolErrorError(pe *http.ProtocolError) string

//go:linkname requestcookies http.sub_requestcookies
func requestcookies(r *http.Request) []*http.Cookie {
	return r.Cookies()
}

//go:linkname RequestCookies http.sub_requestcookies
//go:noescape
func RequestCookies(r *http.Request) []*http.Cookie

//go:linkname Get net/http.Get
//go:noescape
func Get(url string) (*http.Response, error)

//go:linkname clientpostform http.sub_clientpostform
func clientpostform(c *http.Client, url string, data url.Values) (*http.Response, error) {
	return c.PostForm(url, data)
}

//go:linkname ClientPostForm http.sub_clientpostform
//go:noescape
func ClientPostForm(c *http.Client, url string, data url.Values) (*http.Response, error)

//go:linkname diropen http.sub_diropen
func diropen(d http.Dir, name string) (http.File, error) {
	return d.Open(name)
}

//go:linkname DirOpen http.sub_diropen
//go:noescape
func DirOpen(d http.Dir, name string) (http.File, error)

//go:linkname requestformvalue http.sub_requestformvalue
func requestformvalue(r *http.Request, key string) string {
	return r.FormValue(key)
}

//go:linkname RequestFormValue http.sub_requestformvalue
//go:noescape
func RequestFormValue(r *http.Request, key string) string

//go:linkname servemuxhandler http.sub_servemuxhandler
func servemuxhandler(mux *http.ServeMux, r *http.Request) (http.Handler, string) {
	return mux.Handler(r)
}

//go:linkname ServeMuxHandler http.sub_servemuxhandler
//go:noescape
func ServeMuxHandler(mux *http.ServeMux, r *http.Request) (http.Handler, string)

//go:linkname requestwriteproxy http.sub_requestwriteproxy
func requestwriteproxy(r *http.Request, w io.Writer) error {
	return r.WriteProxy(w)
}

//go:linkname RequestWriteProxy http.sub_requestwriteproxy
//go:noescape
func RequestWriteProxy(r *http.Request, w io.Writer) error

//go:linkname NewServeMux net/http.NewServeMux
//go:noescape
func NewServeMux() *http.ServeMux

//go:linkname CanonicalHeaderKey net/http.CanonicalHeaderKey
//go:noescape
func CanonicalHeaderKey(s string) string

//go:linkname ListenAndServe net/http.ListenAndServe
//go:noescape
func ListenAndServe(addr string, handler http.Handler) error

//go:linkname headerwritesubset http.sub_headerwritesubset
func headerwritesubset(h http.Header, w io.Writer, exclude map[string]bool) error {
	return h.WriteSubset(w, exclude)
}

//go:linkname HeaderWriteSubset http.sub_headerwritesubset
//go:noescape
func HeaderWriteSubset(h http.Header, w io.Writer, exclude map[string]bool) error

//go:linkname ReadRequest net/http.ReadRequest
//go:noescape
func ReadRequest(b *bufio.Reader) (*http.Request, error)

//go:linkname requestmultipartreader http.sub_requestmultipartreader
func requestmultipartreader(r *http.Request) (*multipart.Reader, error) {
	return r.MultipartReader()
}

//go:linkname RequestMultipartReader http.sub_requestmultipartreader
//go:noescape
func RequestMultipartReader(r *http.Request) (*multipart.Reader, error)

//go:linkname ListenAndServeTLS net/http.ListenAndServeTLS
//go:noescape
func ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error

//go:linkname StripPrefix net/http.StripPrefix
//go:noescape
func StripPrefix(prefix string, h http.Handler) http.Handler

//go:linkname requestformfile http.sub_requestformfile
func requestformfile(r *http.Request, key string) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile(key)
}

//go:linkname RequestFormFile http.sub_requestformfile
//go:noescape
func RequestFormFile(r *http.Request, key string) (multipart.File, *multipart.FileHeader, error)

//go:linkname responsecookies http.sub_responsecookies
func responsecookies(r *http.Response) []*http.Cookie {
	return r.Cookies()
}

//go:linkname ResponseCookies http.sub_responsecookies
//go:noescape
func ResponseCookies(r *http.Response) []*http.Cookie

//go:linkname NewFileTransport net/http.NewFileTransport
//go:noescape
func NewFileTransport(fs http.FileSystem) http.RoundTripper

//go:linkname FileServer net/http.FileServer
//go:noescape
func FileServer(root http.FileSystem) http.Handler

//go:linkname NewRequest net/http.NewRequest
//go:noescape
func NewRequest(method, url string, body io.Reader) (*http.Request, error)

//go:linkname requestuseragent http.sub_requestuseragent
func requestuseragent(r *http.Request) string {
	return r.UserAgent()
}

//go:linkname RequestUserAgent http.sub_requestuseragent
//go:noescape
func RequestUserAgent(r *http.Request) string

//go:linkname requestwithcontext http.sub_requestwithcontext
func requestwithcontext(r *http.Request, ctx context.Context) *http.Request {
	return r.WithContext(ctx)
}

//go:linkname RequestWithContext http.sub_requestwithcontext
//go:noescape
func RequestWithContext(r *http.Request, ctx context.Context) *http.Request

//go:linkname serverservetls http.sub_serverservetls
func serverservetls(srv *http.Server, l net.Listener, certFile, keyFile string) error {
	return srv.ServeTLS(l, certFile, keyFile)
}

//go:linkname ServerServeTLS http.sub_serverservetls
//go:noescape
func ServerServeTLS(srv *http.Server, l net.Listener, certFile, keyFile string) error

//go:linkname responseprotoatleast http.sub_responseprotoatleast
func responseprotoatleast(r *http.Response, major, minor int) bool {
	return r.ProtoAtLeast(major, minor)
}

//go:linkname ResponseProtoAtLeast http.sub_responseprotoatleast
//go:noescape
func ResponseProtoAtLeast(r *http.Response, major, minor int) bool

//go:linkname serverclose http.sub_serverclose
func serverclose(srv *http.Server) error {
	return srv.Close()
}

//go:linkname ServerClose http.sub_serverclose
//go:noescape
func ServerClose(srv *http.Server) error

//go:linkname DetectContentType net/http.DetectContentType
//go:noescape
func DetectContentType(data []byte) string

//go:linkname cookiestring http.sub_cookiestring
func cookiestring(c *http.Cookie) string {
	return c.String()
}

//go:linkname CookieString http.sub_cookiestring
//go:noescape
func CookieString(c *http.Cookie) string

//go:linkname NewRequestWithContext net/http.NewRequestWithContext
//go:noescape
func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error)

//go:linkname requestparsemultipartform http.sub_requestparsemultipartform
func requestparsemultipartform(r *http.Request, maxMemory int64) error {
	return r.ParseMultipartForm(maxMemory)
}

//go:linkname RequestParseMultipartForm http.sub_requestparsemultipartform
//go:noescape
func RequestParseMultipartForm(r *http.Request, maxMemory int64) error

//go:linkname Post net/http.Post
//go:noescape
func Post(url, contentType string, body io.Reader) (*http.Response, error)

//go:linkname ParseHTTPVersion net/http.ParseHTTPVersion
//go:noescape
func ParseHTTPVersion(vers string) (int, bool)

//go:linkname RedirectHandler net/http.RedirectHandler
//go:noescape
func RedirectHandler(url string, code int) http.Handler

//go:linkname TimeoutHandler net/http.TimeoutHandler
//go:noescape
func TimeoutHandler(h http.Handler, dt time.Duration, msg string) http.Handler

//go:linkname requestparseform http.sub_requestparseform
func requestparseform(r *http.Request) error {
	return r.ParseForm()
}

//go:linkname RequestParseForm http.sub_requestparseform
//go:noescape
func RequestParseForm(r *http.Request) error

//go:linkname serverserve http.sub_serverserve
func serverserve(srv *http.Server, l net.Listener) error {
	return srv.Serve(l)
}

//go:linkname ServerServe http.sub_serverserve
//go:noescape
func ServerServe(srv *http.Server, l net.Listener) error

//go:linkname ProxyURL net/http.ProxyURL
//go:noescape
func ProxyURL(fixedURL *url.URL) func(*http.Request) (*url.URL, error)

//go:linkname clientdo http.sub_clientdo
func clientdo(c *http.Client, req *http.Request) (*http.Response, error) {
	return c.Do(req)
}

//go:linkname ClientDo http.sub_clientdo
//go:noescape
func ClientDo(c *http.Client, req *http.Request) (*http.Response, error)

//go:linkname requestclone http.sub_requestclone
func requestclone(r *http.Request, ctx context.Context) *http.Request {
	return r.Clone(ctx)
}

//go:linkname RequestClone http.sub_requestclone
//go:noescape
func RequestClone(r *http.Request, ctx context.Context) *http.Request

//go:linkname servershutdown http.sub_servershutdown
func servershutdown(srv *http.Server, ctx context.Context) error {
	return srv.Shutdown(ctx)
}

//go:linkname ServerShutdown http.sub_servershutdown
//go:noescape
func ServerShutdown(srv *http.Server, ctx context.Context) error

//go:linkname transportroundtrip http.sub_transportroundtrip
func transportroundtrip(t *http.Transport, req *http.Request) (*http.Response, error) {
	return t.RoundTrip(req)
}

//go:linkname TransportRoundTrip http.sub_transportroundtrip
//go:noescape
func TransportRoundTrip(t *http.Transport, req *http.Request) (*http.Response, error)
