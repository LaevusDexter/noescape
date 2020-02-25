// This file has automatically been generated on Wed Feb 26 02:10:05 +05 2020.
// DO NOT EDIT.
package net

import (
	"context"
	"io"
	"net"
	"os"
	"syscall"
	"time"
	_ "unsafe"
)

//go:linkname ParseCIDR net.ParseCIDR
//go:noescape
func ParseCIDR(s string) (net.IP, *net.IPNet, error)

//go:linkname ipconnsetreaddeadline net.sub_ipconnsetreaddeadline
func ipconnsetreaddeadline(c *net.IPConn, t time.Time) error {
	return c.SetReadDeadline(t)
}

//go:linkname IPConnSetReadDeadline net.sub_ipconnsetreaddeadline
//go:noescape
func IPConnSetReadDeadline(c *net.IPConn, t time.Time) error

//go:linkname ipmasksize net.sub_ipmasksize
func ipmasksize(m net.IPMask) int {
	return m.Size()
}

//go:linkname IPMaskSize net.sub_ipmasksize
//go:noescape
func IPMaskSize(m net.IPMask) int

//go:linkname ListenMulticastUDP net.ListenMulticastUDP
//go:noescape
func ListenMulticastUDP(network string, ifi *net.Interface, gaddr *net.UDPAddr) (*net.UDPConn, error)

//go:linkname unixconnread net.sub_unixconnread
func unixconnread(c *net.UnixConn, b []byte) (int, error) {
	return c.Read(b)
}

//go:linkname UnixConnRead net.sub_unixconnread
//go:noescape
func UnixConnRead(c *net.UnixConn, b []byte) (int, error)

//go:linkname unixconnsetreaddeadline net.sub_unixconnsetreaddeadline
func unixconnsetreaddeadline(c *net.UnixConn, t time.Time) error {
	return c.SetReadDeadline(t)
}

//go:linkname UnixConnSetReadDeadline net.sub_unixconnsetreaddeadline
//go:noescape
func UnixConnSetReadDeadline(c *net.UnixConn, t time.Time) error

//go:linkname unixconnwritemsgunix net.sub_unixconnwritemsgunix
func unixconnwritemsgunix(c *net.UnixConn, b, oob []byte, addr *net.UnixAddr) (int, error) {
	return c.WriteMsgUnix(b, oob, addr)
}

//go:linkname UnixConnWriteMsgUnix net.sub_unixconnwritemsgunix
//go:noescape
func UnixConnWriteMsgUnix(c *net.UnixConn, b, oob []byte, addr *net.UnixAddr) (int, error)

//go:linkname ParseIP net.ParseIP
//go:noescape
func ParseIP(s string) net.IP

//go:linkname ipislinklocalunicast net.sub_ipislinklocalunicast
func ipislinklocalunicast(ip net.IP) bool {
	return ip.IsLinkLocalUnicast()
}

//go:linkname IPIsLinkLocalUnicast net.sub_ipislinklocalunicast
//go:noescape
func IPIsLinkLocalUnicast(ip net.IP) bool

//go:linkname tcpconnremoteaddr net.sub_tcpconnremoteaddr
func tcpconnremoteaddr(c *net.TCPConn) net.Addr {
	return c.RemoteAddr()
}

//go:linkname TCPConnRemoteAddr net.sub_tcpconnremoteaddr
//go:noescape
func TCPConnRemoteAddr(c *net.TCPConn) net.Addr

//go:linkname tcpconnsetnodelay net.sub_tcpconnsetnodelay
func tcpconnsetnodelay(c *net.TCPConn, noDelay bool) error {
	return c.SetNoDelay(noDelay)
}

//go:linkname TCPConnSetNoDelay net.sub_tcpconnsetnodelay
//go:noescape
func TCPConnSetNoDelay(c *net.TCPConn, noDelay bool) error

//go:linkname udpconnsetreaddeadline net.sub_udpconnsetreaddeadline
func udpconnsetreaddeadline(c *net.UDPConn, t time.Time) error {
	return c.SetReadDeadline(t)
}

//go:linkname UDPConnSetReadDeadline net.sub_udpconnsetreaddeadline
//go:noescape
func UDPConnSetReadDeadline(c *net.UDPConn, t time.Time) error

//go:linkname JoinHostPort net.JoinHostPort
//go:noescape
func JoinHostPort(host, port string) string

//go:linkname LookupHost net.LookupHost
//go:noescape
func LookupHost(host string) ([]string, error)

//go:linkname interfaceaddrs net.sub_interfaceaddrs
func interfaceaddrs(ifi *net.Interface) ([]net.Addr, error) {
	return ifi.Addrs()
}

//go:linkname InterfaceAddrs net.sub_interfaceaddrs
//go:noescape
func InterfaceAddrs(ifi *net.Interface) ([]net.Addr, error)

//go:linkname dnsconfigerrortimeout net.sub_dnsconfigerrortimeout
func dnsconfigerrortimeout(e *net.DNSConfigError) bool {
	return e.Timeout()
}

//go:linkname DNSConfigErrorTimeout net.sub_dnsconfigerrortimeout
//go:noescape
func DNSConfigErrorTimeout(e *net.DNSConfigError) bool

//go:linkname ipismulticast net.sub_ipismulticast
func ipismulticast(ip net.IP) bool {
	return ip.IsMulticast()
}

//go:linkname IPIsMulticast net.sub_ipismulticast
//go:noescape
func IPIsMulticast(ip net.IP) bool

//go:linkname tcplisteneraccept net.sub_tcplisteneraccept
func tcplisteneraccept(l *net.TCPListener) (net.Conn, error) {
	return l.Accept()
}

//go:linkname TCPListenerAccept net.sub_tcplisteneraccept
//go:noescape
func TCPListenerAccept(l *net.TCPListener) (net.Conn, error)

//go:linkname ResolveUDPAddr net.ResolveUDPAddr
//go:noescape
func ResolveUDPAddr(network, address string) (*net.UDPAddr, error)

//go:linkname unixconnreadmsgunix net.sub_unixconnreadmsgunix
func unixconnreadmsgunix(c *net.UnixConn, b, oob []byte) (int, *net.UnixAddr, error) {
	return c.ReadMsgUnix(b, oob)
}

//go:linkname UnixConnReadMsgUnix net.sub_unixconnreadmsgunix
//go:noescape
func UnixConnReadMsgUnix(c *net.UnixConn, b, oob []byte) (int, *net.UnixAddr, error)

//go:linkname buffersread net.sub_buffersread
func buffersread(v *net.Buffers, p []byte) (int, error) {
	return v.Read(p)
}

//go:linkname BuffersRead net.sub_buffersread
//go:noescape
func BuffersRead(v *net.Buffers, p []byte) (int, error)

//go:linkname hardwareaddrstring net.sub_hardwareaddrstring
func hardwareaddrstring(a net.HardwareAddr) string {
	return a.String()
}

//go:linkname HardwareAddrString net.sub_hardwareaddrstring
//go:noescape
func HardwareAddrString(a net.HardwareAddr) string

//go:linkname ipmarshaltext net.sub_ipmarshaltext
func ipmarshaltext(ip net.IP) ([]byte, error) {
	return ip.MarshalText()
}

//go:linkname IPMarshalText net.sub_ipmarshaltext
//go:noescape
func IPMarshalText(ip net.IP) ([]byte, error)

//go:linkname ListenIP net.ListenIP
//go:noescape
func ListenIP(network string, laddr *net.IPAddr) (*net.IPConn, error)

//go:linkname unixconnsetdeadline net.sub_unixconnsetdeadline
func unixconnsetdeadline(c *net.UnixConn, t time.Time) error {
	return c.SetDeadline(t)
}

