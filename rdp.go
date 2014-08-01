package udt

//#include "udt_go.h"
import "C"


import (
	"unsafe"
)

//EPOLLOPT
const (
	UDT_EPOLL_IN  = 0x1
	UDT_EPOLL_OUT = 0x4
	UDT_EPOLL_ERR = 0x8
)

//SOCKSTATUS
const (
	INIT = 1
	OPENED
	LISTENING
	CONNECTING
	CONNECTED
	BROKEN
	CLOSING
	CLOSED
	NONEXIST
)

//SOCKOPT
const (
	UDT_MSS = iota // the Maximum Transfer Unit
	UDT_SNDSYN     // if sending is blocking
	UDT_RCVSYN     // if receiving is blocking
	UDT_CC         // custom congestion control algorithm
	UDT_FC         // Flight flag size (window size)
	UDT_SNDBUF     // maximum buffer in sending queue
	UDT_RCVBUF     // UDT receiving buffer size
	UDT_LINGER     // waiting for unsent data when closing
	UDP_SNDBUF     // UDP sending buffer size
	UDP_RCVBUF     // UDP receiving buffer size
	UDT_MAXMSG     // maximum datagram message size
	UDT_MSGTTL     // time-to-live of a datagram message
	UDT_RENDEZVOUS // rendezvous connection mode
	UDT_SNDTIMEO   // send() timeout
	UDT_RCVTIMEO   // recv() timeout
	UDT_REUSEADDR  // reuse an existing port or create a new one
	UDT_MAXBW      // maximum bandwidth (bytes per second) that the connection can use
	UDT_STATE      // current socket state, see UDTSTATUS, read only
	UDT_EVENT      // current avalable events associated with the socket
	UDT_SNDDATA    // size of data in the sending buffer
	UDT_RCVDATA    // size of data available for recv
)

type PerfMon struct {
	// global measurements
	msTimeStamp        int64 // time since the UDT entity is started, in milliseconds
	pktSentTotal       int64 // total number of sent data packets, including retransmissions
	pktRecvTotal       int64 // total number of received packets
	pktSndLossTotal    int32 // total number of lost packets (sender side)
	pktRcvLossTotal    int32 // total number of lost packets (receiver side)
	pktRetransTotal    int32 // total number of retransmitted packets
	pktSentACKTotal    int32 // total number of sent ACK packets
	pktRecvACKTotal    int32 // total number of received ACK packets
	pktSentNAKTotal    int32 // total number of sent NAK packets
	pktRecvNAKTotal    int32 // total number of received NAK packets
	usSndDurationTotal int64 // total time duration when UDT is sending data (idle time exclusive)

	// local measurements
	pktSent       int64   // number of sent data packets, including retransmissions
	pktRecv       int64   // number of received packets
	pktSndLoss    int32   // number of lost packets (sender side)
	pktRcvLoss    int32   // number of lost packets (receiver side)
	pktRetrans    int32   // number of retransmitted packets
	pktSentACK    int32   // number of sent ACK packets
	pktRecvACK    int32   // number of received ACK packets
	pktSentNAK    int32   // number of sent NAK packets
	pktRecvNAK    int32   // number of received NAK packets
	mbpsSendRate  float64 // sending rate in Mb/s
	mbpsRecvRate  float64 // receiving rate in Mb/s
	usSndDuration int64   // busy sending time (i.e., idle time exclusive)

	// instant measurements
	usPktSndPeriod      float64 // packet sending period, in microseconds
	pktFlowWindow       int32   // flow window size, in number of packets
	pktCongestionWindow int32   // congestion window size, in number of packets
	pktFlightSize       int32   // number of packets on flight
	msRTT               float64 // RTT, in milliseconds
	mbpsBandwidth       float64 // estimated bandwidth, in Mb/s
	byteAvailSndBuf     int32   // available UDT sender buffer size
	byteAvailRcvBuf     int32   // available UDT receiver buffer size
}

