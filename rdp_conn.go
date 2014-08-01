package rdp

import (
	"syscall"
	"net"
	"time"
)

type RDPConn struct {
	Sock RDPSOCKET
	Addr RDPAddr
}

func (c *RDPConn) ok() bool { return c.Sock > 0 }

// Implementation of the Conn interface.

// Read implements the Conn Read method.
func (c *RDPConn) Read(b []byte) (int, error) {
	if !c.ok() {
		return 0, syscall.EINVAL
	}
	n, err := RDP_recvmsg(c.Sock, b)
	return int(n), err
}

// Write implements the Conn Write method.
func (c *RDPConn) Write(b []byte) (int, error) {
	if !c.ok() {
		return 0, syscall.EINVAL
	}
	n, err := RDP_sendmsg(c.Sock, b, 32, false)
	return int(n), err
}

// Close closes the connection.
func (c *RDPConn) Close() error {
	if !c.ok() {
		return syscall.EINVAL
	}
	_, err := RDP_close(c.Sock)
	return err
}

// LocalAddr returns the local network address.
func (c *RDPConn) LocalAddr() net.Addr {
	if !c.ok() {
		return nil
	}
	return &c.Addr
}

// RemoteAddr returns the remote network address.
func (c *RDPConn) RemoteAddr() net.Addr {
	if !c.ok() {
		return nil
	}
	return nil
}

// SetDeadline implements the Conn SetDeadline method.
func (c *RDPConn) SetDeadline(t time.Time) error {
	if !c.ok() {
		return syscall.EINVAL
	}
	return nil
}

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *RDPConn) SetReadDeadline(t time.Time) error {
	if !c.ok() {
		return syscall.EINVAL
	}
	return nil
}

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *RDPConn) SetWriteDeadline(t time.Time) error {
	if !c.ok() {
		return syscall.EINVAL
	}
	return nil
}

// SetReadBuffer sets the size of the operating system's
// receive buffer associated with the connection.
func (c *RDPConn) SetReadBuffer(bytes int) error {
	if !c.ok() {
		return syscall.EINVAL
	}
	return nil
}

// SetWriteBuffer sets the size of the operating system's
// transmit buffer associated with the connection.
func (c *RDPConn) SetWriteBuffer(bytes int) error {
	if !c.ok() {
		return syscall.EINVAL
	}
	return nil
}
