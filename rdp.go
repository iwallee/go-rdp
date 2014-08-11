package rdp

//#include "rdp_go.h"


import "C"


import "unsafe"

const (
	RDP_SDK_VERSION = 0x00010001
)

//RDPSOCKETSTATUS
const (
	RDPSOCKETSTATUS_INIT      = C.enum_RDPSOCKETSTATUS_INIT
	RDPSOCKETSTATUS_BINDED    = C.enum_RDPSOCKETSTATUS_BINDED
	RDPSOCKETSTATUS_LISTENING = C.enum_RDPSOCKETSTATUS_LISTENING
)

//RDPSESSIONSTAUS
const (
	RDPSESSIONSTATUS_INIT       = C.enum_RDPSESSIONSTATUS_INIT
	RDPSESSIONSTATUS_CONNECTING = C.enum_RDPSESSIONSTATUS_CONNECTING
	RDPSESSIONSTATUS_CONNECTED  = C.enum_RDPSESSIONSTATUS_CONNECTED
)

//RDPERROR
const (
	RDPERROR_SUCCESS = C.enum_RDPERROR_SUCCESS

	RDPERROR_UNKNOWN      = C.enum_RDPERROR_UNKNOWN
	RDPERROR_NOTINIT      = C.enum_RDPERROR_NOTINIT
	RDPERROR_INVALIDPARAM = C.enum_RDPERROR_INVALIDPARAM
	RDPERROR_SYSERROR     = C.enum_RDPERROR_SYSERROR

	RDPERROR_SOCKET_RUNOUT        = C.enum_RDPERROR_SOCKET_RUNOUT
	RDPERROR_SOCKET_INVALIDSOCKET = C.enum_RDPERROR_SOCKET_INVALIDSOCKET
	RDPERROR_SOCKET_BADSTATE      = C.enum_RDPERROR_SOCKET_BADSTATE

	RDPERROR_SOCKET_ONCONNECTNOTSET    = C.enum_RDPERROR_SOCKET_ONCONNECTNOTSET
	RDPERROR_SOCKET_ONACCEPTNOTSET     = C.enum_RDPERROR_SOCKET_ONACCEPTNOTSET
	RDPERROR_SOCKET_ONDISCONNECTNOTSET = C.enum_RDPERROR_SOCKET_ONDISCONNECTNOTSET
	RDPERROR_SOCKET_ONRECVNOTSET       = C.enum_RDPERROR_SOCKET_ONRECVNOTSET
	RDPERROR_SOCKET_ONUDPRECVNOTSET    = C.enum_RDPERROR_SOCKET_ONUDPRECVNOTSET

	RDPERROR_SESSION_INVALIDSESSIONID = C.enum_RDPERROR_SESSION_INVALIDSESSIONID
	RDPERROR_SESSION_BADSTATE         = C.enum_RDPERROR_SESSION_BADSTATE
	RDPERROR_SESSION_CONNTIMEOUT      = C.enum_RDPERROR_SESSION_CONNTIMEOUT
	RDPERROR_SESSION_HEARTBEATTIMEOUT = C.enum_RDPERROR_SESSION_HEARTBEATTIMEOUT
	RDPERROR_SESSION_CONNRESET        = C.enum_RDPERROR_SESSION_CONNRESET
)

//RDPSESSIONSENDFLAG
const (
	RDPSESSIONSENDFLAG_ACK     = C.enum_RDPSESSIONSENDFLAG_ACK
	RDPSESSIONSENDFLAG_INORDER = C.enum_RDPSESSIONSENDFLAG_INORDER
)

const (
	disconnect_reason_none = 0
)



type RDPSOCKET uint32
type RDPSESSIONID uint64


type RDP_on_connect_param struct{
	Err        int32
	Sock       RDPSOCKET
	Session_id RDPSESSIONID
}

type RDP_on_before_accept_param struct{
	Sock       RDPSOCKET
	Session_id RDPSESSIONID
	Addr       RDPAddr
	Buf        []byte
}

type RDP_on_before_accept_param RDP_on_accept_param


type RDP_on_disconnect_param struct{
	Err        int32
	Reason     uint16
	Sock       RDPSOCKET
	Session_id RDPSESSIONID
}

type RDP_on_recv_param struct{
	Sock       RDPSOCKET
	Session_id RDPSESSIONID
	Buf        []byte
}

type RDP_on_send_param struct{
	Err                   int32
	Sock                  RDPSOCKET
	Session_id            RDPSESSIONID
	Local_send_queue_size uint32
	Peer_window_size_     uint32
}

type RDP_on_udp_recv_param struct{
	Sock    RDPSOCKET
	Addr    RDPAddr
	Buf     []byte;
}

type RDP_on_udp_send_param struct{
	Err        int32
	Sock       RDPSOCKET
	Session_id RDPSESSIONID
	Addr       RDPAddr
}