//go:linkname UnixConnSetDeadline net.sub_unixconnsetdeadline
//go:noescape
func UnixConnSetDeadline(c *net.UnixConn, t time.Time) error

//go:linkname SplitHostPort net.SplitHostPort
//go:noescape
func SplitHostPort(hostport string) (string, error)

//go:linkname LookupIP net.LookupIP
//go:noescape
func LookupIP(host string) ([]net.IP, error)

//go:linkname ipnetcontains net.sub_ipnetcontains
func ipnetcontains(n *net.IPNet, ip net.IP) bool {
	return n.Contains(ip)
}

//go:linkname IPNetContains net.sub_ipnetcontains
//go:noescape
func IPNetContains(n *net.IPNet, ip net.IP) bool

//go:linkname ipnetstring net.sub_ipnetstring
func ipnetstring(n *net.IPNet) string {
	return n.String()
}

//go:linkname IPNetString net.sub_ipnetstring
//go:noescape
func IPNetString(n *net.IPNet) string

//go:linkname dnserrortimeout net.sub_dnserrortimeout
func dnserrortimeout(e *net.DNSError) bool {
	return e.Timeout()
}

//go:linkname DNSErrorTimeout net.sub_dnserrortimeout
//go:noescape
func DNSErrorTimeout(e *net.DNSError) bool

//go:linkname ipisinterfacelocalmulticast net.sub_ipisinterfacelocalmulticast
func ipisinterfacelocalmulticast(ip net.IP) bool {
	return ip.IsInterfaceLocalMulticast()
}

//go:linkname IPIsInterfaceLocalMulticast net.sub_ipisinterfacelocalmulticast
//go:noescape
func IPIsInterfaceLocalMulticast(ip net.IP) bool

//go:linkname ipconnclose net.sub_ipconnclose
func ipconnclose(c *net.IPConn) error {
	return c.Close()
}

//go:linkname IPConnClose net.sub_ipconnclose
//go:noescape
func IPConnClose(c *net.IPConn) error

//go:linkname ipconnwriteto net.sub_ipconnwriteto
func ipconnwriteto(c *net.IPConn, b []byte, addr net.Addr) (int, error) {
	return c.WriteTo(b, addr)
}

//go:linkname IPConnWriteTo net.sub_ipconnwriteto
//go:noescape
func IPConnWriteTo(c *net.IPConn, b []byte, addr net.Addr) (int, error)

//go:linkname resolverlookupmx net.sub_resolverlookupmx
func resolverlookupmx(r *net.Resolver, ctx context.Context, name string) ([]*net.MX, error) {
	return r.LookupMX(ctx, name)
}

//go:linkname ResolverLookupMX net.sub_resolverlookupmx
//go:noescape
func ResolverLookupMX(r *net.Resolver, ctx context.Context, name string) ([]*net.MX, error)

//go:linkname unixconncloseread net.sub_unixconncloseread
func unixconncloseread(c *net.UnixConn) error {
	return c.CloseRead()
}

//go:linkname UnixConnCloseRead net.sub_unixconncloseread
//go:noescape
func UnixConnCloseRead(c *net.UnixConn) error

//go:linkname unixconnwritetounix net.sub_unixconnwritetounix
func unixconnwritetounix(c *net.UnixConn, b []byte, addr *net.UnixAddr) (int, error) {
	return c.WriteToUnix(b, addr)
}

//go:linkname UnixConnWriteToUnix net.sub_unixconnwritetounix
//go:noescape
func UnixConnWriteToUnix(c *net.UnixConn, b []byte, addr *net.UnixAddr) (int, error)

//go:linkname unixlisteneracceptunix net.sub_unixlisteneracceptunix
func unixlisteneracceptunix(l *net.UnixListener) (*net.UnixConn, error) {
	return l.AcceptUnix()
}

//go:linkname UnixListenerAcceptUnix net.sub_unixlisteneracceptunix
//go:noescape
func UnixListenerAcceptUnix(l *net.UnixListener) (*net.UnixConn, error)

//go:linkname bufferswriteto net.sub_bufferswriteto
func bufferswriteto(v *net.Buffers, w io.Writer) (int64, error) {
	return v.WriteTo(w)
}

//go:linkname BuffersWriteTo net.sub_bufferswriteto
//go:noescape
func BuffersWriteTo(v *net.Buffers, w io.Writer) (int64, error)

//go:linkname Listen net.Listen
//go:noescape
func Listen(network, address string) (net.Listener, error)

//go:linkname tcpconnsetreaddeadline net.sub_tcpconnsetreaddeadline
func tcpconnsetreaddeadline(c *net.TCPConn, t time.Time) error {
	return c.SetReadDeadline(t)
}

//go:linkname TCPConnSetReadDeadline net.sub_tcpconnsetreaddeadline
//go:noescape
func TCPConnSetReadDeadline(c *net.TCPConn, t time.Time) error

//go:linkname ipequal net.sub_ipequal
func ipequal(ip net.IP, x net.IP) bool {
	return ip.Equal(x)
}

//go:linkname IPEqual net.sub_ipequal
//go:noescape
func IPEqual(ip net.IP, x net.IP) bool

//go:linkname ListenPacket net.ListenPacket
//go:noescape
func ListenPacket(network, address string) (net.PacketConn, error)

//go:linkname ListenTCP net.ListenTCP
//go:noescape
func ListenTCP(network string, laddr *net.TCPAddr) (*net.TCPListener, error)

//go:linkname ListenUDP net.ListenUDP
//go:noescape
func ListenUDP(network string, laddr *net.UDPAddr) (*net.UDPConn, error)

//go:linkname udpconnreadmsgudp net.sub_udpconnreadmsgudp
func udpconnreadmsgudp(c *net.UDPConn, b, oob []byte) (int, *net.UDPAddr, error) {
	return c.ReadMsgUDP(b, oob)
}

//go:linkname UDPConnReadMsgUDP net.sub_udpconnreadmsgudp
//go:noescape
func UDPConnReadMsgUDP(c *net.UDPConn, b, oob []byte) (int, *net.UDPAddr, error)

//go:linkname dialerdial net.sub_dialerdial
func dialerdial(d *net.Dialer, network, address string) (net.Conn, error) {
	return d.Dial(network, address)
}

//go:linkname DialerDial net.sub_dialerdial
//go:noescape
func DialerDial(d *net.Dialer, network, address string) (net.Conn, error)

//go:linkname ipisglobalunicast net.sub_ipisglobalunicast
func ipisglobalunicast(ip net.IP) bool {
	return ip.IsGlobalUnicast()
}

//go:linkname IPIsGlobalUnicast net.sub_ipisglobalunicast
//go:noescape
func IPIsGlobalUnicast(ip net.IP) bool

//go:linkname ipconnsetwritebuffer net.sub_ipconnsetwritebuffer
func ipconnsetwritebuffer(c *net.IPConn, bytes int) error {
	return c.SetWriteBuffer(bytes)
}

//go:linkname IPConnSetWriteBuffer net.sub_ipconnsetwritebuffer
//go:noescape
func IPConnSetWriteBuffer(c *net.IPConn, bytes int) error

//go:linkname ipnetnetwork net.sub_ipnetnetwork
func ipnetnetwork(n *net.IPNet) string {
	return n.Network()
}

//go:linkname IPNetNetwork net.sub_ipnetnetwork
//go:noescape
func IPNetNetwork(n *net.IPNet) string

//go:linkname operrorunwrap net.sub_operrorunwrap
func operrorunwrap(e *net.OpError) error {
	return e.Unwrap()
}

//go:linkname OpErrorUnwrap net.sub_operrorunwrap
//go:noescape
func OpErrorUnwrap(e *net.OpError) error

