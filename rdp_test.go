package rdp

import (
	"testing"
	"log"
	"net"
	"syscall"
)

func TestRun(t *testing.T) {
	log.Println("rdp server start")
	n, _ := RDP_startup()
	log.Println(n)

	s, _ := RDP_socket(syscall.AF_INET)
	log.Println(n)

	addr := RDPAddr {
		IP: net.IPv4(0, 0, 0, 0),
		Port: 8389,
	}

	_, _ = RDP_bind(s, syscall.AF_INET, &addr)
	log.Println(n)

	_, _ = RDP_listen(s, 10)
	log.Println(n)

	for {
		r, remote, _ := RDP_accept(s, syscall.AF_INET)
		if r < 0 {
			break
		}
		log.Println("accept", r, remote)
	}
	log.Println("rdp server stop")
}


