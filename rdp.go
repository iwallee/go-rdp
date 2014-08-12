package rdp

//#include "rdp_go.h"
import "C"
import "unsafe"

import (
	"sync"
)

const (
	RDP_SDK_VERSION = C.RDP_SDK_VERSION
)

//RDPSOCKETSTATUS
const (
	RDPSOCKETSTATUS_INIT      = C.RDPSOCKETSTATUS_INIT
	RDPSOCKETSTATUS_BINDED    = C.RDPSOCKETSTATUS_BINDED
	RDPSOCKETSTATUS_LISTENING = C.RDPSOCKETSTATUS_LISTENING
)

//RDPSESSIONSTAUS
const (
	RDPSESSIONSTATUS_INIT       = C.RDPSESSIONSTATUS_INIT
	RDPSESSIONSTATUS_CONNECTING = C.RDPSESSIONSTATUS_CONNECTING
	RDPSESSIONSTATUS_CONNECTED  = C.RDPSESSIONSTATUS_CONNECTED
)

//RDPERROR
const (
	RDPERROR_SUCCESS = C.RDPERROR_SUCCESS

	RDPERROR_UNKNOWN      = C.RDPERROR_UNKNOWN
	RDPERROR_NOTINIT      = C.RDPERROR_NOTINIT
	RDPERROR_INVALIDPARAM = C.RDPERROR_INVALIDPARAM
	RDPERROR_SYSERROR     = C.RDPERROR_SYSERROR

	RDPERROR_SOCKET_RUNOUT        = C.RDPERROR_SOCKET_RUNOUT
	RDPERROR_SOCKET_INVALIDSOCKET = C.RDPERROR_SOCKET_INVALIDSOCKET
	RDPERROR_SOCKET_BADSTATE      = C.RDPERROR_SOCKET_BADSTATE

	RDPERROR_SOCKET_ONCONNECTNOTSET    = C.RDPERROR_SOCKET_ONCONNECTNOTSET
	RDPERROR_SOCKET_ONACCEPTNOTSET     = C.RDPERROR_SOCKET_ONACCEPTNOTSET
	RDPERROR_SOCKET_ONDISCONNECTNOTSET = C.RDPERROR_SOCKET_ONDISCONNECTNOTSET
	RDPERROR_SOCKET_ONRECVNOTSET       = C.RDPERROR_SOCKET_ONRECVNOTSET
	RDPERROR_SOCKET_ONUDPRECVNOTSET    = C.RDPERROR_SOCKET_ONUDPRECVNOTSET

	RDPERROR_SESSION_INVALIDSESSIONID = C.RDPERROR_SESSION_INVALIDSESSIONID
	RDPERROR_SESSION_BADSTATE         = C.RDPERROR_SESSION_BADSTATE
	RDPERROR_SESSION_CONNTIMEOUT      = C.RDPERROR_SESSION_CONNTIMEOUT
	RDPERROR_SESSION_HEARTBEATTIMEOUT = C.RDPERROR_SESSION_HEARTBEATTIMEOUT
	RDPERROR_SESSION_CONNRESET        = C.RDPERROR_SESSION_CONNRESET
)

//RDPSESSIONSENDFLAG
const (
	RDPSESSIONSENDFLAG_ACK     = C.RDPSESSIONSENDFLAG_ACK
	RDPSESSIONSENDFLAG_INORDER = C.RDPSESSIONSENDFLAG_INORDER
)

const (
	DISCONNECTRESSON_NONE = C.DISCONNECTRESSON_NONE
)

type RDPSOCKET uint32
type RDPSESSIONID uint64

type RDP_on_connect_param struct {
	User_data  uintptr
	Err        int32
	Sock       RDPSOCKET
	Session_id RDPSESSIONID
}

type RDP_on_before_accept_param struct {
	User_data  uintptr
	Sock       RDPSOCKET
	Session_id RDPSESSIONID
	Addr       *RDPAddr
	Buf        []byte
}

type RDP_on_accept_param struct {
	User_data  uintptr
	Sock       RDPSOCKET
	Session_id RDPSESSIONID
	Addr       *RDPAddr
	Buf        []byte
}

type RDP_on_disconnect_param struct {
	User_data  uintptr
	Err        int32
	Reason     uint16
	Sock       RDPSOCKET
	Session_id RDPSESSIONID
}

type RDP_on_recv_param struct {
	User_data  uintptr
	Sock       RDPSOCKET
	Session_id RDPSESSIONID
	Buf        []byte
}

type RDP_on_send_param struct {
	User_data             uintptr
	Err                   int32
	Sock                  RDPSOCKET
	Session_id            RDPSESSIONID
	Local_send_queue_size uint32
	Peer_window_size_     uint32
}