//go:linkname udpconnfile net.sub_udpconnfile
func udpconnfile(c *net.UDPConn) (*os.File, error) {
	return c.File()
}

//go:linkname UDPConnFile net.sub_udpconnfile
//go:noescape
func UDPConnFile(c *net.UDPConn) (*os.File, error)

//go:linkname ipconnsetwritedeadline net.sub_ipconnsetwritedeadline
func ipconnsetwritedeadline(c *net.IPConn, t time.Time) error {
	return c.SetWriteDeadline(t)
}

//go:linkname IPConnSetWriteDeadline net.sub_ipconnsetwritedeadline
//go:noescape
func IPConnSetWriteDeadline(c *net.IPConn, t time.Time) error

//go:linkname operrortemporary net.sub_operrortemporary
func operrortemporary(e *net.OpError) bool {
	return e.Temporary()
}

//go:linkname OpErrorTemporary net.sub_operrortemporary
//go:noescape
func OpErrorTemporary(e *net.OpError) bool

//go:linkname ListenUnixgram net.ListenUnixgram
//go:noescape
func ListenUnixgram(network string, laddr *net.UnixAddr) (*net.UnixConn, error)

//go:linkname unixconnsetreadbuffer net.sub_unixconnsetreadbuffer
func unixconnsetreadbuffer(c *net.UnixConn, bytes int) error {
	return c.SetReadBuffer(bytes)
}

//go:linkname UnixConnSetReadBuffer net.sub_unixconnsetreadbuffer
//go:noescape
func UnixConnSetReadBuffer(c *net.UnixConn, bytes int) error

//go:linkname unixconnsetwritedeadline net.sub_unixconnsetwritedeadline
func unixconnsetwritedeadline(c *net.UnixConn, t time.Time) error {
	return c.SetWriteDeadline(t)
}

//go:linkname UnixConnSetWriteDeadline net.sub_unixconnsetwritedeadline
//go:noescape
func UnixConnSetWriteDeadline(c *net.UnixConn, t time.Time) error

//go:linkname unixconnwriteto net.sub_unixconnwriteto
func unixconnwriteto(c *net.UnixConn, b []byte, addr net.Addr) (int, error) {
	return c.WriteTo(b, addr)
}

//go:linkname UnixConnWriteTo net.sub_unixconnwriteto
//go:noescape
func UnixConnWriteTo(c *net.UnixConn, b []byte, addr net.Addr) (int, error)

//go:linkname dnserrortemporary net.sub_dnserrortemporary
func dnserrortemporary(e *net.DNSError) bool {
	return e.Temporary()
}

//go:linkname DNSErrorTemporary net.sub_dnserrortemporary
//go:noescape
func DNSErrorTemporary(e *net.DNSError) bool

//go:linkname ipisunspecified net.sub_ipisunspecified
func ipisunspecified(ip net.IP) bool {
	return ip.IsUnspecified()
}

//go:linkname IPIsUnspecified net.sub_ipisunspecified
//go:noescape
func IPIsUnspecified(ip net.IP) bool

//go:linkname ipconnread net.sub_ipconnread
func ipconnread(c *net.IPConn, b []byte) (int, error) {
	return c.Read(b)
}

//go:linkname IPConnRead net.sub_ipconnread
//go:noescape
func IPConnRead(c *net.IPConn, b []byte) (int, error)

//go:linkname ipconnremoteaddr net.sub_ipconnremoteaddr
func ipconnremoteaddr(c *net.IPConn) net.Addr {
	return c.RemoteAddr()
}

//go:linkname IPConnRemoteAddr net.sub_ipconnremoteaddr
//go:noescape
func IPConnRemoteAddr(c *net.IPConn) net.Addr

//go:linkname InterfaceByName net.InterfaceByName
//go:noescape
func InterfaceByName(name string) (*net.Interface, error)

//go:linkname resolverlookupaddr net.sub_resolverlookupaddr
func resolverlookupaddr(r *net.Resolver, ctx context.Context, addr string) ([]string, error) {
	return r.LookupAddr(ctx, addr)
}

//go:linkname ResolverLookupAddr net.sub_resolverlookupaddr
//go:noescape
func ResolverLookupAddr(r *net.Resolver, ctx context.Context, addr string) ([]string, error)

//go:linkname DialUnix net.DialUnix
//go:noescape
func DialUnix(network string, laddr, raddr *net.UnixAddr) (*net.UnixConn, error)

//go:linkname ipconnlocaladdr net.sub_ipconnlocaladdr
func ipconnlocaladdr(c *net.IPConn) net.Addr {
	return c.LocalAddr()
}

//go:linkname IPConnLocalAddr net.sub_ipconnlocaladdr
//go:noescape
func IPConnLocalAddr(c *net.IPConn) net.Addr

//go:linkname LookupNS net.LookupNS
//go:noescape
func LookupNS(name string) ([]*net.NS, error)

//go:linkname tcplisteneraddr net.sub_tcplisteneraddr
func tcplisteneraddr(l *net.TCPListener) net.Addr {
	return l.Addr()
}

//go:linkname TCPListenerAddr net.sub_tcplisteneraddr
//go:noescape
func TCPListenerAddr(l *net.TCPListener) net.Addr

//go:linkname udpconnremoteaddr net.sub_udpconnremoteaddr
func udpconnremoteaddr(c *net.UDPConn) net.Addr {
	return c.RemoteAddr()
}

//go:linkname UDPConnRemoteAddr net.sub_udpconnremoteaddr
//go:noescape
func UDPConnRemoteAddr(c *net.UDPConn) net.Addr

//go:linkname unixconnreadfromunix net.sub_unixconnreadfromunix
func unixconnreadfromunix(c *net.UnixConn, b []byte) (int, *net.UnixAddr, error) {
	return c.ReadFromUnix(b)
}

//go:linkname UnixConnReadFromUnix net.sub_unixconnreadfromunix
//go:noescape
func UnixConnReadFromUnix(c *net.UnixConn, b []byte) (int, *net.UnixAddr, error)

//go:linkname ipconnfile net.sub_ipconnfile
func ipconnfile(c *net.IPConn) (*os.File, error) {
	return c.File()
}

//go:linkname IPConnFile net.sub_ipconnfile
//go:noescape
func IPConnFile(c *net.IPConn) (*os.File, error)

//go:linkname ipconnsyscallconn net.sub_ipconnsyscallconn
func ipconnsyscallconn(c *net.IPConn) (syscall.RawConn, error) {
	return c.SyscallConn()
}

//go:linkname IPConnSyscallConn net.sub_ipconnsyscallconn
//go:noescape
func IPConnSyscallConn(c *net.IPConn) (syscall.RawConn, error)

//go:linkname tcplistenerclose net.sub_tcplistenerclose
func tcplistenerclose(l *net.TCPListener) error {
	return l.Close()
}

//go:linkname TCPListenerClose net.sub_tcplistenerclose
//go:noescape
func TCPListenerClose(l *net.TCPListener) error

//go:linkname unixlistenersyscallconn net.sub_unixlistenersyscallconn
func unixlistenersyscallconn(l *net.UnixListener) (syscall.RawConn, error) {
	return l.SyscallConn()
}

//go:linkname UnixListenerSyscallConn net.sub_unixlistenersyscallconn
//go:noescape
func UnixListenerSyscallConn(l *net.UnixListener) (syscall.RawConn, error)

//go:linkname unknownnetworkerrortemporary net.sub_unknownnetworkerrortemporary
func unknownnetworkerrortemporary(e net.UnknownNetworkError) bool {
	return e.Temporary()
}

//go:linkname UnknownNetworkErrorTemporary net.sub_unknownnetworkerrortemporary
//go:noescape
func UnknownNetworkErrorTemporary(e net.UnknownNetworkError) bool

