package util

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

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

// GetLocalIP 获取本机IP
func GetLocalIP() string {
	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}
	return ip.String()
}

// GetLocalShortIP 获取本机IP的最后一段地址
// 例如192.168.1.104  返回 104
func GetLocalShortIP() string {
	shorts := strings.Split(GetLocalIP(), ".")
	return shorts[len(shorts)-1]
}

// GetLocalIntShortIP 获取本机IP的最后一段地址的int64类型，用户生成雪花ID的节点值
func GetLocalIntShortIP() int64 {
	ipShortInt, _ := strconv.Atoi(GetLocalIP())
	return int64(ipShortInt)
}
