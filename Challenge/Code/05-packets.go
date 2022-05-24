package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func packetText() {
	fmt.Println("The packet tool")
	fmt.Println("It has the following features")
	fmt.Println("1. Find All Devices")
	fmt.Println("2. Capture Packets")

}

func packetOption() int {
	var option int
	fmt.Scanln(&option)
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
		fmt.Println("Invalid option or invalid permission level. Please try again,", user.Username)
		time.Sleep(2 * time.Second)

		if falseOption == 2 {
			fmt.Println("Closing application, too many wrong inputs")
		} else {
			falseOption++
			fmt.Println("Please be careful, you have", falseOption, "invalid options given/tried.")
			time.Sleep(10 * time.Second)
			selectTool(user, db)
		}
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
