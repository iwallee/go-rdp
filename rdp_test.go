package udt

import (
	"testing"
	"log"
	"net"
	"syscall"
)

func TestRun(t *testing.T) {
	log.Println("udt server start")
	n, _ := UDT_startup()
	log.Println(n)

	s, _ := UDT_socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	log.Println(n)

	addr := UDTAddr {
		IP: net.IPv4(0, 0, 0, 0),
		Port: 8389,
	}

	_, _ = UDT_bind(s, &addr)
	log.Println(n)

	_, _ = UDT_listen(s, 10)
	log.Println(n)

	for {
		remote := UDTAddr {
			IP: net.IPv4(0, 0, 0, 0),
			Port: 0,
		};
		r, _ := UDT_accept(s, &remote)
		if r == UDTSOCKET(INVALID_SOCK) {
			break
		}
	}
	log.Println("udt server stop")
}


