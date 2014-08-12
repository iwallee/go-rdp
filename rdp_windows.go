package rdp

import "C"


import (
	"unsafe"
	"syscall"
)

var (
	rdpDll = syscall.NewLazyDLL("C:/Develop/code/iwallee/rdp2/rdp/win/bin64/Debug/rdp.dll")
)
var (
	_startup           = rdpDll.NewProc("rdp_startup")
	_startup_get_param = rdpDll.NewProc("rdp_startup_get_param")
	_cleanup           = rdpDll.NewProc("rdp_cleanup")
	_getsyserror       = rdpDll.NewProc("rdp_getsyserror")
	_getsyserrordesc   = rdpDll.NewProc("rdp_getsyserrordesc")


	_socket_create           = rdpDll.NewProc("rdp_socket_create")
	_socket_get_create_param = rdpDll.NewProc("rdp_socket_get_create_param")
	_socket_get_state        = rdpDll.NewProc("rdp_socket_get_state")
	_socket_close            = rdpDll.NewProc("rdp_socket_close")
	_socket_bind             = rdpDll.NewProc("rdp_socket_bind")
	_socket_listen           = rdpDll.NewProc("rdp_socket_listen")
	_socket_connect          = rdpDll.NewProc("rdp_socket_connect")
	_socket_recv             = rdpDll.NewProc("rdp_socket_recv")

	_session_close      = rdpDll.NewProc("rdp_session_close")
	_session_get_state  = rdpDll.NewProc("rdp_session_get_state")
	_session_send       = rdpDll.NewProc("rdp_session_send")
	_session_is_in_come = rdpDll.NewProc("rdp_session_is_in_come")

	_udp_send           = rdpDll.NewProc("rdp_udp_send")

	_addr_to            = rdpDll.NewProc("rdp_addr_to")

)

var (
	//SUCCESS      = newProcInt32(rdpDll, "SUCCESS")      //0
)

func newProcInt32(dll *syscall.LazyDLL, name string) int32 {
	return int32(*((*C.int)(unsafe.Pointer(dll.NewProc(name).Addr()))))
}
