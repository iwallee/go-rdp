package rdp

import (
	"testing"
	"log"
	"net"
	"syscall"
)

type TestServer struct {

}
func  (s*TestServer)On_connect(param *RDP_on_connect_param){

}
func  (s*TestServer)On_before_accept(param *RDP_on_before_accept_param) bool{
	return true
}
func  (s*TestServer)On_accept(param *RDP_on_accept_param){

}
func  (s*TestServer)On_disconnect(param *RDP_on_disconnect_param){

}
func  (s*TestServer)On_recv(param *RDP_on_recv_param){

}
func  (s*TestServer)On_send(param *RDP_on_send_param){

}
func  (s*TestServer)On_udp_recv(param *RDP_on_udp_recv_param){

}
func  (s*TestServer)On_hash_addr(addr *RDPAddr) uint32{
	return 0
}

func TestRun(t *testing.T) {
	log.Println("rdp server start")
	var server TestServer
	var startup_param RDP_startup_param

	startup_param.Version   = RDP_SDK_VERSION
	startup_param.Max_sock = 1
	startup_param.Recv_thread_num = 1
	startup_param.Recv_buf_size = 1024*4
	startup_param.On_connect = server
	startup_param.On_before_accept = server
	startup_param.On_accept = server
	startup_param.On_disconnect = server
	startup_param.On_recv = server
	startup_param.On_send = server
	startup_param.On_udp_recv = server
	//startup_param.On_hash_addr = server


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

	n = RDP_socket_bind(s, &addr)
	if n < 0{
		log.Println("RDP_socket_bind failed", n)
		return
	}

	n = RDP_listen(s)
	if n < 0{
		log.Println("RDP_listen failed", n)
		return
	}


	log.Println("rdp server stop")
}


