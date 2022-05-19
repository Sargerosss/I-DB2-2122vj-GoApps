package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ghedo/go.pkt/capture/pcap"
	"github.com/ghedo/go.pkt/filter"
	"github.com/ghedo/go.pkt/layers"
	"github.com/ghedo/go.pkt/packet"
	"github.com/ghedo/go.pkt/packet/arp"
	"github.com/ghedo/go.pkt/packet/eth"
)

// https://github.com/ghedo/go.pkt/tree/master/examples
// https://github.com/ghedo/go.pkt

func packetTool(user User, db *sql.DB) {
	option := packetText()
	packetOption(option, user, db)
}

func packetText() int {
	fmt.Println("This is the packet tool, similar to Wireshark, it has a lot of features around packets")
	fmt.Println("These are the options:")
	fmt.Println("1. Capturing")
	fmt.Println("2. Injection")
	fmt.Println("3. Filtering")
	fmt.Println("4. Encoding")
	fmt.Println("5. Decoding")
	var option int
	fmt.Scanln(&option)
	return option
}

func packetOption(option int, user User, db *sql.DB) {
	if option == 1 {
		src, err := pcap.Open("eth0")
		checkError(err)
		defer src.Close()

		err = src.Activate()
		checkError(err)

		for {
			buf, err := src.Capture()
			checkError(err)
			fmt.Println(buf)
			time.Sleep(1 * time.Second)
			continueTool(user, db)
		}

	} else if option == 2 {
		dst, err := pcap.Open("eth0")
		checkError(err)
		defer dst.Close()

		// you may configure the source further, e.g. by activating
		// promiscuous mode.

		err = dst.Activate()
		checkError(err)
		time.Sleep(2 * time.Second)
		var data string
		fmt.Println("Put something to inject")
		fmt.Scanln(&data)
		err = dst.Inject([]byte(data))
		checkError(err)
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 3 {
		fmt.Println("Choose between UDP or TCP")
		var choice string
		fmt.Scanln(&choice)
		flt, err := filter.Compile(choice, packet.Eth, true)
		checkError(err)
		fmt.Println("Choose some data to try and match")
		var data string
		fmt.Scanln(&data)
		if flt.Match([]byte(data)) {
			log.Println("MATCH!!!")
		}
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 4 {
		fmt.Println("Let's make some packets")
		fmt.Println("Starting with Eth packet")
		fmt.Println("Give a Source Address (IPV6)")
		var srcAddr string
		var dstAddr string
		fmt.Scanln(&srcAddr)
		fmt.Println("Give a destination Address (IPV6)")
		fmt.Scanln(&dstAddr)

		eth_pkt := eth.Make()
		eth_pkt.SrcAddr, _ = net.ParseMAC(srcAddr)
		eth_pkt.DstAddr, _ = net.ParseMAC(dstAddr)

		time.Sleep(2 * time.Second)
		fmt.Println("Now an ARP Packet")
		var ARPSrcAddr string
		var ARPDstAddr string
		var ARPprotoSrcAddr string
		var ARPprotoDstAddr string
		fmt.Println("IPV6 Source Address")
		fmt.Scanln(&ARPSrcAddr)
		fmt.Println("IPV6 Destination Address")
		fmt.Scanln(&ARPDstAddr)
		fmt.Println("IPV4 Source Address")
		fmt.Scanln(&ARPprotoSrcAddr)
		fmt.Println("IPV4 Source Address")
		fmt.Scanln(&ARPprotoDstAddr)
		arp_pkt := arp.Make()
		arp_pkt.HWSrcAddr, _ = net.ParseMAC(ARPSrcAddr)
		arp_pkt.HWDstAddr, _ = net.ParseMAC(ARPDstAddr)
		arp_pkt.ProtoSrcAddr = net.ParseIP(ARPprotoSrcAddr)
		arp_pkt.ProtoDstAddr = net.ParseIP(ARPprotoDstAddr)

		buf, err := layers.Pack(eth_pkt, arp_pkt)
		checkError(err)
		fmt.Println(buf)
		time.Sleep(5 * time.Second)
		continueTool(user, db)
	} else if option == 5 {
		fmt.Println("Give some data to decode (Ethernet layer is used)")
		var data string
		fmt.Scanln(&data)
		buf := []byte(data)

		// Assume Ethernet as datalink layer
		pkt, err := layers.UnpackAll(buf, packet.Eth)
		checkError(err)

		fmt.Println(pkt)
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	}

}
