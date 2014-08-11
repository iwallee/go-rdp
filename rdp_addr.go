package rdp

import (
	"net"
)

type RDPAddr struct {
	IP   net.IP
	Port int
	Zone string // IPv6 scoped addressing zone
}

// Network returns the address's network name, "udt".
func (a *RDPAddr) Network() string { return "rdp" }

func (a *RDPAddr) String() string {
	if a == nil {
		return "<nil>"
	}
	ip := ipEmptyString(a.IP)
	if a.Zone != "" {
		return net.JoinHostPort(ip+"%"+a.Zone, itoa(a.Port))
	}
	return net.JoinHostPort(ip, itoa(a.Port))
}

func (a *RDPAddr) toAddr() net.Addr {
	if a == nil {
		return nil
	}
	return a
}

func ipEmptyString(ip net.IP) string {
	if len(ip) == 0 {
		return ""
	}
	return ip.String()
}

// Integer to decimal.
func itoa(i int) string {
	var buf [30]byte
	n := len(buf)
	neg := false
	if i < 0 {
		i = -i
		neg = true
	}
	ui := uint(i)
	for ui > 0 || n == len(buf) {
		n--
		buf[n] = byte('0' + ui % 10)
		ui /= 10
	}
	if neg {
		n--
		buf[n] = '-'
	}
	return string(buf[n:])
}