const (
	SUCCESS      = 0
	ECONNSETUP   = 1000
	ENOSERVER    = 1001
	ECONNREJ     = 1002
	ESOCKFAIL    = 1003
	ESECFAIL     = 1004
	ECONNFAIL    = 2000
	ECONNLOST    = 2001
	ENOCONN      = 2002
	ERESOURCE    = 3000
	ETHREAD      = 3001
	ENOBUF       = 3002
	EFILE        = 4000
	EINVRDOFF    = 4001
	ERDPERM      = 4002
	EINVWROFF    = 4003
	EWRPERM      = 4004
	EINVOP       = 5000
	EBOUNDSOCK   = 5001
	ECONNSOCK    = 5002
	EINVPARAM    = 5003
	EINVSOCK     = 5004
	EUNBOUNDSOCK = 5005
	ENOLISTEN    = 5006
	ERDVNOSERV   = 5007
	ERDVUNBOUND  = 5008
	ESTREAMILL   = 5009
	EDGRAMILL    = 5010
	EDUPLISTEN   = 5011
	ELARGEMSG    = 5012
	EINVPOLLID   = 5013
	EASYNCFAIL   = 6000
	EASYNCSND    = 6001
	EASYNCRCV    = 6002
	ETIMEOUT     = 6003
	EPEERERR     = 7000
	EUNKNOWN     = -1
)

type SOCKET struct {
	sock     int32
	af       int32
	socktype int32
}

type SYSSOCKET int32
type UDTSOCKET int32
type UDPSOCKET SYSSOCKET

const (
	INVALID_SOCK  int32 = -1
	ERROR         int32 = -1
	EXCEPTION     int32 = -2
)