type RDP_on_udp_recv_param struct {
	User_data   uintptr
	Sock        RDPSOCKET
	Addr       *RDPAddr
	Buf         []byte
}

type RDP_on_udp_send_param struct {
	User_data  uintptr
	Err        int32
	Sock       RDPSOCKET
	Session_id RDPSESSIONID
	Addr       *RDPAddr
}

type RDP_on_connect func(param *RDP_on_connect_param)
type RDP_on_before_accept func(param *RDP_on_before_accept_param) bool
type RDP_on_accept func(param *RDP_on_accept_param)
type RDP_on_disconnect func(param *RDP_on_disconnect_param)
type RDP_on_recv func(param *RDP_on_recv_param)
type RDP_on_send func(param *RDP_on_send_param)
type RDP_on_udp_recv func(param *RDP_on_udp_recv_param)
type RDP_on_hash_addr func(addr *RDPAddr) uint32

type RDP_startup_param struct {
	Max_sock        uint8
	Recv_thread_num uint16
	Recv_buf_size   uint32

	On_hash_addr     RDP_on_hash_addr
}

type RDP_socket_create_param struct {
	User_data            uintptr
	Is_v4                bool
	Ack_timeout          uint16
	Heart_beat_timeout   uint16
	Max_send_queue_size  uint16
	Max_recv_queue_size  uint16
	In_session_hash_size uint16

	On_connect       RDP_on_connect
	On_before_accept RDP_on_before_accept
	On_accept        RDP_on_accept
	On_disconnect    RDP_on_disconnect
	On_recv          RDP_on_recv
	On_send          RDP_on_send
	On_udp_recv      RDP_on_udp_recv
}

func RDP_startup(param *RDP_startup_param) int32 {
	starup_param = *param

	var cparam C.struct_rdp_startup_param
	cparam.version = RDP_SDK_VERSION
	cparam.max_sock = (C.ui8)(param.Max_sock)
	cparam.recv_thread_num = (C.ui16)(param.Recv_thread_num)
	cparam.recv_buf_size = (C.ui32)(param.Recv_buf_size)

	cparam.on_hash_addr = nil

	r, _, _ := _startup.Call(uintptr(unsafe.Pointer(&cparam)))
	return int32(r)
}
func RDP_startup_get_param() (*RDP_startup_param, int32) {
	return &starup_param, 0
}
func RDP_cleanup() int32 {
	r, _, _ := _cleanup.Call()
	return int32(r)
}

func RDP_getsyserror() int32 {
	r, _, _ := _getsyserror.Call()
	return int32(r)
}
func RDP_getsyserrordesc(err int32) (string, int32) {
	var errdesc [1024]byte
	len := len(errdesc)
	var ed string
	r, _, _ := _getsyserrordesc.Call(uintptr(err),
		uintptr(unsafe.Pointer(&errdesc[0])),
		uintptr(unsafe.Pointer(&len)),
	)
	if int32(r) >= 0 {
		ed = string(errdesc[:len])
	}
	return ed, int32(r)
}

func RDP_socket_create(param *RDP_socket_create_param) (RDPSOCKET, int32) {
	var sock C.RDPSOCKET
	var cparam C.struct_rdp_socket_create_param
	if param.Is_v4 {
		cparam.is_v4 = 1
	} else {
		cparam.is_v4 = 0
	}
	cparam.userdata = unsafe.Pointer(param.User_data)
	cparam.ack_timeout = (C.ui16)(param.Ack_timeout)
	cparam.heart_beat_timeout = (C.ui16)(param.Heart_beat_timeout)
	cparam.max_send_queue_size = (C.ui16)(param.Max_send_queue_size)
	cparam.max_recv_queue_size = (C.ui16)(param.Max_recv_queue_size)
	cparam.in_session_hash_size = (C.ui16)(param.In_session_hash_size)

	/*cparam.on_connect = C.__on_connect
	cparam.on_before_accept = C.__on_before_accept
	cparam.on_accept = C.__on_accept
	cparam.on_disconnect = C.__on_disconnect
	cparam.on_recv = C.__on_recv
	cparam.on_send = C.__on_send
	cparam.on_udp_recv = C.__on_udp_recv*/
	r := C.socket_create(unsafe.Pointer(_socket_create.Addr()),
		&cparam,
		&sock)

	if sock >= 0 {
		mutex_lock.Lock()
		socket_create_param[RDPSOCKET(sock)] = param
		mutex_lock.Unlock()
	}
	return RDPSOCKET(sock), int32(r)
}
func RDP_socket_get_create_param(sock RDPSOCKET) (*RDP_socket_create_param, int32) {
	param, ok := socket_create_param[sock]
	if ok {
		return param, RDPERROR_SUCCESS
	}

	return nil, RDPERROR_SOCKET_INVALIDSOCKET
}
func RDP_socket_get_state(sock RDPSOCKET) (int32, int32) {
	var state int
	r, _, _ := _socket_get_state.Call(uintptr(sock),
		uintptr(unsafe.Pointer(&state)))
	return int32(state), int32(r)
}
func RDP_socket_close(sock RDPSOCKET) int32 {
	r, _, _ := _socket_close.Call(uintptr(sock))
	mutex_lock.Lock()
	delete (socket_create_param, sock)
	mutex_lock.Unlock()
	return int32(r)
}
func RDP_socket_bind(sock RDPSOCKET, addr *RDPAddr) int32 {
	r, _, _ := _socket_bind.Call(uintptr(sock),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(addr.Port))
	return int32(r)
}
func RDP_socket_listen(sock RDPSOCKET) int32 {
	r, _, _ := _socket_listen.Call(uintptr(sock))
	return int32(r)
}

