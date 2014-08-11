package rdp

import (
	"testing"
	"log"
	"net"
)


func  On_connect(param *RDP_on_connect_param){

}
func  On_before_accept(param *RDP_on_before_accept_param) bool{
	return true
}
func On_accept(param *RDP_on_accept_param){

}
func On_disconnect(param *RDP_on_disconnect_param){

}
func On_recv(param *RDP_on_recv_param){

}
func On_send(param *RDP_on_send_param){

}
func On_udp_recv(param *RDP_on_udp_recv_param){

}
func On_hash_addr(addr *RDPAddr) uint32{
	return 0
}

func TestRun(t *testing.T) {
	return
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
	startup_param.On_udp_recv =On_udp_recv
	//startup_param.On_hash_addr = On_hash_addr


	n := RDP_startup(&startup_param)
	if n < 0{
		log.Println("RDP_startup failed", n)
		return
	}

	var socket_param RDP_socket_create_param
	socket_param.Is_v4   = true
	socket_param.Ack_timeout  = 100
	socket_param.Heart_beat_timeout  = 180
	socket_param.Max_send_queue_size = 0
	socket_param.Max_recv_queue_size  = 0
	socket_param.In_session_hash_size = 1024



	var sock RDPSOCKET
	sock, n = RDP_socket_create(&socket_param)
	if n < 0{
		log.Println("RDP_socket_create failed", n)
		return
	}

	addr := RDPAddr {
		IP: net.IPv4(0, 0, 0, 0),
		Port: 9000,
	}

	n = RDP_socket_bind(sock, &addr)
	if n < 0{
		log.Println("RDP_socket_bind failed", n)
		return
	}

	n = RDP_socket_listen(sock)
	if n < 0{
		log.Println("RDP_listen failed", n)
		return
	}


	log.Println("rdp server stop")
}


