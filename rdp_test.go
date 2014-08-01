package rdp

import (
	"testing"
	"log"
	"net"
	"syscall"
)

func TestRun(t *testing.T) {
	log.Println("udt server start")
	n, _ := RDP_startup()
	log.Println(n)

	s, _ := RDP_socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	log.Println(n)

	addr := RDPAddr {
		IP: net.IPv4(0, 0, 0, 0),
		Port: 8389,
	}

	_, _ = RDP_bind(s, &addr)
	log.Println(n)

	_, _ = RDP_listen(s, 10)
	log.Println(n)

	for {
		remote := RDPAddr {
			IP: net.IPv4(0, 0, 0, 0),
			Port: 0,
		};
		r, _ := RDP_accept(s, &remote)
		if r == RDPSOCKET(INVALID_SOCK) {
			break
		}
	}
	log.Println("udt server stop")
}


