package caller

import (
	"bytes"
	"net"
	"net/http"
	"strings"
)

type ipRange struct {
	start net.IP
	end   net.IP
}

func inRange(r ipRange, ipAddress net.IP) bool {
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return true
	}
	return false
}

var privateRanges = []ipRange{
	ipRange{ // This host on this network - RFC 1122
		start: net.ParseIP("0.0.0.0"),
		end:   net.ParseIP("0.255.255.255"),
	},
	ipRange{ // Private Networks - RFC 1918
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	ipRange{ // Carrier grade NAT - RFC 6598
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	ipRange{ // Loopback - RFC 1122
		start: net.ParseIP("127.0.0.0"),
		end:   net.ParseIP("127.255.255.255"),
	},
	ipRange{ // Link Local - RFC 3927
		start: net.ParseIP("169.254.0.0"),
		end:   net.ParseIP("169.254.255.255"),
	},
	ipRange{ // Private Networks - RFC 1918
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	ipRange{ // IANA IPv4 Special Purpose Address Registry - RFC 5736
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	ipRange{ // TEST-NET-1 - RFC 5737
		start: net.ParseIP("192.0.2.0"),
		end:   net.ParseIP("192.0.2.255"),
	},
	ipRange{ // 6to4 Relay Anycast - RFC 3068
		start: net.ParseIP("192.88.99.0"),
		end:   net.ParseIP("192.88.99.255"),
	},
	ipRange{ // Private Networks - RFC 1918
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	ipRange{ // Inter-network testing - RFC 2544
		start: net.ParseIP("198.18.0.0"),
		end:   net.ParseIP("198.19.255.255"),
	},
	ipRange{ // TEST-NET-2 - RFC 5737
		start: net.ParseIP("198.51.100.0"),
		end:   net.ParseIP("198.51.100.255"),
	},
	ipRange{ // TEST-NET-3 - RFC 5737
		start: net.ParseIP("203.0.113.0"),
		end:   net.ParseIP("203.0.113.255"),
	},
	ipRange{ // Future use - RFC 1122
		start: net.ParseIP("240.0.0.0"),
		end:   net.ParseIP("240.255.255.255"),
	},
	ipRange{ // Limited Broadcast - RFC 0919
		start: net.ParseIP("255.255.255.255"),
		end:   net.ParseIP("255.255.255.255"),
	},
}

func isPrivateSubnet(ipAddress net.IP) bool {
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		for _, r := range privateRanges {
			if inRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}

func IP(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addresses := strings.Split(r.Header.Get(h), ",")
		for i := len(addresses) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addresses[i])
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() || isPrivateSubnet(realIP) {
				continue
			}
			return ip
		}
	}
	return "127.0.0.1"
}