//go:linkname DialTimeout net.DialTimeout
//go:noescape
func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)

//go:linkname ipislinklocalmulticast net.sub_ipislinklocalmulticast
func ipislinklocalmulticast(ip net.IP) bool {
	return ip.IsLinkLocalMulticast()
}

//go:linkname IPIsLinkLocalMulticast net.sub_ipislinklocalmulticast
//go:noescape
func IPIsLinkLocalMulticast(ip net.IP) bool

//go:linkname ipstring net.sub_ipstring
func ipstring(ip net.IP) string {
	return ip.String()
}

//go:linkname IPString net.sub_ipstring
//go:noescape
func IPString(ip net.IP) string

//go:linkname ipaddrstring net.sub_ipaddrstring
func ipaddrstring(a *net.IPAddr) string {
	return a.String()
}

//go:linkname IPAddrString net.sub_ipaddrstring
//go:noescape
func IPAddrString(a *net.IPAddr) string

//go:linkname udpconnsetdeadline net.sub_udpconnsetdeadline
func udpconnsetdeadline(c *net.UDPConn, t time.Time) error {
	return c.SetDeadline(t)
}

//go:linkname UDPConnSetDeadline net.sub_udpconnsetdeadline
//go:noescape
func UDPConnSetDeadline(c *net.UDPConn, t time.Time) error

//go:linkname unixlisteneraddr net.sub_unixlisteneraddr
func unixlisteneraddr(l *net.UnixListener) net.Addr {
	return l.Addr()
}

//go:linkname UnixListenerAddr net.sub_unixlisteneraddr
//go:noescape
func UnixListenerAddr(l *net.UnixListener) net.Addr

//go:linkname ipconnwritemsgip net.sub_ipconnwritemsgip
func ipconnwritemsgip(c *net.IPConn, b, oob []byte, addr *net.IPAddr) (int, error) {
	return c.WriteMsgIP(b, oob, addr)
}

//go:linkname IPConnWriteMsgIP net.sub_ipconnwritemsgip
//go:noescape
func IPConnWriteMsgIP(c *net.IPConn, b, oob []byte, addr *net.IPAddr) (int, error)

//go:linkname InterfaceByIndex net.InterfaceByIndex
//go:noescape
func InterfaceByIndex(index int) (*net.Interface, error)

//go:linkname tcpconnsetkeepalive net.sub_tcpconnsetkeepalive
func tcpconnsetkeepalive(c *net.TCPConn, keepalive bool) error {
	return c.SetKeepAlive(keepalive)
}

//go:linkname TCPConnSetKeepAlive net.sub_tcpconnsetkeepalive
//go:noescape
func TCPConnSetKeepAlive(c *net.TCPConn, keepalive bool) error

//go:linkname tcpconnsetwritedeadline net.sub_tcpconnsetwritedeadline
func tcpconnsetwritedeadline(c *net.TCPConn, t time.Time) error {
	return c.SetWriteDeadline(t)
}

//go:linkname TCPConnSetWriteDeadline net.sub_tcpconnsetwritedeadline
//go:noescape
func TCPConnSetWriteDeadline(c *net.TCPConn, t time.Time) error

//go:linkname ResolveUnixAddr net.ResolveUnixAddr
//go:noescape
func ResolveUnixAddr(network, address string) (*net.UnixAddr, error)

//go:linkname unixaddrstring net.sub_unixaddrstring
func unixaddrstring(a *net.UnixAddr) string {
	return a.String()
}

//go:linkname UnixAddrString net.sub_unixaddrstring
//go:noescape
func UnixAddrString(a *net.UnixAddr) string

//go:linkname ipunmarshaltext net.sub_ipunmarshaltext
func ipunmarshaltext(ip *net.IP, text []byte) error {
	return ip.UnmarshalText(text)
}

//go:linkname IPUnmarshalText net.sub_ipunmarshaltext
//go:noescape
func IPUnmarshalText(ip *net.IP, text []byte) error

//go:linkname DialIP net.DialIP
//go:noescape
func DialIP(network string, laddr, raddr *net.IPAddr) (*net.IPConn, error)

//go:linkname ipconnwritetoip net.sub_ipconnwritetoip
func ipconnwritetoip(c *net.IPConn, b []byte, addr *net.IPAddr) (int, error) {
	return c.WriteToIP(b, addr)
}

//go:linkname IPConnWriteToIP net.sub_ipconnwritetoip
//go:noescape
func IPConnWriteToIP(c *net.IPConn, b []byte, addr *net.IPAddr) (int, error)

//go:linkname tcpconnwrite net.sub_tcpconnwrite
func tcpconnwrite(c *net.TCPConn, b []byte) (int, error) {
	return c.Write(b)
}

//go:linkname TCPConnWrite net.sub_tcpconnwrite
//go:noescape
func TCPConnWrite(c *net.TCPConn, b []byte) (int, error)

//go:linkname Pipe net.Pipe
//go:noescape
func Pipe() (net.Conn, net.Conn)

//go:linkname ParseMAC net.ParseMAC
//go:noescape
func ParseMAC(s string) (net.HardwareAddr, error)

//go:linkname Interfaces net.Interfaces
//go:noescape
func Interfaces() ([]net.Interface, error)

//go:linkname udpconnread net.sub_udpconnread
func udpconnread(c *net.UDPConn, b []byte) (int, error) {
	return c.Read(b)
}

//go:linkname UDPConnRead net.sub_udpconnread
//go:noescape
func UDPConnRead(c *net.UDPConn, b []byte) (int, error)

//go:linkname udpconnsetreadbuffer net.sub_udpconnsetreadbuffer
func udpconnsetreadbuffer(c *net.UDPConn, bytes int) error {
	return c.SetReadBuffer(bytes)
}

//go:linkname UDPConnSetReadBuffer net.sub_udpconnsetreadbuffer
//go:noescape
func UDPConnSetReadBuffer(c *net.UDPConn, bytes int) error

//go:linkname parseerrorerror net.sub_parseerrorerror
func parseerrorerror(e *net.ParseError) string {
	return e.Error()
}

//go:linkname ParseErrorError net.sub_parseerrorerror
//go:noescape
func ParseErrorError(e *net.ParseError) string

//go:linkname tcplistenerfile net.sub_tcplistenerfile
func tcplistenerfile(l *net.TCPListener) (*os.File, error) {
	return l.File()
}

//go:linkname TCPListenerFile net.sub_tcplistenerfile
//go:noescape
func TCPListenerFile(l *net.TCPListener) (*os.File, error)

//go:linkname udpconnwrite net.sub_udpconnwrite
func udpconnwrite(c *net.UDPConn, b []byte) (int, error) {
	return c.Write(b)
}

//go:linkname UDPConnWrite net.sub_udpconnwrite
//go:noescape
func UDPConnWrite(c *net.UDPConn, b []byte) (int, error)

//go:linkname unixconnfile net.sub_unixconnfile
func unixconnfile(c *net.UnixConn) (*os.File, error) {
	return c.File()
}

//go:linkname UnixConnFile net.sub_unixconnfile
//go:noescape
func UnixConnFile(c *net.UnixConn) (*os.File, error)

//go:linkname udpconnclose net.sub_udpconnclose
func udpconnclose(c *net.UDPConn) error {
	return c.Close()
}

//go:linkname UDPConnClose net.sub_udpconnclose
//go:noescape
func UDPConnClose(c *net.UDPConn) error

//go:linkname LookupAddr net.LookupAddr
//go:noescape
func LookupAddr(addr string) ([]string, error)

//go:linkname invalidaddrerrortimeout net.sub_invalidaddrerrortimeout
func invalidaddrerrortimeout(e net.InvalidAddrError) bool {
	return e.Timeout()
}

