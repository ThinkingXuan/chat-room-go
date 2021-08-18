package util

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(GetIp())
}

func externalIP() (net.IP, error) {
	var ip net.IP
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		for _, addr := range addrs {

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			if len(ip.String()) > 0 {
				return ip, err
			}
			//fmt.Println("ip: ", ip.String(), "mac: ", iface.HardwareAddr.String())
		}
	}
	return ip, err
}

func GetIp() string {
	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(ip.String())
	return ip.String()
}
