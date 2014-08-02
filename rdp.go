package rdp

//#include "rdp_def.h"
import "C"


import (
	"unsafe"
	"syscall"
)

//RDPEPOLLOPT
const (
	RDP_EPOLL_IN  = C.RDP_EPOLL_IN
	RDP_EPOLL_OUT = C.RDP_EPOLL_OUT
	RDP_EPOLL_ERR = C.RDP_EPOLL_ERR
)

//RDPSTATUS
const (
	INIT       = C.INIT
	OPENED     = C.OPENED
	LISTENING  = C.LISTENING
	CONNECTING = C.CONNECTING
	CONNECTED  = C.CONNECTED
	BROKEN     = C.BROKEN
	CLOSING    = C.CLOSING
	CLOSED     = C.CLOSED
	NONEXIST   = C.NONEXIST
)

//RDPSOCKOPT
const (
	RDP_MSS        = C.RDP_MSS        // the Maximum Transfer Unit
	RDP_SNDSYN     = C.RDP_SNDSYN     // if sending is blocking
	RDP_RCVSYN     = C.RDP_RCVSYN     // if receiving is blocking
	RDP_CC         = C.RDP_CC         // custom congestion control algorithm
	RDP_FC         = C.RDP_FC         // Flight flag size (window size)
	RDP_SNDBUF     = C.RDP_SNDBUF     // maximum buffer in sending queue
	RDP_RCVBUF     = C.RDP_RCVBUF     // RDP receiving buffer size
	RDP_LINGER     = C.RDP_LINGER     // waiting for unsent data when closing
	UDP_SNDBUF     = C.UDP_SNDBUF     // UDP sending buffer size
	UDP_RCVBUF     = C.UDP_RCVBUF     // UDP receiving buffer size
	RDP_MAXMSG     = C.RDP_MAXMSG     // maximum datagram message size
	RDP_MSGTTL     = C.RDP_MSGTTL     // time-to-live of a datagram message
	RDP_RENDEZVOUS = C.RDP_RENDEZVOUS // rendezvous connection mode
	RDP_SNDTIMEO   = C.RDP_SNDTIMEO   // send() timeout
	RDP_RCVTIMEO   = C.RDP_RCVTIMEO   // recv() timeout
	RDP_REUSEADDR  = C.RDP_REUSEADDR  // reuse an existing port or create a new one
	RDP_MAXBW      = C.RDP_MAXBW      // maximum bandwidth (bytes per second) that the connection can use
	RDP_STATE      = C.RDP_STATE      // current socket state, see RDPSTATUS, read only
	RDP_EVENT      = C.RDP_EVENT      // current avalable events associated with the socket
	RDP_SNDDATA    = C.RDP_SNDDATA    // size of data in the sending buffer
	RDP_RCVDATA    = C.RDP_RCVDATA    // size of data available for recv
)

type RDPPerfMon C.RDPPerfMon

type RDPSOCKET int32

