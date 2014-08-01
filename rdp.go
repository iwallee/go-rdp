package rdp

//#include ".h"
//import "C"


import (
	"unsafe"
	"syscall"
	"strings"
	"strconv"
	"net"
)

//EPOLLOPT
const (
	RDP_EPOLL_IN  = 0x1
	RDP_EPOLL_OUT = 0x4
	RDP_EPOLL_ERR = 0x8
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
	RDP_MSS = iota // the Maximum Transfer Unit
	RDP_SNDSYN     // if sending is blocking
	RDP_RCVSYN     // if receiving is blocking
	RDP_CC         // custom congestion control algorithm
	RDP_FC         // Flight flag size (window size)
	RDP_SNDBUF     // maximum buffer in sending queue
	RDP_RCVBUF     // RDP receiving buffer size
	RDP_LINGER     // waiting for unsent data when closing
	UDP_SNDBUF     // UDP sending buffer size
	UDP_RCVBUF     // UDP receiving buffer size
	RDP_MAXMSG     // maximum datagram message size
	RDP_MSGTTL     // time-to-live of a datagram message
	RDP_RENDEZVOUS // rendezvous connection mode
	RDP_SNDTIMEO   // send() timeout
	RDP_RCVTIMEO   // recv() timeout
	RDP_REUSEADDR  // reuse an existing port or create a new one
	RDP_MAXBW      // maximum bandwidth (bytes per second) that the connection can use
	RDP_STATE      // current socket state, see RDPSTATUS, read only
	RDP_EVENT      // current avalable events associated with the socket
	RDP_SNDDATA    // size of data in the sending buffer
	RDP_RCVDATA    // size of data available for recv
)

type PerfMon struct {
	// global measurements
	msTimeStamp        int64 // time since the RDP entity is started, in milliseconds
	pktSentTotal       int64 // total number of sent data packets, including retransmissions
	pktRecvTotal       int64 // total number of received packets
	pktSndLossTotal    int32 // total number of lost packets (sender side)
	pktRcvLossTotal    int32 // total number of lost packets (receiver side)
	pktRetransTotal    int32 // total number of retransmitted packets
	pktSentACKTotal    int32 // total number of sent ACK packets
	pktRecvACKTotal    int32 // total number of received ACK packets
	pktSentNAKTotal    int32 // total number of sent NAK packets
	pktRecvNAKTotal    int32 // total number of received NAK packets
	usSndDurationTotal int64 // total time duration when RDP is sending data (idle time exclusive)

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
	byteAvailSndBuf     int32   // available RDP sender buffer size
	byteAvailRcvBuf     int32   // available RDP receiver buffer size
}