//go:linkname InvalidAddrErrorTimeout net.sub_invalidaddrerrortimeout
//go:noescape
func InvalidAddrErrorTimeout(e net.InvalidAddrError) bool

//go:linkname FilePacketConn net.FilePacketConn
//go:noescape
func FilePacketConn(f *os.File) (net.PacketConn, error)

//go:linkname resolverlookuphost net.sub_resolverlookuphost
func resolverlookuphost(r *net.Resolver, ctx context.Context, host string) ([]string, error) {
	return r.LookupHost(ctx, host)
}

//go:linkname ResolverLookupHost net.sub_resolverlookuphost
//go:noescape
func ResolverLookupHost(r *net.Resolver, ctx context.Context, host string) ([]string, error)

//go:linkname resolverlookupport net.sub_resolverlookupport
func resolverlookupport(r *net.Resolver, ctx context.Context, network, service string) (int, error) {
	return r.LookupPort(ctx, network, service)
}

//go:linkname ResolverLookupPort net.sub_resolverlookupport
//go:noescape
func ResolverLookupPort(r *net.Resolver, ctx context.Context, network, service string) (int, error)

//go:linkname udpaddrnetwork net.sub_udpaddrnetwork
func udpaddrnetwork(a *net.UDPAddr) string {
	return a.Network()
}

//go:linkname UDPAddrNetwork net.sub_udpaddrnetwork
//go:noescape
func UDPAddrNetwork(a *net.UDPAddr) string

//go:linkname udpaddrstring net.sub_udpaddrstring
func udpaddrstring(a *net.UDPAddr) string {
	return a.String()
}

//go:linkname UDPAddrString net.sub_udpaddrstring
//go:noescape
func UDPAddrString(a *net.UDPAddr) string

//go:linkname udpconnreadfrom net.sub_udpconnreadfrom
func udpconnreadfrom(c *net.UDPConn, b []byte) (int, net.Addr, error) {
	return c.ReadFrom(b)
}

//go:linkname UDPConnReadFrom net.sub_udpconnreadfrom
//go:noescape
func UDPConnReadFrom(c *net.UDPConn, b []byte) (int, net.Addr, error)

//go:linkname unixlistenerfile net.sub_unixlistenerfile
func unixlistenerfile(l *net.UnixListener) (*os.File, error) {
	return l.File()
}

//go:linkname UnixListenerFile net.sub_unixlistenerfile
//go:noescape
func UnixListenerFile(l *net.UnixListener) (*os.File, error)

//go:linkname tcpconnsetlinger net.sub_tcpconnsetlinger
func tcpconnsetlinger(c *net.TCPConn, sec int) error {
	return c.SetLinger(sec)
}

//go:linkname TCPConnSetLinger net.sub_tcpconnsetlinger
//go:noescape
func TCPConnSetLinger(c *net.TCPConn, sec int) error

//go:linkname LookupPort net.LookupPort
//go:noescape
func LookupPort(network, service string) (int, error)

//go:linkname addrerrortimeout net.sub_addrerrortimeout
func addrerrortimeout(e *net.AddrError) bool {
	return e.Timeout()
}

//go:linkname AddrErrorTimeout net.sub_addrerrortimeout
//go:noescape
func AddrErrorTimeout(e *net.AddrError) bool

//go:linkname FileConn net.FileConn
//go:noescape
func FileConn(f *os.File) (net.Conn, error)

//go:linkname dnserrorerror net.sub_dnserrorerror
func dnserrorerror(e *net.DNSError) string {
	return e.Error()
}

//go:linkname DNSErrorError net.sub_dnserrorerror
//go:noescape
func DNSErrorError(e *net.DNSError) string

//go:linkname interfacemulticastaddrs net.sub_interfacemulticastaddrs
func interfacemulticastaddrs(ifi *net.Interface) ([]net.Addr, error) {
	return ifi.MulticastAddrs()
}

//go:linkname InterfaceMulticastAddrs net.sub_interfacemulticastaddrs
//go:noescape
func InterfaceMulticastAddrs(ifi *net.Interface) ([]net.Addr, error)

//go:linkname invalidaddrerrorerror net.sub_invalidaddrerrorerror
func invalidaddrerrorerror(e net.InvalidAddrError) string {
	return e.Error()
}

//go:linkname InvalidAddrErrorError net.sub_invalidaddrerrorerror
//go:noescape
func InvalidAddrErrorError(e net.InvalidAddrError) string

//go:linkname listenconfiglisten net.sub_listenconfiglisten
func listenconfiglisten(lc *net.ListenConfig, ctx context.Context, network, address string) (net.Listener, error) {
	return lc.Listen(ctx, network, address)
}

//go:linkname ListenConfigListen net.sub_listenconfiglisten
//go:noescape
func ListenConfigListen(lc *net.ListenConfig, ctx context.Context, network, address string) (net.Listener, error)

//go:linkname unixlisteneraccept net.sub_unixlisteneraccept
func unixlisteneraccept(l *net.UnixListener) (net.Conn, error) {
	return l.Accept()
}

//go:linkname UnixListenerAccept net.sub_unixlisteneraccept
//go:noescape
func UnixListenerAccept(l *net.UnixListener) (net.Conn, error)

//go:linkname unixlistenersetdeadline net.sub_unixlistenersetdeadline
func unixlistenersetdeadline(l *net.UnixListener, t time.Time) error {
	return l.SetDeadline(t)
}

//go:linkname UnixListenerSetDeadline net.sub_unixlistenersetdeadline
//go:noescape
func UnixListenerSetDeadline(l *net.UnixListener, t time.Time) error

//go:linkname LookupMX net.LookupMX
//go:noescape
func LookupMX(name string) ([]*net.MX, error)

//go:linkname ListenUnix net.ListenUnix
//go:noescape
func ListenUnix(network string, laddr *net.UnixAddr) (*net.UnixListener, error)

//go:linkname tcpconnsetkeepaliveperiod net.sub_tcpconnsetkeepaliveperiod
func tcpconnsetkeepaliveperiod(c *net.TCPConn, d time.Duration) error {
	return c.SetKeepAlivePeriod(d)
}

//go:linkname TCPConnSetKeepAlivePeriod net.sub_tcpconnsetkeepaliveperiod
//go:noescape
func TCPConnSetKeepAlivePeriod(c *net.TCPConn, d time.Duration) error

//go:linkname ipconnsetdeadline net.sub_ipconnsetdeadline
func ipconnsetdeadline(c *net.IPConn, t time.Time) error {
	return c.SetDeadline(t)
}

//go:linkname IPConnSetDeadline net.sub_ipconnsetdeadline
//go:noescape
func IPConnSetDeadline(c *net.IPConn, t time.Time) error

//go:linkname ipconnsetreadbuffer net.sub_ipconnsetreadbuffer
func ipconnsetreadbuffer(c *net.IPConn, bytes int) error {
	return c.SetReadBuffer(bytes)
}

//go:linkname IPConnSetReadBuffer net.sub_ipconnsetreadbuffer
//go:noescape
func IPConnSetReadBuffer(c *net.IPConn, bytes int) error

//go:linkname CIDRMask net.CIDRMask
//go:noescape
func CIDRMask(ones, bits int) net.IPMask

//go:linkname FileListener net.FileListener
//go:noescape
func FileListener(f *os.File) (net.Listener, error)

//go:linkname LookupSRV net.LookupSRV
//go:noescape
func LookupSRV(service, proto, name string) (string, []*net.SRV, error)

//go:linkname tcpaddrstring net.sub_tcpaddrstring
func tcpaddrstring(a *net.TCPAddr) string {
	return a.String()
}

//go:linkname TCPAddrString net.sub_tcpaddrstring
//go:noescape
func TCPAddrString(a *net.TCPAddr) string