type RDP_on_connect interface {
	On_connect(param *RDP_on_connect_param)
}
type RDP_on_before_accept interface {
	On_before_accept(param *RDP_on_before_accept_param) bool
}
type RDP_on_accept interface {
	On_accept(param *RDP_on_accept_param)
}
type RDP_on_disconnect interface {
	On_disconnect(param *RDP_on_disconnect_param)
}
type RDP_on_recv interface {
	On_recv(param *RDP_on_recv_param)
}
type RDP_on_send interface {
	On_send(param *RDP_on_send_param)
}
type RDP_on_udp_recv interface {
	On_udp_recv(param *RDP_on_udp_recv_param)
}
type RDP_on_hash_addr interface {
	On_hash_addr(addr *RDPAddr) uint32
}


type RDP_startup_param struct{
	Version         uint32
	Max_sock        uint8
	Recv_thread_num uint16
	Recv_buf_size   uint32

	On_connect       RDP_on_connect
	On_before_accept RDP_on_before_accept
	On_accept        RDP_on_accept
	On_disconnect    RDP_on_disconnect
	On_recv          RDP_on_recv
	On_send          RDP_on_send
	On_udp_recv      RDP_on_udp_recv
	On_hash_addr     RDP_on_hash_addr
}

type RDP_socket_create_param struct{
	Is_v4                bool
	Ack_timeout          uint16
	Heart_beat_timeout   uint16
	Max_send_queue_size  uint16
	Max_recv_queue_size  uint16
	In_session_hash_size uint16
}

func RDP_startup(param *RDP_startup_param) (int32) {
	starup_param = param
	starup_param.Version = RDP_SDK_VERSION

	var cparam C.struct_rdp_startup_param
	cparam.version = param.Version
	cparam.max_sock = param.Max_sock
	cparam.recv_thread_num = param.Recv_thread_num
	cparam.recv_buf_size = param.Recv_buf_size

	cparam.on_connect = on_connect
	cparam.on_before_accept = on_before_accept
	cparam.on_accept = on_accept
	cparam.on_disconnect = on_disconnect
	cparam.on_recv = on_recv
	cparam.on_send = on_send
	cparam.on_udp_recv = on_udp_recv
	cparam.on_hash_addr = on_hash_addr

	r, _, _ := _startup.Call(unsafe.Pointer(&cparam))
	if int32(r) >= 0{
		starup_param.Max_sock = cparam.max_sock
		param.Version = starup_param.Version
		param.Max_sock = starup_param.Max_sock
	}
	return int32(r)
}
func RDP_startup_get_param() (*RDP_startup_param, int32) {
	return &starup_param, 0
}
func RDP_cleanup() (int32) {
	r, _, _ := _cleanup.Call()
	return int32(r)
}

func RDP_getsyserror() (int32) {
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

func RDP_socket_create(param* RDP_socket_create_param) (RDPSOCKET, int32) {
	var cparam C.struct_rdp_socket_create_param
	cparam.is_v4 = param.Is_v4
	cparam.ack_timeout = param.Ack_timeout
	cparam.heart_beat_timeout = param.Heart_beat_timeout
	cparam.max_send_queue_size = param.Max_send_queue_size
	cparam.max_recv_queue_size = param.Max_recv_queue_size
	cparam.in_session_hash_size = param.In_session_hash_size

	r, _, _ := _socket_create.Call(unsafe.Pointer(&cparam))
	if int32(r) >= 0{
		param.Ack_timeout = cparam.ack_timeout
		param.Heart_beat_timeout = cparam.heart_beat_timeout
		param.Max_send_queue_size = cparam.max_send_queue_size
		param.Max_recv_queue_size = cparam.max_recv_queue_size
		param.In_session_hash_size = cparam.in_session_hash_size
	}
	return RDPSOCKET(r), int32(r)
}
func RDP_socket_get_create_param(sock RDPSOCKET) (*RDP_socket_create_param, int32) {
	var cparam C.struct_rdp_socket_create_param
	r, _, _ := _socket_get_create_param.Call(uintptr(sock),
		unsafe.Pointer(&cparam))

	var param *RDP_socket_create_param
	if int32(r) >= 0 {
		param = &RDP_socket_create_param{
			Is_v4: cparam.is_v4,
			Ack_timeout : cparam.ack_timeout,
			Heart_beat_timeout: cparam.heart_beat_timeout,
			Max_send_queue_size : cparam.max_send_queue_size,
			Max_recv_queue_size : cparam.max_recv_queue_size,
			In_session_hash_size : cparam.in_session_hash_size,
		}
	}

	return param, int32(r)
}
func RDP_socket_get_state(sock RDPSOCKET) (int32, int32) {
	var state int
	r, _, _ := _socket_get_state.Call(uintptr(sock),
		uintptr(unsafe.Pointer(&state)))
	return int32(state), int32(r)
}
func RDP_socket_close(sock RDPSOCKET) (int32) {
	r, _, _ := _close.Call(uintptr(sock))
	return int32(r)
}
func RDP_socket_bind(sock RDPSOCKET, addr *RDPAddr) (int32) {
	r, _, _ := _socket_bind.Call(uintptr(sock),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(addr.Port))
	return int32(r)
}
func RDP_socket_listen(sock RDPSOCKET) (int32) {
	r, _, _ := _socket_listen.Call(uintptr(sock))
	return int32(r)
}

func RDP_socket_connect(sock RDPSOCKET, addr *RDPAddr, timeout int32, data[]byte) (RDPSESSIONID, int32) {
	ip = C.CString(string(addr.IP))
	r, _, _ := _socket_connect.Call(uintptr(sock),
		uintptr(unsafe.Pointer(ip)),
		uintptr(addr.Port))
	return int32(r)
}
func RDP_session_close(sock RDPSOCKET, session_id RDPSESSIONID, reason int32) (int32) {
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

func RDP_session_send(sock RDPSOCKET, session_id RDPSESSIONID, data[]byte) (int32) {
	r, _, _ := _session_send.Call(uintptr(sock),
		uintptr(sock),
		uintptr(session_id),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(0))
	return int32(r)
}
func RDP_session_is_in_come(session_id RDPSESSIONID) (bool) {
	r, _, _ := _session_is_in_come.Call(uintptr(session_id))
	return int32(r)
}
func RDP_udp_send(sock RDPSOCKET, addr *RDPAddr, data[]byte) (int32) {
	ip = C.CString(string(addr.IP))
	r, _, _ := _udp_send.Call(uintptr(sock),
		uintptr(sock),
		uintptr(session_id),
		uintptr(unsafe.Pointer(ip)),
		uintptr(addr.Port),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)))
	return int32(r)
}