func RDP_socket_connect(sock RDPSOCKET, addr *RDPAddr, timeout int32, data []byte) (RDPSESSIONID, int32) {
	ip := C.CString(string(addr.IP))
	var session_id RDPSESSIONID
	r, _, _ := _socket_connect.Call(uintptr(sock),
		uintptr(unsafe.Pointer(ip)),
		uintptr(addr.Port),
		uintptr(unsafe.Pointer(&session_id)))
	return session_id, int32(r)
}
func RDP_socket_recv(timeout int32) (int32) {
	r, _, _ := _socket_recv.Call(uintptr(timeout))
	return int32(r)
}
func RDP_session_close(sock RDPSOCKET, session_id RDPSESSIONID, reason int32) int32 {
	r, _, _ := _session_close.Call(uintptr(sock),
		uintptr(sock),
		uintptr(session_id),
		uintptr(reason))
	return int32(r)
}
func RDP_session_get_state(sock RDPSOCKET, session_id RDPSESSIONID) (int32, int32) {
	var state int
	r, _, _ := _socket_get_state.Call(uintptr(sock),
		uintptr(session_id),
		uintptr(unsafe.Pointer(&state)))
	return int32(state), int32(r)
}

func RDP_session_send(sock RDPSOCKET, session_id RDPSESSIONID, data []byte) int32 {
	r, _, _ := _session_send.Call(uintptr(sock),
		uintptr(sock),
		uintptr(session_id),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(0))
	return int32(r)
}
func RDP_session_is_in_come(session_id RDPSESSIONID) bool {
	r, _, _ := _session_is_in_come.Call(uintptr(session_id))
	return r == 1
}
func RDP_udp_send(sock RDPSOCKET, addr *RDPAddr, data []byte) int32 {
	ip := C.CString(string(addr.IP))
	r, _, _ := _udp_send.Call(uintptr(sock),
		uintptr(sock),
		uintptr(unsafe.Pointer(ip)),
		uintptr(addr.Port),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)))
	return int32(r)
}

/////////////////////////////////////////////////////////////////
var mutex_lock sync.Mutex
var starup_param RDP_startup_param
var socket_create_param map[RDPSOCKET] *RDP_socket_create_param = make(map[RDPSOCKET] *RDP_socket_create_param)

//export on_connect
func on_connect(param *C.struct_rdp_on_connect_param) {
	mutex_lock.Lock()
	s , ok := socket_create_param[RDPSOCKET(param.sock)]
	if ok {
		p := RDP_on_connect_param{
			User_data:  uintptr(param.userdata),
			Err:        int32(param.err),
			Sock:       RDPSOCKET(param.sock),
			Session_id: RDPSESSIONID(param.session_id),
		}
		s.On_connect(&p)
	}
	mutex_lock.Unlock()
}

//export on_disconnect
func on_disconnect(param *C.struct_rdp_on_disconnect_param) {
	mutex_lock.Lock()
	s , ok := socket_create_param[RDPSOCKET(param.sock)]
	if ok {
		p := RDP_on_disconnect_param{
			User_data:  uintptr(param.userdata),
			Err:        int32(param.err),
			Reason:     uint16(param.reason),
			Sock:       RDPSOCKET(param.sock),
			Session_id: RDPSESSIONID(param.session_id),
		}
		s.On_disconnect(&p)
	}
	mutex_lock.Unlock()
}

