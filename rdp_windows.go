package rdp


import "C"


import (
	"syscall"
)

var (
	rdpDll = syscall.NewLazyDLL("C:/Develop/code/iwallee/kara/src/rdp/rdp4/win/bin/bin/rdpd.dll")
)
var (
	_startup = rdpDll.NewProc("rdp_startup")
	_cleanup = rdpDll.NewProc("rdp_cleanup")
	_socket = rdpDll.NewProc("rdp_socket")
	_bind = rdpDll.NewProc("rdp_bind")
	_bind2 = rdpDll.NewProc("rdp_bind2")
	_listen = rdpDll.NewProc("rdp_listen")
	_accept = rdpDll.NewProc("rdp_accept")
	_connect = rdpDll.NewProc("rdp_connect")
	_close = rdpDll.NewProc("rdp_close")
	_getpeername = rdpDll.NewProc("rdp_getpeername")
	_getsockname = rdpDll.NewProc("rdp_getsockname")
	_getsockopt = rdpDll.NewProc("rdp_getsockopt")
	_setsockopt = rdpDll.NewProc("rdp_setsockopt")
	_sendmsg = rdpDll.NewProc("rdp_sendmsg")
	_recvmsg = rdpDll.NewProc("rdp_recvmsg")
	_epoll_create = rdpDll.NewProc("rdp_epoll_create")
	_epoll_add_usock = rdpDll.NewProc("rdp_epoll_add_usock")
	_epoll_add_ssock = rdpDll.NewProc("rdp__epoll_add_ssock")
	_epoll_remove_usock = rdpDll.NewProc("rdp_epoll_remove_usock")
	_epoll_remove_ssock = rdpDll.NewProc("rdp_epoll_remove_ssock")
	_epoll_wait = rdpDll.NewProc("rdp_epoll_wait")
	_epoll_release = rdpDll.NewProc("rdp_epoll_release")
	_getlasterror = rdpDll.NewProc("rdp_getlasterror")
	_perfmon = rdpDll.NewProc("rdp_perfmon")
	_getsockstate = rdpDll.NewProc("rdp_getsockstate")

)