//go:linkname tcpconnsetdeadline net.sub_tcpconnsetdeadline
func tcpconnsetdeadline(c *net.TCPConn, t time.Time) error {
	return c.SetDeadline(t)
}

//go:linkname TCPConnSetDeadline net.sub_tcpconnsetdeadline
//go:noescape
func TCPConnSetDeadline(c *net.TCPConn, t time.Time) error

//go:linkname udpconnsetwritedeadline net.sub_udpconnsetwritedeadline
func udpconnsetwritedeadline(c *net.UDPConn, t time.Time) error {
	return c.SetWriteDeadline(t)
}

//go:linkname UDPConnSetWriteDeadline net.sub_udpconnsetwritedeadline
//go:noescape
func UDPConnSetWriteDeadline(c *net.UDPConn, t time.Time) error

//go:linkname udpconnwritemsgudp net.sub_udpconnwritemsgudp
func udpconnwritemsgudp(c *net.UDPConn, b, oob []byte, addr *net.UDPAddr) (int, error) {
	return c.WriteMsgUDP(b, oob, addr)
}

//go:linkname UDPConnWriteMsgUDP net.sub_udpconnwritemsgudp
//go:noescape
func UDPConnWriteMsgUDP(c *net.UDPConn, b, oob []byte, addr *net.UDPAddr) (int, error)

//go:linkname unixconnwrite net.sub_unixconnwrite
func unixconnwrite(c *net.UnixConn, b []byte) (int, error) {
	return c.Write(b)
}

//go:linkname UnixConnWrite net.sub_unixconnwrite
//go:noescape
func UnixConnWrite(c *net.UnixConn, b []byte) (int, error)

//go:linkname unixconnsetwritebuffer net.sub_unixconnsetwritebuffer
func unixconnsetwritebuffer(c *net.UnixConn, bytes int) error {
	return c.SetWriteBuffer(bytes)
}

//go:linkname UnixConnSetWriteBuffer net.sub_unixconnsetwritebuffer
//go:noescape
func UnixConnSetWriteBuffer(c *net.UnixConn, bytes int) error

//go:linkname dnsconfigerrorunwrap net.sub_dnsconfigerrorunwrap
func dnsconfigerrorunwrap(e *net.DNSConfigError) error {
	return e.Unwrap()
}

//go:linkname DNSConfigErrorUnwrap net.sub_dnsconfigerrorunwrap
//go:noescape
func DNSConfigErrorUnwrap(e *net.DNSConfigError) error

//go:linkname flagsstring net.sub_flagsstring
func flagsstring(f net.Flags) string {
	return f.String()
}

//go:linkname FlagsString net.sub_flagsstring
//go:noescape
func FlagsString(f net.Flags) string

//go:linkname ipconnreadfrom net.sub_ipconnreadfrom
func ipconnreadfrom(c *net.IPConn, b []byte) (int, net.Addr, error) {
	return c.ReadFrom(b)
}

//go:linkname IPConnReadFrom net.sub_ipconnreadfrom
//go:noescape
func IPConnReadFrom(c *net.IPConn, b []byte) (int, net.Addr, error)

//go:linkname invalidaddrerrortemporary net.sub_invalidaddrerrortemporary
func invalidaddrerrortemporary(e net.InvalidAddrError) bool {
	return e.Temporary()
}

//go:linkname InvalidAddrErrorTemporary net.sub_invalidaddrerrortemporary
//go:noescape
func InvalidAddrErrorTemporary(e net.InvalidAddrError) bool

//go:linkname resolverlookuptxt net.sub_resolverlookuptxt
func resolverlookuptxt(r *net.Resolver, ctx context.Context, name string) ([]string, error) {
	return r.LookupTXT(ctx, name)
}

//go:linkname ResolverLookupTXT net.sub_resolverlookuptxt
//go:noescape
func ResolverLookupTXT(r *net.Resolver, ctx context.Context, name string) ([]string, error)

//go:linkname DialTCP net.DialTCP
//go:noescape
func DialTCP(network string, laddr, raddr *net.TCPAddr) (*net.TCPConn, error)

//go:linkname udpconnreadfromudp net.sub_udpconnreadfromudp
func udpconnreadfromudp(c *net.UDPConn, b []byte) (int, *net.UDPAddr, error) {
	return c.ReadFromUDP(b)
}

//go:linkname UDPConnReadFromUDP net.sub_udpconnreadfromudp
//go:noescape
func UDPConnReadFromUDP(c *net.UDPConn, b []byte) (int, *net.UDPAddr, error)

//go:linkname tcpconnreadfrom net.sub_tcpconnreadfrom
func tcpconnreadfrom(c *net.TCPConn, r io.Reader) (int64, error) {
	return c.ReadFrom(r)
}

//go:linkname TCPConnReadFrom net.sub_tcpconnreadfrom
//go:noescape
func TCPConnReadFrom(c *net.TCPConn, r io.Reader) (int64, error)

//go:linkname LookupCNAME net.LookupCNAME
//go:noescape
func LookupCNAME(host string) (string, error)

//go:linkname LookupTXT net.LookupTXT
//go:noescape
func LookupTXT(name string) ([]string, error)

//go:linkname ipisloopback net.sub_ipisloopback
func ipisloopback(ip net.IP) bool {
	return ip.IsLoopback()
}

//go:linkname IPIsLoopback net.sub_ipisloopback
//go:noescape
func IPIsLoopback(ip net.IP) bool

//go:linkname ipmaskstring net.sub_ipmaskstring
func ipmaskstring(m net.IPMask) string {
	return m.String()
}

//go:linkname IPMaskString net.sub_ipmaskstring
//go:noescape
func IPMaskString(m net.IPMask) string

//go:linkname operrorerror net.sub_operrorerror
func operrorerror(e *net.OpError) string {
	return e.Error()
}

//go:linkname OpErrorError net.sub_operrorerror
//go:noescape
func OpErrorError(e *net.OpError) string

//go:linkname operrortimeout net.sub_operrortimeout
func operrortimeout(e *net.OpError) bool {
	return e.Timeout()
}

//go:linkname OpErrorTimeout net.sub_operrortimeout
//go:noescape
func OpErrorTimeout(e *net.OpError) bool

//go:linkname resolverlookupcname net.sub_resolverlookupcname
func resolverlookupcname(r *net.Resolver, ctx context.Context, host string) (string, error) {
	return r.LookupCNAME(ctx, host)
}

//go:linkname ResolverLookupCNAME net.sub_resolverlookupcname
//go:noescape
func ResolverLookupCNAME(r *net.Resolver, ctx context.Context, host string) (string, error)

//go:linkname tcplisteneraccepttcp net.sub_tcplisteneraccepttcp
func tcplisteneraccepttcp(l *net.TCPListener) (*net.TCPConn, error) {
	return l.AcceptTCP()
}

//go:linkname TCPListenerAcceptTCP net.sub_tcplisteneraccepttcp
//go:noescape
func TCPListenerAcceptTCP(l *net.TCPListener) (*net.TCPConn, error)

//go:linkname unixconnclosewrite net.sub_unixconnclosewrite
func unixconnclosewrite(c *net.UnixConn) error {
	return c.CloseWrite()
}

//go:linkname UnixConnCloseWrite net.sub_unixconnclosewrite
//go:noescape
func UnixConnCloseWrite(c *net.UnixConn) error

//go:linkname unixconnreadfrom net.sub_unixconnreadfrom
func unixconnreadfrom(c *net.UnixConn, b []byte) (int, net.Addr, error) {
	return c.ReadFrom(b)
}

//go:linkname UnixConnReadFrom net.sub_unixconnreadfrom
//go:noescape
func UnixConnReadFrom(c *net.UnixConn, b []byte) (int, net.Addr, error)