func RDP_startup() (int32, error) {
	r, _, _ := _startup.Call()
	return int32(r), nil
}
func RDP_cleanup() (int32, error) {
	r, _, _ := _cleanup.Call()
	return int32(r), nil
}
func RDP_socket(af int32) (RDPSOCKET, error) {
	r, _, _ := _socket.Call(uintptr(af))
	return RDPSOCKET(r), nil
}
func RDP_bind(u RDPSOCKET, af int32, addr *RDPAddr) (int32, error) {
	r, _, _ := _bind.Call(uintptr(u),
		uintptr(af),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(addr.Port))
	return int32(r), nil
}
func RDP_listen(u RDPSOCKET, backlog int) (int32, error) {
	r, _, _ := _listen.Call(uintptr(u),
		uintptr(backlog))
	return int32(r), nil
}
func RDP_accept(u RDPSOCKET, af int32) (RDPSOCKET, *RDPAddr, error) {
	var ip [64]byte
	len := len(ip)
	var addr RDPAddr
	r, _, _ := _accept.Call(uintptr(u),
		uintptr(af),
		uintptr(unsafe.Pointer(&ip[0])),
		uintptr(unsafe.Pointer(&len)),
		uintptr(unsafe.Pointer(&addr.Port)))
	if RDPSOCKET(r) >= 0 {
		if af == syscall.AF_INET {
			addr.IP = parseIPv4(string(ip[:len]))
		} else if af == syscall.AF_INET6 {
			addr.IP, addr.Zone = parseIPv6(string(ip[:len]), true)
		}
	}

	return RDPSOCKET(r), &addr, nil
}
func RDP_connect(u RDPSOCKET, af int32, addr *RDPAddr) (int32, error) {
	r, _, _ := _connect.Call(uintptr(u),
		uintptr(af),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(len(addr.IP)),
		uintptr(addr.Port))
	return int32(r), nil
}
func RDP_close(u RDPSOCKET) (int32, error) {
	r, _, _ := _close.Call(uintptr(u))
	return int32(r), nil
}
func RDP_getpeername(u RDPSOCKET, af int32) (int32, *RDPAddr, error) {
	var ip [64]byte
	len := len(ip)
	var addr RDPAddr
	r, _, _ := _getpeername.Call(uintptr(u),
		uintptr(af),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(unsafe.Pointer(&len)),
		uintptr(unsafe.Pointer(&addr.Port)))
	if int32(r) >= 0 {
		if af == syscall.AF_INET {
			addr.IP = parseIPv4(string(ip[:len]))
		} else if af == syscall.AF_INET6 {
			addr.IP, addr.Zone = parseIPv6(string(ip[:len]), true)
		}
	}

	return int32(r), &addr, nil
}
func RDP_getsockname(u RDPSOCKET, af int32) (int32, *RDPAddr, error) {
	var ip [64]byte
	len := len(ip)
	var addr RDPAddr
	r, _, _ := _getsockname.Call(uintptr(u),
		uintptr(af),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(unsafe.Pointer(&len)),
		uintptr(unsafe.Pointer(&addr.Port)))
	if int32(r) >= 0 {
		if af == syscall.AF_INET {
			addr.IP = parseIPv4(string(ip[:len]))
		} else if af == syscall.AF_INET6 {
			addr.IP, addr.Zone = parseIPv6(string(ip[:len]), true)
		}
	}

	return int32(r), &addr, nil
}
func RDP_getsockopt(u RDPSOCKET, level int32, optname int32, optval interface{}, optlen int32) (int32, error) {
	r, _, _ := _getsockopt.Call(uintptr(u),
		uintptr(optname),
		uintptr(optlen))
	return int32(r), nil
}
func RDP_setsockopt(u RDPSOCKET, level int32, optname int32, optval interface{}, optlen int32) (int32, error) {
	r, _, _ := _setsockopt.Call(uintptr(u),
		uintptr(optname),
		uintptr(optlen))
	return int32(r), nil
}
func RDP_sendmsg(u RDPSOCKET, buf[]byte, ttl int32, inorder bool) (int32, error) {
	var _inorder int
	if inorder {
		_inorder = 1
	} else {
		_inorder = 0
	}
	r, _, _ := _sendmsg.Call(uintptr(u),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)), uintptr(ttl),
		uintptr(_inorder))
	return int32(r), nil
}
func RDP_recvmsg(u RDPSOCKET, buf[]byte) (int32, error) {
	r, _, _ := _recvmsg.Call(uintptr(u),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)))
	return int32(r), nil
}
func RDP_epoll_create() (int32, error) {
	r, _, err := _epoll_create.Call()
	return int32(r), err
}
func RDP_epoll_add_usock(eid int32, u RDPSOCKET, events *int32/*=nil*/) (int32, error) {
	r, _, _ := _epoll_add_usock.Call(uintptr(eid),
		uintptr(u),
		uintptr(unsafe.Pointer(events)))
	return int32(r), nil
}
func RDP_epoll_remove_usock(eid int32, u RDPSOCKET) (int32, error) {
	r, _, _ := _epoll_remove_usock.Call(uintptr(eid),
		uintptr(u))
	return int32(r), nil
}
func RDP_epoll_wait(eid int32,
	readfds []RDPSOCKET, rnum* int32,
	writefds []RDPSOCKET, wnum *int32,
	msTimeOut int64) (int32, error) {
	r, _, _ := _epoll_wait.Call(uintptr(eid),
		uintptr(unsafe.Pointer(&readfds[0])),
		uintptr(unsafe.Pointer(rnum)),
		uintptr(unsafe.Pointer(&writefds[0])),
		uintptr(unsafe.Pointer(wnum)),
		uintptr(msTimeOut),
	)
	return int32(r), nil
}
func RDP_epoll_release(eid int32) (int32, error) {
	r, _, _ := _epoll_release.Call(uintptr(eid))
	return int32(r), nil
}
func RDP_geterrordesc(err int) (int32, string, error) {
	var errdesc [1024]byte
	len := len(errdesc)
	var ed string
	r, _, _ := _geterrordesc.Call(uintptr(err),
		uintptr(unsafe.Pointer(&errdesc[0])),
		uintptr(unsafe.Pointer(&len)),
	)
	if int32(r) >= 0 {
		ed = string(errdesc[:len])
	}
	return int32(r), ed, nil
}
func RDP_getsystemerror() (int32, error) {
	r, _, _ := _getsystemerror.Call()
	return int32(r), nil
}
func RDP_getsystemerrordesc(err int) (int32, string, error) {
	var errdesc [1024]byte
	len := len(errdesc)
	var ed string
	r, _, _ := _getsystemerrordesc.Call(uintptr(err),
		uintptr(unsafe.Pointer(&errdesc[0])),
		uintptr(unsafe.Pointer(&len)),
	)
	if int32(r) >= 0 {
		ed = string(errdesc[:len])
	}
	return int32(r), ed, nil
}
func RDP_perfmon(u *RDPSOCKET, perf *RDPPerfMon, clear bool/*=true*/) (int32, error) {
	return 0, nil
}
func RDP_getsockstate(u RDPSOCKET) (int32, error) {
	r, _, _ := _getsockstate.Call(uintptr(u))
	return int32(r), nil
}


func init() {
	RDP_startup()
}
