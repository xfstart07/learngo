// Author: Xu Fei
// Date: 2018/9/6
package main

import (
	"net"
	"strings"
	"fmt"
)

func main() {
	fmt.Println(getIP("172.20"))
}

// 获取指定ip段的ip
func getIP(ipSeg string) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	var ips []string
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	if len(ips) == 0 {
		return ""
	}

	if ipSeg == "" {
		return ips[0]
	}

	for index := range ips {
		if strings.HasPrefix(ips[index], ipSeg) {
			return ips[index]
		}
	}

	return ""
}