type SYSSOCKET int32
type RDPSOCKET int32
type UDPSOCKET SYSSOCKET


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
func RDP_bind2(u RDPSOCKET, udpsock UDPSOCKET) (int32, error) {
	r, _, _ := _bind2.Call(uintptr(u),
		uintptr(udpsock))
	return int32(r), nil
}
func RDP_listen(u RDPSOCKET, backlog int) (int32, error) {
	r, _, _ := _listen.Call(uintptr(u),
		uintptr(backlog))
	return int32(r), nil
}
func RDP_accept(u RDPSOCKET, af int32, addr *RDPAddr) (RDPSOCKET, error) {
	var ip [64]byte
	len := len(ip)
	r, _, _ := _accept.Call(uintptr(u),
		uintptr(af),
		uintptr(unsafe.Pointer(&ip[0])),
		uintptr(unsafe.Pointer(&len)),
		uintptr(unsafe.Pointer(&addr.Port)))
	if RDPSOCKET(r) >= 0{
		if af == syscall.AF_INET {
			what := strings.Split(string(ip[:len]), ".")
			i0 ,_ := strconv.Atoi(what[0])
			i1 ,_ := strconv.Atoi(what[1])
			i2 ,_ := strconv.Atoi(what[2])
			i3 ,_ := strconv.Atoi(what[3])
			addr.IP = net.IPv4(byte(i0),byte(i1), byte(i2), byte(i3))
		} else if af == syscall.AF_INET6{
			addr.IP, addr.Zone = parseIPv6(string(addr.IP), true)
		}
	}


	return RDPSOCKET(r), nil
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
func RDP_getpeername(u RDPSOCKET, af int32, addr *RDPAddr) (int32, error) {
	var ip [64]byte
	len := len(ip)
	r, _, _ := _getpeername.Call(uintptr(u),
		uintptr(af),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(unsafe.Pointer(&len)),
		uintptr(unsafe.Pointer(&addr.Port)))
	if RDPSOCKET(r) >= 0{
		if af == syscall.AF_INET {
			what := strings.Split(string(ip[:len]), ".")
			i0 ,_ := strconv.Atoi(what[0])
			i1 ,_ := strconv.Atoi(what[1])
			i2 ,_ := strconv.Atoi(what[2])
			i3 ,_ := strconv.Atoi(what[3])
			addr.IP = net.IPv4(byte(i0),byte(i1), byte(i2), byte(i3))
		} else if af == syscall.AF_INET6{
			addr.IP, addr.Zone = parseIPv6(string(addr.IP), true)
		}
	}
	return int32(r), nil
}
func RDP_getsockname(u RDPSOCKET, af int32, addr *RDPAddr) (int32, error) {
	var ip [64]byte
	len := len(ip)
	r, _, _ := _getsockname.Call(uintptr(u),
		uintptr(af),
		uintptr(unsafe.Pointer(&addr.IP[0])),
		uintptr(unsafe.Pointer(&len)),
		uintptr(unsafe.Pointer(&addr.Port)))
	if RDPSOCKET(r) >= 0{
		if af == syscall.AF_INET {
			what := strings.Split(string(ip[:len]), ".")
			i0 ,_ := strconv.Atoi(what[0])
			i1 ,_ := strconv.Atoi(what[1])
			i2 ,_ := strconv.Atoi(what[2])
			i3 ,_ := strconv.Atoi(what[3])
			addr.IP = net.IPv4(byte(i0),byte(i1), byte(i2), byte(i3))
		} else if af == syscall.AF_INET6{
			addr.IP, addr.Zone = parseIPv6(string(addr.IP), true)
		}
	}
	return int32(r), nil
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
func RDP_epoll_add_ssock(eid int32, s SYSSOCKET, events *int32/*=nil*/) (int32, error) {
	r, _, _ := _epoll_add_ssock.Call(uintptr(eid),
		uintptr(s),
		uintptr(unsafe.Pointer(events)))
	return int32(r), nil
}
func RDP_epoll_remove_usock(eid int32, u RDPSOCKET) (int32, error) {
	r, _, _ := _epoll_remove_usock.Call(uintptr(eid),
		uintptr(u))
	return int32(r), nil
}
func RDP_epoll_remove_ssock(eid int32 , s SYSSOCKET) (int32, error) {
	r, _, _ := _epoll_remove_ssock.Call(uintptr(eid),
		uintptr(s))
	return int32(r), nil
}
func RDP_epoll_wait(eid int32,
	readfds []RDPSOCKET, rnum* int32,
	writefds []RDPSOCKET, wnum *int32,
	msTimeOut int64,
	lrfds []SYSSOCKET/*=nil*/, lrnum *int32,
	lwfds []SYSSOCKET/*=nil*/, lwnum *int32) (int32, error) {
	r, _, _ := _epoll_wait.Call(uintptr(eid),
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
func RDP_epoll_release(eid int32) (int32, error) {
	r, _, _ := _epoll_release.Call(uintptr(eid))
	return int32(r), nil
}
func RDP_geterrordesc()(int32, error){
	//r, _, _ := _geterrordesc.Call()
	//return int32(r), nil
	return 0, nil
}
func RDP_getsystemerror()(int32, error){
	r, _, _ := _getsystemerror.Call()
	return int32(r), nil
}
func RDP_getsystemerrordesc()(int32, error){
	//r, _, _ := _geterrordesc.Call()
	//return int32(r), nil
	return 0, nil
}
func RDP_perfmon(u *RDPSOCKET, perf *PerfMon, clear bool/*=true*/) (int32, error) {
	return 0, nil
}
func RDP_getsockstate(u RDPSOCKET) (int32, error) {
	r, _, _ := _getsockstate.Call(uintptr(u))
	return int32(r), nil
}




func init() {
	RDP_startup()
}
