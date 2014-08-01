package rdp


import "C"


import (
	"unsafe"
	"syscall"
)

var (
	rdpDll = syscall.NewLazyDLL("rdpx64d.dll")
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
	_geterrordesc = rdpDll.NewProc("rdp_geterrordesc")
	_getsystemerror = rdpDll.NewProc("rdp_getsystemerror")
	_getsystemerrordesc = rdpDll.NewProc("rdp_getsystemerrordesc")
	_perfmon = rdpDll.NewProc("rdp_perfmon")
	_getsockstate = rdpDll.NewProc("rdp_getsockstate")

)

var (
	SUCCESS      = newProcInt32(rdpDll,"SUCCESS")//0
	EUNKNOWN     = newProcInt32(rdpDll,"EUNKNOWN")//-1
	ECONNSETUP   = newProcInt32(rdpDll,"ECONNSETUP")//-1000
	ENOSERVER    = newProcInt32(rdpDll,"ENOSERVER")//-1001
	ECONNREJ     = newProcInt32(rdpDll,"ECONNREJ")//-1002
	ESOCKFAIL    = newProcInt32(rdpDll,"ESOCKFAIL")//-1003
	ESECFAIL     = newProcInt32(rdpDll,"ESECFAIL")//-1004
	ECONNFAIL    = newProcInt32(rdpDll,"ECONNFAIL")//-2000
	ECONNLOST    = newProcInt32(rdpDll,"ECONNLOST")//-2001
	ENOCONN      = newProcInt32(rdpDll,"ENOCONN")//-2002
	ERESOURCE    = newProcInt32(rdpDll,"ERESOURCE")//-3000
	ETHREAD      = newProcInt32(rdpDll,"ETHREAD")//-3001
	ENOBUF       = newProcInt32(rdpDll,"ENOBUF")//-3002
	EFILE        = newProcInt32(rdpDll,"EFILE")//-4000
	EINVRDOFF    = newProcInt32(rdpDll,"EINVRDOFF")//-4001
	ERDPERM      = newProcInt32(rdpDll,"ERDPERM")//-4002
	EINVWROFF    = newProcInt32(rdpDll,"EINVWROFF")//-4003
	EWRPERM      = newProcInt32(rdpDll,"EWRPERM")//-4004
	EINVOP       = newProcInt32(rdpDll,"EINVOP")//-5000
	EBOUNDSOCK   = newProcInt32(rdpDll,"EBOUNDSOCK")//-5001
	ECONNSOCK    = newProcInt32(rdpDll,"ECONNSOCK")//-5002
	EINVPARAM    = newProcInt32(rdpDll,"EINVPARAM")//-5003
	EINVSOCK     = newProcInt32(rdpDll,"EINVSOCK")//-5004
	EUNBOUNDSOCK = newProcInt32(rdpDll,"EUNBOUNDSOCK")//-5005
	ENOLISTEN    = newProcInt32(rdpDll,"ENOLISTEN")//-5006
	ERDVNOSERV   = newProcInt32(rdpDll,"ERDVNOSERV")//-5007
	ERDVUNBOUND  = newProcInt32(rdpDll,"ERDVUNBOUND")//-5008
	ESTREAMILL   = newProcInt32(rdpDll,"ESTREAMILL")//-5009
	EDGRAMILL    = newProcInt32(rdpDll,"EDGRAMILL")//-5010
	EDUPLISTEN   = newProcInt32(rdpDll,"EDUPLISTEN")//-5011
	ELARGEMSG    = newProcInt32(rdpDll,"ELARGEMSG")//-5012
	EINVPOLLID   = newProcInt32(rdpDll,"EINVPOLLID")//-5013
	EASYNCFAIL   = newProcInt32(rdpDll,"EASYNCFAIL")//-6000
	EASYNCSND    = newProcInt32(rdpDll,"EASYNCSND")//-6001
	EASYNCRCV    = newProcInt32(rdpDll,"EASYNCRCV")//-6002
	ETIMEOUT     = newProcInt32(rdpDll,"ETIMEOUT")//-6003
	EPEERERR     = newProcInt32(rdpDll,"EPEERERR")//-7000
)

func newProcInt32(dll *syscall.LazyDLL, name string) int32{
	return int32(*((*C.int)(unsafe.Pointer(dll.NewProc(name).Addr()))))
}