//go:linkname unixconnlocaladdr net.sub_unixconnlocaladdr
func unixconnlocaladdr(c *net.UnixConn) net.Addr {
	return c.LocalAddr()
}

//go:linkname UnixConnLocalAddr net.sub_unixconnlocaladdr
//go:noescape
func UnixConnLocalAddr(c *net.UnixConn) net.Addr

//go:linkname addrerrortemporary net.sub_addrerrortemporary
func addrerrortemporary(e *net.AddrError) bool {
	return e.Temporary()
}

//go:linkname AddrErrorTemporary net.sub_addrerrortemporary
//go:noescape
func AddrErrorTemporary(e *net.AddrError) bool

//go:linkname ipmask net.sub_ipmask
func ipmask(ip net.IP, mask net.IPMask) net.IP {
	return ip.Mask(mask)
}

//go:linkname IPMask net.sub_ipmask
//go:noescape
func IPMask(ip net.IP, mask net.IPMask) net.IP

//go:linkname listenconfiglistenpacket net.sub_listenconfiglistenpacket
func listenconfiglistenpacket(lc *net.ListenConfig, ctx context.Context, network, address string) (net.PacketConn, error) {
	return lc.ListenPacket(ctx, network, address)
}

//go:linkname ListenConfigListenPacket net.sub_listenconfiglistenpacket
//go:noescape
func ListenConfigListenPacket(lc *net.ListenConfig, ctx context.Context, network, address string) (net.PacketConn, error)

//go:linkname resolverlookupipaddr net.sub_resolverlookupipaddr
func resolverlookupipaddr(r *net.Resolver, ctx context.Context, host string) ([]net.IPAddr, error) {
	return r.LookupIPAddr(ctx, host)
}

//go:linkname ResolverLookupIPAddr net.sub_resolverlookupipaddr
//go:noescape
func ResolverLookupIPAddr(r *net.Resolver, ctx context.Context, host string) ([]net.IPAddr, error)

//go:linkname tcpconnclose net.sub_tcpconnclose
func tcpconnclose(c *net.TCPConn) error {
	return c.Close()
}

//go:linkname TCPConnClose net.sub_tcpconnclose
//go:noescape
func TCPConnClose(c *net.TCPConn) error

//go:linkname tcpconnsyscallconn net.sub_tcpconnsyscallconn
func tcpconnsyscallconn(c *net.TCPConn) (syscall.RawConn, error) {
	return c.SyscallConn()
}

//go:linkname TCPConnSyscallConn net.sub_tcpconnsyscallconn
//go:noescape
func TCPConnSyscallConn(c *net.TCPConn) (syscall.RawConn, error)

//go:linkname unixconnclose net.sub_unixconnclose
func unixconnclose(c *net.UnixConn) error {
	return c.Close()
}

//go:linkname UnixConnClose net.sub_unixconnclose
//go:noescape
func UnixConnClose(c *net.UnixConn) error

//go:linkname unknownnetworkerrortimeout net.sub_unknownnetworkerrortimeout
func unknownnetworkerrortimeout(e net.UnknownNetworkError) bool {
	return e.Timeout()
}

//go:linkname UnknownNetworkErrorTimeout net.sub_unknownnetworkerrortimeout
//go:noescape
func UnknownNetworkErrorTimeout(e net.UnknownNetworkError) bool

//go:linkname tcpconnfile net.sub_tcpconnfile
func tcpconnfile(c *net.TCPConn) (*os.File, error) {
	return c.File()
}

//go:linkname TCPConnFile net.sub_tcpconnfile
//go:noescape
func TCPConnFile(c *net.TCPConn) (*os.File, error)

//go:linkname tcpconnread net.sub_tcpconnread
func tcpconnread(c *net.TCPConn, b []byte) (int, error) {
	return c.Read(b)
}

//go:linkname TCPConnRead net.sub_tcpconnread
//go:noescape
func TCPConnRead(c *net.TCPConn, b []byte) (int, error)

//go:linkname udpconnwriteto net.sub_udpconnwriteto
func udpconnwriteto(c *net.UDPConn, b []byte, addr net.Addr) (int, error) {
	return c.WriteTo(b, addr)
}

//go:linkname UDPConnWriteTo net.sub_udpconnwriteto
//go:noescape
func UDPConnWriteTo(c *net.UDPConn, b []byte, addr net.Addr) (int, error)

//go:linkname unixconnremoteaddr net.sub_unixconnremoteaddr
func unixconnremoteaddr(c *net.UnixConn) net.Addr {
	return c.RemoteAddr()
}

//go:linkname UnixConnRemoteAddr net.sub_unixconnremoteaddr
//go:noescape
func UnixConnRemoteAddr(c *net.UnixConn) net.Addr

//go:linkname DialUDP net.DialUDP
//go:noescape
func DialUDP(network string, laddr, raddr *net.UDPAddr) (*net.UDPConn, error)

//go:linkname Dial net.Dial
//go:noescape
func Dial(network, address string) (net.Conn, error)

//go:linkname ResolveIPAddr net.ResolveIPAddr
//go:noescape
func ResolveIPAddr(network, address string) (*net.IPAddr, error)

//go:linkname ipconnreadmsgip net.sub_ipconnreadmsgip
func ipconnreadmsgip(c *net.IPConn, b, oob []byte) (int, *net.IPAddr, error) {
	return c.ReadMsgIP(b, oob)
}

//go:linkname IPConnReadMsgIP net.sub_ipconnreadmsgip
//go:noescape
func IPConnReadMsgIP(c *net.IPConn, b, oob []byte) (int, *net.IPAddr, error)

//go:linkname ResolveTCPAddr net.ResolveTCPAddr
//go:noescape
func ResolveTCPAddr(network, address string) (*net.TCPAddr, error)

//go:linkname tcpconncloseread net.sub_tcpconncloseread
func tcpconncloseread(c *net.TCPConn) error {
	return c.CloseRead()
}

//go:linkname TCPConnCloseRead net.sub_tcpconncloseread
//go:noescape
func TCPConnCloseRead(c *net.TCPConn) error

//go:linkname tcpconnclosewrite net.sub_tcpconnclosewrite
func tcpconnclosewrite(c *net.TCPConn) error {
	return c.CloseWrite()
}

//go:linkname TCPConnCloseWrite net.sub_tcpconnclosewrite
//go:noescape
func TCPConnCloseWrite(c *net.TCPConn) error

//go:linkname tcpconnsetwritebuffer net.sub_tcpconnsetwritebuffer
func tcpconnsetwritebuffer(c *net.TCPConn, bytes int) error {
	return c.SetWriteBuffer(bytes)
}

//go:linkname TCPConnSetWriteBuffer net.sub_tcpconnsetwritebuffer
//go:noescape
func TCPConnSetWriteBuffer(c *net.TCPConn, bytes int) error

//go:linkname unknownnetworkerrorerror net.sub_unknownnetworkerrorerror
func unknownnetworkerrorerror(e net.UnknownNetworkError) string {
	return e.Error()
}

//go:linkname UnknownNetworkErrorError net.sub_unknownnetworkerrorerror
//go:noescape
func UnknownNetworkErrorError(e net.UnknownNetworkError) string

//go:linkname dialerdialcontext net.sub_dialerdialcontext
func dialerdialcontext(d *net.Dialer, ctx context.Context, network, address string) (net.Conn, error) {
	return d.DialContext(ctx, network, address)
}

//go:linkname DialerDialContext net.sub_dialerdialcontext
//go:noescape
func DialerDialContext(d *net.Dialer, ctx context.Context, network, address string) (net.Conn, error)

