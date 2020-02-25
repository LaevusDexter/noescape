// This file has automatically been generated on Wed Feb 26 02:09:47 +05 2020.
// DO NOT EDIT.
package tls

import (
	"crypto/tls"
	"net"
	"time"
	_ "unsafe"
)

//go:linkname connclose tls.sub_connclose
func connclose(c *tls.Conn) error {
	return c.Close()
}

//go:linkname ConnClose tls.sub_connclose
//go:noescape
func ConnClose(c *tls.Conn) error

//go:linkname connconnectionstate tls.sub_connconnectionstate
func connconnectionstate(c *tls.Conn) tls.ConnectionState {
	return c.ConnectionState()
}

//go:linkname ConnConnectionState tls.sub_connconnectionstate
//go:noescape
func ConnConnectionState(c *tls.Conn) tls.ConnectionState

//go:linkname certificaterequestinfosupportscertificate tls.sub_certificaterequestinfosupportscertificate
func certificaterequestinfosupportscertificate(cri *tls.CertificateRequestInfo, c *tls.Certificate) error {
	return cri.SupportsCertificate(c)
}

//go:linkname CertificateRequestInfoSupportsCertificate tls.sub_certificaterequestinfosupportscertificate
//go:noescape
func CertificateRequestInfoSupportsCertificate(cri *tls.CertificateRequestInfo, c *tls.Certificate) error

//go:linkname DialWithDialer crypto/tls.DialWithDialer
//go:noescape
func DialWithDialer(dialer *net.Dialer, network, addr string, config *tls.Config) (*tls.Conn, error)

//go:linkname Client crypto/tls.Client
//go:noescape
func Client(conn net.Conn, config *tls.Config) *tls.Conn

//go:linkname CipherSuiteName crypto/tls.CipherSuiteName
//go:noescape
func CipherSuiteName(id uint16) string

//go:linkname NewListener crypto/tls.NewListener
//go:noescape
func NewListener(inner net.Listener, config *tls.Config) net.Listener

//go:linkname Server crypto/tls.Server
//go:noescape
func Server(conn net.Conn, config *tls.Config) *tls.Conn

//go:linkname connhandshake tls.sub_connhandshake
func connhandshake(c *tls.Conn) error {
	return c.Handshake()
}

//go:linkname ConnHandshake tls.sub_connhandshake
//go:noescape
func ConnHandshake(c *tls.Conn) error

//go:linkname connocspresponse tls.sub_connocspresponse
func connocspresponse(c *tls.Conn) []byte {
	return c.OCSPResponse()
}

//go:linkname ConnOCSPResponse tls.sub_connocspresponse
//go:noescape
func ConnOCSPResponse(c *tls.Conn) []byte

//go:linkname connectionstateexportkeyingmaterial tls.sub_connectionstateexportkeyingmaterial
func connectionstateexportkeyingmaterial(cs *tls.ConnectionState, label string, context []byte, length int) ([]byte, error) {
	return cs.ExportKeyingMaterial(label, context, length)
}

//go:linkname ConnectionStateExportKeyingMaterial tls.sub_connectionstateexportkeyingmaterial
//go:noescape
func ConnectionStateExportKeyingMaterial(cs *tls.ConnectionState, label string, context []byte, length int) ([]byte, error)

//go:linkname Listen crypto/tls.Listen
//go:noescape
func Listen(network, laddr string, config *tls.Config) (net.Listener, error)

//go:linkname configclone tls.sub_configclone
func configclone(c *tls.Config) *tls.Config {
	return c.Clone()
}

//go:linkname ConfigClone tls.sub_configclone
//go:noescape
func ConfigClone(c *tls.Config) *tls.Config

//go:linkname connsetwritedeadline tls.sub_connsetwritedeadline
func connsetwritedeadline(c *tls.Conn, t time.Time) error {
	return c.SetWriteDeadline(t)
}

//go:linkname ConnSetWriteDeadline tls.sub_connsetwritedeadline
//go:noescape
func ConnSetWriteDeadline(c *tls.Conn, t time.Time) error

