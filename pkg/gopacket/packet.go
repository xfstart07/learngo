package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket/layers"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	device      string = "en0"
	snapshotLen int32  = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = 3 * time.Second
	offdump     string        = "dump.pcap"
)

func main() {
	// 在线显示
	handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)

	// 离线分析，读取文件
	// 使用 tcpdump -w dump.pcap 抓取包，并存储 dump.pcap 文件
	// handle, err := pcap.OpenOffline(offdump)

	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// 过滤
	// handle.SetBPFFilter("tcp")
	// 过滤例子
	// 过滤IP： 10.1.1.3
	// 过滤TCP：tcp
	// 过滤CIDR： 128.3/16
	// 过滤端口： port 53
	// 过滤TCP和端口: tcp and (port 80)
	// 过滤主机和端口： host 8.8.8.8 and udp port 53
	// 过滤网段和端口： net 199.16.156.0/22 and port
	// 过滤非本机 Web 流量： (port 80 and port 443) and not host 192.168.0.1

	// 从 packet source 中读取抓到的包
	packageSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packageSource.Packets() {
		fmt.Println(packet)

		// 列出包的层
		// for _, layer := range packet.Layers() {
		// 	fmt.Println(layer.LayerType())
		// }

		// go parseIP(packet)
		// go parseTCP(packet)
		go parseApplicationLayer(packet)
	}
}

func parseIP(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("解析IP层的包")
		ip, _ := ipLayer.(*layers.IPv4)
		// fmt.Printf("%+v", ip)
		fmt.Println(ip.SrcIP, ip.DstIP)
		fmt.Println(ip.Protocol)
	}
}

func parseTCP(packet gopacket.Packet) {
	iLayer := packet.Layer(layers.LayerTypeTCP)
	if iLayer != nil {
		fmt.Println("解析TCP层的包")
		info, _ := iLayer.(*layers.TCP)
		// fmt.Printf("%+v", ip)
		fmt.Println(info.SrcPort)
		fmt.Println(info.DstPort)
	}
}

func parseApplicationLayer(packet gopacket.Packet) {
	iLayer := packet.ApplicationLayer()
	if iLayer != nil {
		fmt.Println("解析应用层层的包")
		// fmt.Printf("%+v", ip)
		fmt.Println(iLayer.Payload())
	}
}
