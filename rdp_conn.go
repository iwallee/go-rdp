package udt

import (
	"syscall"
	"net"
	"time"
)

type UDTConn struct {
	Sock UDTSOCKET
	Addr UDTAddr
}

func (c *UDTConn) ok() bool { return c.Sock != UDTSOCKET(INVALID_SOCK) }

// Implementation of the Conn interface.

// Read implements the Conn Read method.
func (c *UDTConn) Read(b []byte) (int, error) {
	if !c.ok() {
		return 0, syscall.EINVAL
	}
	n, err := UDT_recvmsg(c.Sock, b)
	return int(n), err
}

// Write implements the Conn Write method.
func (c *UDTConn) Write(b []byte) (int, error) {
	if !c.ok() {
		return 0, syscall.EINVAL
	}
	n, err := UDT_sendmsg(c.Sock, b, 32, false)
	return int(n), err
}

// Close closes the connection.
func (c *UDTConn) Close() error {
	if !c.ok() {
		return syscall.EINVAL
	}
	_, err := UDT_close(c.Sock)
	return err
}

// LocalAddr returns the local network address.
func (c *UDTConn) LocalAddr() net.Addr {
	if !c.ok() {
		return nil
	}
	return &c.Addr
}

// RemoteAddr returns the remote network address.
func (c *UDTConn) RemoteAddr() net.Addr {
	if !c.ok() {
		return nil
	}
	return nil
}

// SetDeadline implements the Conn SetDeadline method.
func (c *UDTConn) SetDeadline(t time.Time) error {
	if !c.ok() {
		return syscall.EINVAL
	}
	return nil
}

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *UDTConn) SetReadDeadline(t time.Time) error {
	if !c.ok() {
		return syscall.EINVAL
	}
	return nil
}

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *UDTConn) SetWriteDeadline(t time.Time) error {
	if !c.ok() {
		return syscall.EINVAL
	}
	return nil
}

// SetReadBuffer sets the size of the operating system's
// receive buffer associated with the connection.
func (c *UDTConn) SetReadBuffer(bytes int) error {
	if !c.ok() {
		return syscall.EINVAL
	}
	return nil
}

// SetWriteBuffer sets the size of the operating system's
// transmit buffer associated with the connection.
func (c *UDTConn) SetWriteBuffer(bytes int) error {
	if !c.ok() {
		return syscall.EINVAL
	}
	return nil
}
