package rdp

/////////////////////////////////////////////////////////////////////////////////
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

/////////////////////////////////////////////////////////////////