//go:linkname resolverlookupns net.sub_resolverlookupns
func resolverlookupns(r *net.Resolver, ctx context.Context, name string) ([]*net.NS, error) {
	return r.LookupNS(ctx, name)
}

//go:linkname ResolverLookupNS net.sub_resolverlookupns
//go:noescape
func ResolverLookupNS(r *net.Resolver, ctx context.Context, name string) ([]*net.NS, error)

//go:linkname tcpconnsetreadbuffer net.sub_tcpconnsetreadbuffer
func tcpconnsetreadbuffer(c *net.TCPConn, bytes int) error {
	return c.SetReadBuffer(bytes)
}

//go:linkname TCPConnSetReadBuffer net.sub_tcpconnsetreadbuffer
//go:noescape
func TCPConnSetReadBuffer(c *net.TCPConn, bytes int) error

//go:linkname tcplistenersetdeadline net.sub_tcplistenersetdeadline
func tcplistenersetdeadline(l *net.TCPListener, t time.Time) error {
	return l.SetDeadline(t)
}

//go:linkname TCPListenerSetDeadline net.sub_tcplistenersetdeadline
//go:noescape
func TCPListenerSetDeadline(l *net.TCPListener, t time.Time) error

//go:linkname udpconnsyscallconn net.sub_udpconnsyscallconn
func udpconnsyscallconn(c *net.UDPConn) (syscall.RawConn, error) {
	return c.SyscallConn()
}

//go:linkname UDPConnSyscallConn net.sub_udpconnsyscallconn
//go:noescape
func UDPConnSyscallConn(c *net.UDPConn) (syscall.RawConn, error)

//go:linkname udpconnwritetoudp net.sub_udpconnwritetoudp
func udpconnwritetoudp(c *net.UDPConn, b []byte, addr *net.UDPAddr) (int, error) {
	return c.WriteToUDP(b, addr)
}

//go:linkname UDPConnWriteToUDP net.sub_udpconnwritetoudp
//go:noescape
func UDPConnWriteToUDP(c *net.UDPConn, b []byte, addr *net.UDPAddr) (int, error)

//go:linkname unixconnsyscallconn net.sub_unixconnsyscallconn
func unixconnsyscallconn(c *net.UnixConn) (syscall.RawConn, error) {
	return c.SyscallConn()
}

//go:linkname UnixConnSyscallConn net.sub_unixconnsyscallconn
//go:noescape
func UnixConnSyscallConn(c *net.UnixConn) (syscall.RawConn, error)

//go:linkname addrerrorerror net.sub_addrerrorerror
func addrerrorerror(e *net.AddrError) string {
	return e.Error()
}

//go:linkname AddrErrorError net.sub_addrerrorerror
//go:noescape
func AddrErrorError(e *net.AddrError) string

//go:linkname dnsconfigerrortemporary net.sub_dnsconfigerrortemporary
func dnsconfigerrortemporary(e *net.DNSConfigError) bool {
	return e.Temporary()
}

//go:linkname DNSConfigErrorTemporary net.sub_dnsconfigerrortemporary
//go:noescape
func DNSConfigErrorTemporary(e *net.DNSConfigError) bool

//go:linkname ipaddrnetwork net.sub_ipaddrnetwork
func ipaddrnetwork(a *net.IPAddr) string {
	return a.Network()
}

//go:linkname IPAddrNetwork net.sub_ipaddrnetwork
//go:noescape
func IPAddrNetwork(a *net.IPAddr) string

//go:linkname ipconnreadfromip net.sub_ipconnreadfromip
func ipconnreadfromip(c *net.IPConn, b []byte) (int, *net.IPAddr, error) {
	return c.ReadFromIP(b)
}

//go:linkname IPConnReadFromIP net.sub_ipconnreadfromip
//go:noescape
func IPConnReadFromIP(c *net.IPConn, b []byte) (int, *net.IPAddr, error)

//go:linkname tcpaddrnetwork net.sub_tcpaddrnetwork
func tcpaddrnetwork(a *net.TCPAddr) string {
	return a.Network()
}

//go:linkname TCPAddrNetwork net.sub_tcpaddrnetwork
//go:noescape
func TCPAddrNetwork(a *net.TCPAddr) string

//go:linkname udpconnlocaladdr net.sub_udpconnlocaladdr
func udpconnlocaladdr(c *net.UDPConn) net.Addr {
	return c.LocalAddr()
}

//go:linkname UDPConnLocalAddr net.sub_udpconnlocaladdr
//go:noescape
func UDPConnLocalAddr(c *net.UDPConn) net.Addr

//go:linkname unixaddrnetwork net.sub_unixaddrnetwork
func unixaddrnetwork(a *net.UnixAddr) string {
	return a.Network()
}

//go:linkname UnixAddrNetwork net.sub_unixaddrnetwork
//go:noescape
func UnixAddrNetwork(a *net.UnixAddr) string

//go:linkname unixlistenerclose net.sub_unixlistenerclose
func unixlistenerclose(l *net.UnixListener) error {
	return l.Close()
}

//go:linkname UnixListenerClose net.sub_unixlistenerclose
//go:noescape
func UnixListenerClose(l *net.UnixListener) error

//go:linkname dnsconfigerrorerror net.sub_dnsconfigerrorerror
func dnsconfigerrorerror(e *net.DNSConfigError) string {
	return e.Error()
}

//go:linkname DNSConfigErrorError net.sub_dnsconfigerrorerror
//go:noescape
func DNSConfigErrorError(e *net.DNSConfigError) string

//go:linkname ipdefaultmask net.sub_ipdefaultmask
func ipdefaultmask(ip net.IP) net.IPMask {
	return ip.DefaultMask()
}

//go:linkname IPDefaultMask net.sub_ipdefaultmask
//go:noescape
func IPDefaultMask(ip net.IP) net.IPMask

//go:linkname ipconnwrite net.sub_ipconnwrite
func ipconnwrite(c *net.IPConn, b []byte) (int, error) {
	return c.Write(b)
}

//go:linkname IPConnWrite net.sub_ipconnwrite
//go:noescape
func IPConnWrite(c *net.IPConn, b []byte) (int, error)

//go:linkname resolverlookupsrv net.sub_resolverlookupsrv
func resolverlookupsrv(r *net.Resolver, ctx context.Context, service, proto, name string) (string, []*net.SRV, error) {
	return r.LookupSRV(ctx, service, proto, name)
}

//go:linkname ResolverLookupSRV net.sub_resolverlookupsrv
//go:noescape
func ResolverLookupSRV(r *net.Resolver, ctx context.Context, service, proto, name string) (string, []*net.SRV, error)

//go:linkname tcpconnlocaladdr net.sub_tcpconnlocaladdr
func tcpconnlocaladdr(c *net.TCPConn) net.Addr {
	return c.LocalAddr()
}

//go:linkname TCPConnLocalAddr net.sub_tcpconnlocaladdr
//go:noescape
func TCPConnLocalAddr(c *net.TCPConn) net.Addr

//go:linkname tcplistenersyscallconn net.sub_tcplistenersyscallconn
func tcplistenersyscallconn(l *net.TCPListener) (syscall.RawConn, error) {
	return l.SyscallConn()
}

//go:linkname TCPListenerSyscallConn net.sub_tcplistenersyscallconn
//go:noescape
func TCPListenerSyscallConn(l *net.TCPListener) (syscall.RawConn, error)

//go:linkname udpconnsetwritebuffer net.sub_udpconnsetwritebuffer
func udpconnsetwritebuffer(c *net.UDPConn, bytes int) error {
	return c.SetWriteBuffer(bytes)
}

//go:linkname UDPConnSetWriteBuffer net.sub_udpconnsetwritebuffer
//go:noescape
func UDPConnSetWriteBuffer(c *net.UDPConn, bytes int) error