func UDT_startup() (int32, error) {
	r, _, _ := _startup.Call()
	return int32(r), nil
}
func UDT_cleanup() (int32, error) {
	r, _, _ := _cleanup.Call()
	return int32(r), nil
}
func UDT_socket(af int32, socktype int32, protocol int32) (UDTSOCKET, error) {
	r, _, _ := _socket.Call(uintptr(af),
		uintptr(socktype),
		uintptr(protocol))
	return UDTSOCKET(r), nil
}
func UDT_bind(u UDTSOCKET, addr *UDTAddr) (int32, error) {
	r, _, _ := _bind.Call(uintptr(u),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(len(addr.IP)),
		uintptr(addr.Port))
	return int32(r), nil
}
func UDT_bind2(u UDTSOCKET, udpsock UDPSOCKET) (int32, error) {
	r, _, _ := _bind2.Call(uintptr(u),
		uintptr(udpsock))
	return int32(r), nil
}
func UDT_listen(u UDTSOCKET, backlog int) (int32, error) {
	r, _, _ := _listen.Call(uintptr(u),
		uintptr(backlog))
	return int32(r), nil
}
func UDT_accept(u UDTSOCKET, addr *UDTAddr) (UDTSOCKET, error) {
	r, _, _ := _accept.Call(uintptr(u),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(len(addr.IP)),
		uintptr(unsafe.Pointer(&addr.Port)))
	return UDTSOCKET(r), nil
}
func UDT_connect(u UDTSOCKET, addr *UDTAddr) (int32, error) {
	r, _, _ := _connect.Call(uintptr(u),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(len(addr.IP)),
		uintptr(addr.Port))
	return int32(r), nil
}
func UDT_close(u UDTSOCKET) (int32, error) {
	r, _, _ := _close.Call(uintptr(u))
	return int32(r), nil
}
func UDT_getpeername(u UDTSOCKET, addr *UDTAddr) (int32, error) {
	r, _, _ := _getpeername.Call(uintptr(u),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(len(addr.IP)),
		uintptr(unsafe.Pointer(&addr.Port)))
	return int32(r), nil
}
func UDT_getsockname(u UDTSOCKET, addr *UDTAddr) (int32, error) {
	r, _, _ := _getsockname.Call(uintptr(u),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(len(addr.IP)),
		uintptr(unsafe.Pointer(&addr.Port)))
	return int32(r), nil
}
func UDT_getsockopt(u UDTSOCKET, level int32, optname int32, optval interface{}, optlen int32) (int32, error) {
	r, _, _ := _getsockopt.Call(uintptr(u),
		uintptr(optname),
		uintptr(optlen))
	return int32(r), nil
}
func UDT_setsockopt(u UDTSOCKET, level int32, optname int32, optval interface{}, optlen int32) (int32, error) {
	r, _, _ := _setsockopt.Call(uintptr(u),
		uintptr(optname),
		uintptr(optlen))
	return int32(r), nil
}
func UDT_send(u UDTSOCKET, buf[]byte, flags int32) (int32, error) {
	r, _, _ := _send.Call(uintptr(u),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)), uintptr(flags))
	return int32(r), nil
}
func UDT_recv(u UDTSOCKET, buf[]byte, flags int32) (int32, error) {
	r, _, _ := _recv.Call(uintptr(u),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)), uintptr(flags))
	return int32(r), nil
}
func UDT_sendmsg(u UDTSOCKET, buf[]byte, ttl int32, inorder bool) (int32, error) {
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
func UDT_recvmsg(u UDTSOCKET, buf[]byte) (int32, error) {
	r, _, _ := _recvmsg.Call(uintptr(u),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)))
	return int32(r), nil
}
func UDT_sendfile2(u UDTSOCKET, path string, offset *int64, size int64, block /*=364000*/int32) (int64, error) {
	s := C.CString(path)
	r, _, _ := _sendfile2.Call(uintptr(u),
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(offset)),
		uintptr(size),
		uintptr(block))
	return int64(r), nil
}
func UDT_recvfile2(u UDTSOCKET, path string, offset *int64, size int64, block /*=7280000*/int32) (int64, error) {
	s := C.CString(path)
	r, _, _ := _recvfile2.Call(uintptr(u),
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(offset)),
		uintptr(size),
		uintptr(block))
	return int64(r), nil
}
func UDT_epoll_create() (int32, error) {
	r, _, err := _epoll_create.Call()
	return int32(r), err
}
func UDT_epoll_add_usock(eid int32, u UDTSOCKET, events *int32/*=nil*/) (int32, error) {
	r, _, _ := _epoll_add_usock.Call(uintptr(eid),
		uintptr(u),
		uintptr(unsafe.Pointer(events)))
	return int32(r), nil
}
func UDT_epoll_add_ssock(eid int32, s SYSSOCKET, events *int32/*=nil*/) (int32, error) {
	r, _, _ := _epoll_add_ssock.Call(uintptr(eid),
		uintptr(s),
		uintptr(unsafe.Pointer(events)))
	return int32(r), nil
}
func UDT_epoll_remove_usock(eid int32, u UDTSOCKET) (int32, error) {
	r, _, _ := _epoll_remove_usock.Call(uintptr(eid),
		uintptr(u))
	return int32(r), nil
}
func UDT_epoll_remove_ssock(eid int32 , s SYSSOCKET) (int32, error) {
	r, _, _ := _epoll_remove_ssock.Call(uintptr(eid),
		uintptr(s))
	return int32(r), nil
}
func UDT_epoll_wait2(eid int32,
	readfds []UDTSOCKET, rnum* int32,
	writefds []UDTSOCKET, wnum *int32,
	msTimeOut int64,
	lrfds []SYSSOCKET/*=nil*/, lrnum *int32,
	lwfds []SYSSOCKET/*=nil*/, lwnum *int32) (int32, error) {
	r, _, _ := _epoll_wait2.Call(uintptr(eid),
		uintptr(unsafe.Pointer(&readfds[0])),
		uintptr(unsafe.Pointer(rnum)),
		uintptr(unsafe.Pointer(&writefds[0])),
		uintptr(unsafe.Pointer(wnum)),
		uintptr(msTimeOut),
		uintptr(unsafe.Pointer(&lrfds[0])),
		uintptr(unsafe.Pointer(lrnum)),
		uintptr(unsafe.Pointer(&lwfds[0])),
		uintptr(unsafe.Pointer(lwnum)),
	)
	return int32(r), nil
}
func UDT_epoll_release(eid int32) (int32, error) {
	r, _, _ := _epoll_release.Call(uintptr(eid))
	return int32(r), nil
}
func UDT_getlasterror() *UDTException {
	return nil
}
func UDT_perfmon(u *UDTSOCKET, perf *PerfMon, clear bool/*=true*/) (int32, error) {
	return 0, nil
}
func UDT_getsockstate(u UDTSOCKET) (int32, error) {
	r, _, _ := _getsockstate.Call(uintptr(u))
	return int32(r), nil
}

type UDTException struct {

}

func (self *UDTException) GetErrorMessage() (string) {
	return ""
}
func (self *UDTException) GetErrorCode() (int) {
	return 0
}
func (self *UDTException) Clear() {

}

func init() {
	UDT_startup()
}
