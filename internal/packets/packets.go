package packets

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	maxPacketLength int32 = 65535
	err             error
	handle          *pcap.Handle
)

func Run(
	iface string,
	packetChan chan gopacket.Packet,
	promiscuous bool,
) {
	handle, err = pcap.OpenLive(iface, maxPacketLength, promiscuous, pcap.BlockForever)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		packetChan <- packet
	}
}
