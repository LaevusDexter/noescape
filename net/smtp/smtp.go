// This file has automatically been generated on Wed Feb 26 02:10:08 +05 2020.
// DO NOT EDIT.
package smtp

import (
	"crypto/tls"
	"io"
	"net"
	"net/smtp"
	_ "unsafe"
)

//go:linkname Dial net/smtp.Dial
//go:noescape
func Dial(addr string) (*smtp.Client, error)

//go:linkname clientclose smtp.sub_clientclose
func clientclose(c *smtp.Client) error {
	return c.Close()
}

//go:linkname ClientClose smtp.sub_clientclose
//go:noescape
func ClientClose(c *smtp.Client) error

//go:linkname clientmail smtp.sub_clientmail
func clientmail(c *smtp.Client, from string) error {
	return c.Mail(from)
}

//go:linkname ClientMail smtp.sub_clientmail
//go:noescape
func ClientMail(c *smtp.Client, from string) error

//go:linkname clientreset smtp.sub_clientreset
func clientreset(c *smtp.Client) error {
	return c.Reset()
}

//go:linkname ClientReset smtp.sub_clientreset
//go:noescape
func ClientReset(c *smtp.Client) error

//go:linkname clienttlsconnectionstate smtp.sub_clienttlsconnectionstate
func clienttlsconnectionstate(c *smtp.Client) (tls.ConnectionState, bool) {
	return c.TLSConnectionState()
}

//go:linkname ClientTLSConnectionState smtp.sub_clienttlsconnectionstate
//go:noescape
func ClientTLSConnectionState(c *smtp.Client) (tls.ConnectionState, bool)

//go:linkname SendMail net/smtp.SendMail
//go:noescape
func SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error

//go:linkname PlainAuth net/smtp.PlainAuth
//go:noescape
func PlainAuth(identity, username, password, host string) smtp.Auth

//go:linkname clientrcpt smtp.sub_clientrcpt
func clientrcpt(c *smtp.Client, to string) error {
	return c.Rcpt(to)
}

//go:linkname ClientRcpt smtp.sub_clientrcpt
//go:noescape
func ClientRcpt(c *smtp.Client, to string) error

//go:linkname clientstarttls smtp.sub_clientstarttls
func clientstarttls(c *smtp.Client, config *tls.Config) error {
	return c.StartTLS(config)
}

//go:linkname ClientStartTLS smtp.sub_clientstarttls
//go:noescape
func ClientStartTLS(c *smtp.Client, config *tls.Config) error

//go:linkname NewClient net/smtp.NewClient
//go:noescape
func NewClient(conn net.Conn, host string) (*smtp.Client, error)

//go:linkname clientdata smtp.sub_clientdata
func clientdata(c *smtp.Client) (io.WriteCloser, error) {
	return c.Data()
}

//go:linkname ClientData smtp.sub_clientdata
//go:noescape
func ClientData(c *smtp.Client) (io.WriteCloser, error)

//go:linkname clientextension smtp.sub_clientextension
func clientextension(c *smtp.Client, ext string) (bool, string) {
	return c.Extension(ext)
}

//go:linkname ClientExtension smtp.sub_clientextension
//go:noescape
func ClientExtension(c *smtp.Client, ext string) (bool, string)

//go:linkname clientverify smtp.sub_clientverify
func clientverify(c *smtp.Client, addr string) error {
	return c.Verify(addr)
}

//go:linkname ClientVerify smtp.sub_clientverify
//go:noescape
func ClientVerify(c *smtp.Client, addr string) error

//go:linkname clientauth smtp.sub_clientauth
func clientauth(c *smtp.Client, a smtp.Auth) error {
	return c.Auth(a)
}

//go:linkname ClientAuth smtp.sub_clientauth
//go:noescape
func ClientAuth(c *smtp.Client, a smtp.Auth) error

//go:linkname clienthello smtp.sub_clienthello
func clienthello(c *smtp.Client, localName string) error {
	return c.Hello(localName)
}

//go:linkname ClientHello smtp.sub_clienthello
//go:noescape
func ClientHello(c *smtp.Client, localName string) error

//go:linkname clientnoop smtp.sub_clientnoop
func clientnoop(c *smtp.Client) error {
	return c.Noop()
}

//go:linkname ClientNoop smtp.sub_clientnoop
//go:noescape
func ClientNoop(c *smtp.Client) error

//go:linkname clientquit smtp.sub_clientquit
func clientquit(c *smtp.Client) error {
	return c.Quit()
}

//go:linkname ClientQuit smtp.sub_clientquit
//go:noescape
func ClientQuit(c *smtp.Client) error
