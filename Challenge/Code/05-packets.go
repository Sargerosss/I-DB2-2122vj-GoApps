package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func packetText() {
	fmt.Println("The packet tool")
	fmt.Println("It has the following features")
	fmt.Println("1. Find All Devices")
	fmt.Println("2. Capture Packets")
	// Only bug: No eth0 on Windows?

}

func packetOption() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")
	scanner.Scan()
	optionString := scanner.Text()
	option, err := strconv.Atoi(optionString)
	checkError(err)
	return option
}

func packetTool(user User, db *sql.DB) {
	packetText()
	option := packetOption()
	if option == 1 {
		findDevices()
		time.Sleep(2 * time.Second)
		defer continueTool(user, db)
	} else if option == 2 {
		capturePacket()
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else {
		falseOptionFunc(user, db)
	}

}

func findDevices() {
	fmt.Println("Find All Devices")

	devices, err := pcap.FindAllDevs()
	checkError(err)

	// Print device information
	fmt.Println("Devices found:")
	for _, device := range devices {
		time.Sleep(2 * time.Second)
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addresses: ", device.Description)
		for _, address := range device.Addresses {
			time.Sleep(2 * time.Second)
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
	time.Sleep(2 * time.Second)
}

func capturePacket() {
	var (
		snapshot_len int32 = 1024
		promiscuous  bool  = false
		err          error
		timeout      time.Duration = 30 * time.Second
		handle       *pcap.Handle
		device       string
	)

	fmt.Println("Please choose a device to capture")
	fmt.Println("Example: eth0")
	fmt.Scanln(&device)

	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	checkError(err)
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
	}
}
