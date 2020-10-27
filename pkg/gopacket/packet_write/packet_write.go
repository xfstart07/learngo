package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var (
	device      string = "en0"
	snapshotLen int32  = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = 3 * time.Second
)

// 将抓取的包写入文件
func main() {
	handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)

	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	dumpFile, _ := os.Create("../dump.pcap")
	defer dumpFile.Close()
	packetWriter := pcapgo.NewWriter(dumpFile)
	packetWriter.WriteFileHeader(65535, layers.LinkTypeEthernet)

	// 从 packet source 中读取抓到的包
	packageSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packageSource.Packets() {
		fmt.Println(packet)

		packetWriter.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
	}
}