//export on_before_accept
func on_before_accept(param *C.struct_rdp_on_before_accept_param) bool {
	var ret bool = false
	mutex_lock.Lock()
	s , ok := socket_create_param[RDPSOCKET(param.sock)]
	if ok {
		if s.On_before_accept != nil {
			addr, n := addr_to(param.addr, param.addrlen)
			if n >= 0 {
				p := RDP_on_before_accept_param{
					User_data:  uintptr(param.userdata),
					Sock:       RDPSOCKET(param.sock),
					Session_id: RDPSESSIONID(param.session_id),
					Addr:       addr,
					Buf:        C.GoBytes(unsafe.Pointer(param.buf), C.int(param.buf_len)) ,
				}
				ret = s.On_before_accept(&p)
			}
		}
	}
	mutex_lock.Unlock()
	return ret
}

//export on_accept
func on_accept(param *C.struct_rdp_on_accept_param) {
	mutex_lock.Lock()
	s , ok := socket_create_param[RDPSOCKET(param.sock)]
	if ok {
		if s.On_accept != nil {
			addr, n := addr_to(param.addr, param.addrlen)
			if n >= 0 {
				p := RDP_on_accept_param{
					User_data:  uintptr(param.userdata),
					Sock:       RDPSOCKET(param.sock),
					Session_id: RDPSESSIONID(param.session_id),
					Addr:       addr,
					Buf:        C.GoBytes(unsafe.Pointer(param.buf), C.int(param.buf_len)) ,
				}
				s.On_accept(&p)
			}
		}
	}
	mutex_lock.Unlock()
}

//export on_recv
func on_recv(param *C.struct_rdp_on_recv_param) {
	mutex_lock.Lock()
	s , ok := socket_create_param[RDPSOCKET(param.sock)]
	if ok {
		p := RDP_on_recv_param{
			User_data:  uintptr(param.userdata),
			Sock:       RDPSOCKET(param.sock),
			Session_id: RDPSESSIONID(param.session_id),
			Buf:        C.GoBytes(unsafe.Pointer(param.buf), C.int(param.buf_len)) ,
		}
		s.On_recv(&p)
	}
	mutex_lock.Unlock()
}

//export on_send
func on_send(param *C.struct_rdp_on_send_param) {
	mutex_lock.Lock()
	s , ok := socket_create_param[RDPSOCKET(param.sock)]
	if ok {
		if s.On_send != nil {
			p := RDP_on_send_param{
				User_data:             uintptr(param.userdata),
				Err:                   int32(param.err),
				Sock:                  RDPSOCKET(param.sock),
				Session_id:            RDPSESSIONID(param.session_id),
				Local_send_queue_size: uint32(param.local_send_queue_size),
				Peer_window_size_:     uint32(param.peer_window_size),
			}
			s.On_send(&p)
		}
	}
	mutex_lock.Unlock()
}

//export on_udp_recv
func on_udp_recv(param *C.struct_rdp_on_udp_recv_param) {
	mutex_lock.Lock()
	s , ok := socket_create_param[RDPSOCKET(param.sock)]
	if ok {
		if s.On_udp_recv != nil {
			addr, n := addr_to(param.addr, param.addrlen)
			if n >= 0 {
				p := RDP_on_udp_recv_param{
					User_data:  uintptr(param.userdata),
					Sock:       RDPSOCKET(param.sock),
					Addr:       addr,
					Buf:        C.GoBytes(unsafe.Pointer(param.buf), C.int(param.buf_len)) ,
				}

				s.On_udp_recv(&p)
			}
		}
	}
	mutex_lock.Unlock()
}

//export on_hash_addr
func on_hash_addr(addr *C.struct_sockaddr, addrlen C.ui32) uint32 {
	if starup_param.On_hash_addr != nil {
		add, n := addr_to(addr, addrlen)
		if n >= 0 {
			return starup_param.On_hash_addr(add)
		}
	}
	panic("bad hash_addr")
	return 1
}
func addr_to(addr *C.struct_sockaddr, addrlen C.ui32) (*RDPAddr, int32) {
	var ip [64]byte
	len := len(ip)
	var port int
	var addr1 *RDPAddr
	var addrlen1 = int32(addrlen)
	r, _, _ := _addr_to.Call(uintptr(unsafe.Pointer(addr)),
		uintptr(addrlen),
		uintptr(unsafe.Pointer(&ip[0])),
		uintptr(unsafe.Pointer(&addrlen1)),
		uintptr(unsafe.Pointer(&port)))
	if int32(r) >= 0 {
		addr1 = &RDPAddr{
			IP:   ip[:len],
			Port: port,
		}
	}
	return addr1, int32(r)
}

