package udt


import "C"


import (
	"syscall"
)

var (
	udt = syscall.NewLazyDLL("C:/Develop/code/iwallee/kara/src/udt/udt4/win/bin/bin/udtd.dll")
)
var (
	_startup = udt.NewProc("udt_startup")
	_cleanup = udt.NewProc("udt_cleanup")
	_socket = udt.NewProc("udt_socket")
	_bind = udt.NewProc("udt_bind")
	_bind2 = udt.NewProc("udt_bind2")
	_listen = udt.NewProc("udt_listen")
	_accept = udt.NewProc("udt_accept")
	_connect = udt.NewProc("udt_connect")
	_close = udt.NewProc("udt_close")
	_getpeername = udt.NewProc("udt_getpeername")
	_getsockname = udt.NewProc("udt_getsockname")
	_getsockopt = udt.NewProc("udt_getsockopt")
	_setsockopt = udt.NewProc("udt_setsockopt")
	_send = udt.NewProc("udt_send")
	_recv = udt.NewProc("udt_recv")
	_sendmsg = udt.NewProc("udt_sendmsg")
	_recvmsg = udt.NewProc("udt_recvmsg")
	_sendfile2 = udt.NewProc("udt_sendfile2")
	_recvfile2 = udt.NewProc("udt_recvfile2")
	_epoll_create = udt.NewProc("udt_epoll_create")
	_epoll_add_usock = udt.NewProc("udt_epoll_add_usock")
	_epoll_add_ssock = udt.NewProc("udt__epoll_add_ssock")
	_epoll_remove_usock = udt.NewProc("udt_epoll_remove_usock")
	_epoll_remove_ssock = udt.NewProc("udt_epoll_remove_ssock")
	_epoll_wait2 = udt.NewProc("udt_epoll_wait2")
	_epoll_release = udt.NewProc("udt_epoll_release")
	_getlasterror = udt.NewProc("udt_getlasterror")
	_perfmon = udt.NewProc("udt_perfmon")
	_getsockstate = udt.NewProc("udt_getsockstate")

)


