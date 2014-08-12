package rdp

import (
	"testing"
	"log"
	"net"
)

const (
	op_accept  = 0
	op_disconn = 1
	op_recv    = 2
)
func On_connect(param *RDP_on_connect_param) {

}
func On_before_accept(param *RDP_on_before_accept_param) bool {
	if len(param.Buf) != len("hi,wallee") {
		return false
	}
	if string(param.Buf) != "hi,wallee" {
		return false
	}
	return true
}
func On_accept(param *RDP_on_accept_param) {
	wd := workChannelData{
		operation  :op_accept,
		sock       :param.Sock,
		session_id :param.Session_id,
	}
	work_channel <- wd
}
func On_disconnect(param *RDP_on_disconnect_param) {
	wd := workChannelData{
		operation  :op_disconn,
		sock       :param.Sock,
		session_id :param.Session_id,
	}
	work_channel <- wd
}
func On_recv(param *RDP_on_recv_param) {
	wd := workChannelData{
		operation  :op_recv,
		sock       :param.Sock,
		session_id :param.Session_id,
		data       :make([]byte, len(param.Buf)),
	}
	copy(wd.data, param.Buf)
	work_channel <- wd
}
func On_send(param *RDP_on_send_param) {

}
func On_udp_recv(param *RDP_on_udp_recv_param) {

}
func On_hash_addr(addr *RDPAddr) uint32 {
	return 0
}

type workChannelData struct {
	operation  int
	sock       RDPSOCKET
	session_id RDPSESSIONID
	data       []byte
	addr       *RDPAddr
}

var work_channel chan workChannelData

func workProc() {
	for {
		w := <-work_channel
		if w.operation == op_accept {

		} else if w.operation == op_disconn{

		} else if w.operation == op_recv {

		}
	}
}


func TestRun(t *testing.T) {
	log.Println("rdp server start")

	var startup_param RDP_startup_param

	startup_param.Max_sock = 1
	startup_param.Recv_thread_num = 1
	startup_param.Recv_buf_size = 1024*4
	startup_param.On_connect = On_connect
	startup_param.On_before_accept = On_before_accept
	startup_param.On_accept = On_accept
	startup_param.On_disconnect = On_disconnect
	startup_param.On_recv = On_recv
	//startup_param.On_send = On_send
	//startup_param.On_udp_recv = On_udp_recv
	//startup_param.On_hash_addr = On_hash_addr


	n := RDP_startup(&startup_param)
	if n < 0 {
		log.Println("RDP_startup failed", n)
		return
	}

	var socket_param RDP_socket_create_param
	socket_param.Is_v4 = true
	socket_param.Ack_timeout = 100
	socket_param.Heart_beat_timeout = 180
	socket_param.Max_send_queue_size = 0
	socket_param.Max_recv_queue_size = 0
	socket_param.In_session_hash_size = 1024



	var sock RDPSOCKET
	sock, n = RDP_socket_create(&socket_param)
	if n < 0 {
		log.Println("RDP_socket_create failed", n)
		return
	}

	addr := RDPAddr {
		IP: net.IPv4(0, 0, 0, 0),
		Port: 9000,
	}

	n = RDP_socket_bind(sock, &addr)
	if n < 0 {
		log.Println("RDP_socket_bind failed", n)
		return
	}

	n = RDP_socket_listen(sock)
	if n < 0 {
		log.Println("RDP_listen failed", n)
		return
	}
	log.Println("rdp server on ", addr.IP, ":", addr.Port)

	work_channel = make(chan workChannelData, 1)
	workProc()

	n = RDP_socket_close(sock);
	if n < 0 {
		log.Println("RDP_socket_close failed", n)
	}
	RDP_cleanup();
	log.Println("rdp server stop")
}


