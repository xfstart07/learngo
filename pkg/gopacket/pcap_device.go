package main

import (
	"fmt"
	"strings"

	"github.com/google/gopacket/pcap"
)

// 获取设备信息
func main() {
	var devices []pcap.Interface
	devices, _ = pcap.FindAllDevs()

	for _, value := range devices {
		fmt.Printf("%+v \n", value)

		if strings.Contains(value.Name, "e") {
			fmt.Println("默认网卡名称 en0, 当前是:", value.Name)
		}
	}
}