//go:linkname InsecureCipherSuites crypto/tls.InsecureCipherSuites
//go:noescape
func InsecureCipherSuites() []*tls.CipherSuite

//go:linkname clienthelloinfosupportscertificate tls.sub_clienthelloinfosupportscertificate
func clienthelloinfosupportscertificate(chi *tls.ClientHelloInfo, c *tls.Certificate) error {
	return chi.SupportsCertificate(c)
}

//go:linkname ClientHelloInfoSupportsCertificate tls.sub_clienthelloinfosupportscertificate
//go:noescape
func ClientHelloInfoSupportsCertificate(chi *tls.ClientHelloInfo, c *tls.Certificate) error

//go:linkname connlocaladdr tls.sub_connlocaladdr
func connlocaladdr(c *tls.Conn) net.Addr {
	return c.LocalAddr()
}

//go:linkname ConnLocalAddr tls.sub_connlocaladdr
//go:noescape
func ConnLocalAddr(c *tls.Conn) net.Addr

//go:linkname connsetdeadline tls.sub_connsetdeadline
func connsetdeadline(c *tls.Conn, t time.Time) error {
	return c.SetDeadline(t)
}

//go:linkname ConnSetDeadline tls.sub_connsetdeadline
//go:noescape
func ConnSetDeadline(c *tls.Conn, t time.Time) error

//go:linkname recordheadererrorerror tls.sub_recordheadererrorerror
func recordheadererrorerror(e tls.RecordHeaderError) string {
	return e.Error()
}

//go:linkname RecordHeaderErrorError tls.sub_recordheadererrorerror
//go:noescape
func RecordHeaderErrorError(e tls.RecordHeaderError) string

//go:linkname NewLRUClientSessionCache crypto/tls.NewLRUClientSessionCache
//go:noescape
func NewLRUClientSessionCache(capacity int) tls.ClientSessionCache

//go:linkname connclosewrite tls.sub_connclosewrite
func connclosewrite(c *tls.Conn) error {
	return c.CloseWrite()
}

//go:linkname ConnCloseWrite tls.sub_connclosewrite
//go:noescape
func ConnCloseWrite(c *tls.Conn) error

//go:linkname connverifyhostname tls.sub_connverifyhostname
func connverifyhostname(c *tls.Conn, host string) error {
	return c.VerifyHostname(host)
}

//go:linkname ConnVerifyHostname tls.sub_connverifyhostname
//go:noescape
func ConnVerifyHostname(c *tls.Conn, host string) error

//go:linkname connwrite tls.sub_connwrite
func connwrite(c *tls.Conn, b []byte) (int, error) {
	return c.Write(b)
}

//go:linkname ConnWrite tls.sub_connwrite
//go:noescape
func ConnWrite(c *tls.Conn, b []byte) (int, error)

//go:linkname Dial crypto/tls.Dial
//go:noescape
func Dial(network, addr string, config *tls.Config) (*tls.Conn, error)

//go:linkname connremoteaddr tls.sub_connremoteaddr
func connremoteaddr(c *tls.Conn) net.Addr {
	return c.RemoteAddr()
}

//go:linkname ConnRemoteAddr tls.sub_connremoteaddr
//go:noescape
func ConnRemoteAddr(c *tls.Conn) net.Addr

//go:linkname connsetreaddeadline tls.sub_connsetreaddeadline
func connsetreaddeadline(c *tls.Conn, t time.Time) error {
	return c.SetReadDeadline(t)
}

//go:linkname ConnSetReadDeadline tls.sub_connsetreaddeadline
//go:noescape
func ConnSetReadDeadline(c *tls.Conn, t time.Time) error

//go:linkname CipherSuites crypto/tls.CipherSuites
//go:noescape
func CipherSuites() []*tls.CipherSuite

//go:linkname connread tls.sub_connread
func connread(c *tls.Conn, b []byte) (int, error) {
	return c.Read(b)
}

//go:linkname ConnRead tls.sub_connread
//go:noescape
func ConnRead(c *tls.Conn, b []byte) (int, error)
