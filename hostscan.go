package gonmap

import (
	"kscan/lib/gosping"
	"time"
)

func HostDiscovery(ip string) bool {
	if HostDiscoveryForIcmp(ip) {
		return true
	}
	if HostDiscoveryForTcp(ip) {
		return true
	}
	return false
}

func HostDiscoveryForIcmp(ip string) bool {
	return gosping.OsPing(ip)
}

func HostDiscoveryForTcp(ip string) bool {
	tcpArr := []int{22, 23, 80, 139, 443, 445, 3389}
	for _, port := range tcpArr {
		if PortScan("tcp", ip, port, time.Second*2) {
			return true
		}
	}
	return false
}
