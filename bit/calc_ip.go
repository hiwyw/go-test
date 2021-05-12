package main

import (
	"fmt"
	"net"
)

const (
	ipv4StartIp = "192.168.1.11"
	ipv4EndIp   = "192.168.1.12"

	ipv6StartIp = "2001:da8:1:1::123"
	ipv6EndIp   = "2001:da8:1:1::1234:8888"
)

func main() {
	ip1 := net.ParseIP(ipv4StartIp)
	ip2 := net.ParseIP(ipv4EndIp)
	fmt.Println(calcSubnet(ip1, ip2))

	ip3 := net.ParseIP(ipv6StartIp)
	ip4 := net.ParseIP(ipv6EndIp)
	fmt.Println(calcSubnet(ip3, ip4))
}

func calcSubnet(ip1, ip2 net.IP) *net.IPNet {
	var start, end []byte
	if isIpv4(ip1) {
		start = []byte(ip1)[12:]
		end = []byte(ip2)[12:]
	}
	if isIpv6(ip1) {
		start = []byte(ip1)
		end = []byte(ip2)
	}

	netbits := 0
	for i, sb := range start {
		eb := end[i]
		if uint8(sb) == uint8(eb) {
			netbits += 8
			continue
		}
		for j := 1; j <= 8; j++ {
			if uint8(sb)>>j == uint8(eb)>>j {
				netbits += (8 - j)
				break
			}
		}
		break
	}
	_, subnet, _ := net.ParseCIDR(fmt.Sprintf("%s/%d", ip1.String(), netbits))
	return subnet
}

func isIpv4(ip net.IP) bool {
	return ip != nil && ip.To4() != nil
}

func isIpv6(ip net.IP) bool {
	return ip != nil && ip.To4() == nil
}
