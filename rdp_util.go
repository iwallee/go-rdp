package rdp

import (
	"net"
)

//copy from ip.go


// Parse IPv4 address (d.d.d.d).
func parseIPv4(s string) net.IP {
	var p [net.IPv4len]byte
	i := 0
	for j := 0; j < net.IPv4len; j++ {
		if i >= len(s) {
			// Missing octets.
			return nil
		}
		if j > 0 {
			if s[i] != '.' {
				return nil
			}
			i++
		}
		var (
			n  int
			ok bool
		)
		n, i, ok = dtoi(s, i)
		if !ok || n > 0xFF {
			return nil
		}
		p[j] = byte(n)
	}
	if i != len(s) {
		return nil
	}
	return net.IPv4(p[0], p[1], p[2], p[3])
}

// parseIPv6 parses s as a literal IPv6 address described in RFC 4291
// and RFC 5952.  It can also parse a literal scoped IPv6 address with
// zone identifier which is described in RFC 4007 when zoneAllowed is
// true.
func parseIPv6(s string, zoneAllowed bool) (ip net.IP, zone string) {
	ip = make(net.IP, net.IPv6len)
	ellipsis := -1 // position of ellipsis in p
	i := 0         // index in string s

	if zoneAllowed {
		s, zone = splitHostZone(s)
	}

	// Might have leading ellipsis
	if len(s) >= 2 && s[0] == ':' && s[1] == ':' {
		ellipsis = 0
		i = 2
		// Might be only ellipsis
		if i == len(s) {
			return ip, zone
		}
	}

	// Loop, parsing hex numbers followed by colon.
	j := 0
	for j < net.IPv6len {
		// Hex number.
		n, i1, ok := xtoi(s, i)
		if !ok || n > 0xFFFF {
			return nil, zone
		}

		// If followed by dot, might be in trailing IPv4.
		if i1 < len(s) && s[i1] == '.' {
			if ellipsis < 0 && j != net.IPv6len-net.IPv4len {
				// Not the right place.
				return nil, zone
			}
			if j+net.IPv4len > net.IPv6len {
				// Not enough room.
				return nil, zone
			}
			ip4 := parseIPv4(s[i:])
			if ip4 == nil {
				return nil, zone
			}
			ip[j] = ip4[12]
			ip[j+1] = ip4[13]
			ip[j+2] = ip4[14]
			ip[j+3] = ip4[15]
			i = len(s)
			j += net.IPv4len
			break
		}

		// Save this 16-bit chunk.
		ip[j] = byte(n >> 8)
		ip[j+1] = byte(n)
		j += 2

		// Stop at end of string.
		i = i1
		if i == len(s) {
			break
		}

		// Otherwise must be followed by colon and more.
		if s[i] != ':' || i+1 == len(s) {
			return nil, zone
		}
		i++

		// Look for ellipsis.
		if s[i] == ':' {
			if ellipsis >= 0 { // already have one
				return nil, zone
			}
			ellipsis = j
			if i++; i == len(s) { // can be at end
				break
			}
		}
	}

	// Must have used entire string.
	if i != len(s) {
		return nil, zone
	}

	// If didn't parse enough, expand ellipsis.
	if j < net.IPv6len {
		if ellipsis < 0 {
			return nil, zone
		}
		n := net.IPv6len - j
		for k := j - 1; k >= ellipsis; k-- {
			ip[k+n] = ip[k]
		}
		for k := ellipsis + n - 1; k >= ellipsis; k-- {
			ip[k] = 0
		}
	} else if ellipsis >= 0 {
		// Ellipsis must represent at least one 0 group.
		return nil, zone
	}
	return ip, zone
}

// copy from ipsock.go


func splitHostZone(s string) (host, zone string) {
	// The IPv6 scoped addressing zone identifier starts after the
	// last percent sign.
	if i := last(s, '%'); i > 0 {
		host, zone = s[:i], s[i+1:]
	} else {
		host = s
	}
	return
}


// copy from parse.go

// Bigger than we need, not too big to worry about overflow
const big = 0xFFFFFF
// Decimal to integer starting at &s[i0].
// Returns number, new offset, success.
func dtoi(s string, i0 int) (n int, i int, ok bool) {
	n = 0
	for i = i0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		n = n*10 + int(s[i]-'0')
		if n >= big {
			return 0, i, false
		}
	}
	if i == i0 {
		return 0, i, false
	}
	return n, i, true
}

// Hexadecimal to integer starting at &s[i0].
// Returns number, new offset, success.
func xtoi(s string, i0 int) (n int, i int, ok bool) {
	n = 0
	for i = i0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			n *= 16
			n += int(s[i] - '0')
		} else if 'a' <= s[i] && s[i] <= 'f' {
			n *= 16
			n += int(s[i]-'a') + 10
		} else if 'A' <= s[i] && s[i] <= 'F' {
			n *= 16
			n += int(s[i]-'A') + 10
		} else {
			break
		}
		if n >= big {
			return 0, i, false
		}
	}
	if i == i0 {
		return 0, i, false
	}
	return n, i, true
}

// Index of rightmost occurrence of b in s.
func last(s string, b byte) int {
	i := len(s)
	for i--; i >= 0; i-- {
		if s[i] == b {
			break
		}
	}
	return i
}