/////////////////////////////////////////////////////////////////
var starup_param RDP_startup_param

//export on_connect
//export on_disconnect
//export on_before_accept
//export on_accept
//export on_recv
//export on_send
//export on_udp_recv
//export on_hash_addr

func on_connect(param *C.struct_rdp_on_connect_param) {
	p := RDP_on_connect_param {
		Err : param.err,
		Sock   : param.sock,
		Session_id : param.session_id,
	}
	callback.On_connect(&p)
}
func on_disconnect(param *C.struct_rdp_on_disconnect_param) {
	p := RDP_on_disconnect_param {
		Err     : param.err,
		Reason     : param.reason,
		Sock       : param.sock,
		Session_id : param.session_id,
	}
	callback.On_disconnect(&p)
}
func on_before_accept(param *C.struct_rdp_on_before_accept_param) {
	if starup_param.On_before_accept {
		n, addr := addr_to(param.addr, param.addrlen)
		if n >= 0 {
			p := RDP_on_before_accept_param {
				Sock      : param.sock,
				Session_id: param.session_id,
				Addr      : param.addr,
				Buf       : param.buf,
			}
			callback.On_before_accept(&p)
		}
	}
}
func on_accept(param *C.struct_rdp_on_accept_param) {
	if starup_param.On_accept {
		n, addr := addr_to(param.addr, param.addrlen)
		if n >= 0{
			p := RDP_on_accept_param {
				Sock      : param.sock,
				Session_id: param.session_id,
				Addr      : param.addr,
				Buf       : param.buf,
			}
			callback.On_accept(&p)
		}
	}
}
func on_recv(param *C.struct_rdp_on_recv_param) {
	p := RDP_on_recv_param {
		Sock      : param.sock,
		Session_id: param.session_id,
		Buf       : param.buf,
	}
	callback.On_recv(&p)
}
func on_send(param* C.struct_rdp_on_send_param) {
	if starup_param.On_send {
		p := RDP_on_send_param {
			Err       : param.err,
			Sock      : param.sock,
			Session_id: param.session_id,
			Local_send_queue_size      : param.local_send_queue_size,
			Peer_window_size_       : param.peer_window_size_,
		}
		callback.On_send(&p)
	}
}
func on_udp_recv(param *C.struct_rdp_on_udp_recv_param) {
	if starup_param.On_udp_recv {
		n, addr := addr_to(param.addr, param.addrlen)
		if n >= 0 {
			p := RDP_on_udp_recv_param {
				Sock      : param.sock,
				Addr      : param.addr,
				Buf       : param.buf,
			}
			callback.On_udp_recv(&p)
		}
	}
}
func on_hash_addr(addr *C.struct_sockaddr, addrlen uint32) uint32 {
	if starup_param.On_hash_addr {
		n, add := addr_to(addr, addrlen)
		if n >= 0 {
			return callback.On_hash_addr(add)
		}
	}
	panic("bad hash_addr")
}
func addr_to(addr *C.struct_sockaddr, addrlen uint32) (*RDPAddr, int32) {
	var ip [64]byte
	len := len(ip)
	var port int
	var addr *RDPAddr
	r, _, _ := _addr_to.Call(uintptr(sock),
		uintptr(unsafe.Pointer(addr)),
		uintptr(addrlen),
		uintptr(unsafe.Pointer(&ip[0])),
		uintptr(len),
		uintptr(unsafe.Pointer(&port)))
	if int32(r) >= 0 {
		addr = &RDPAddr{
			IP:string(ip[:len]),
			Port:port,
		}
	}
	return addr, int32(r)
}

