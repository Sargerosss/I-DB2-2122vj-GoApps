package main

// Standaard package die nodig is
import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.EvenParity,
		DataBits: 7,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open("COM3", mode)
	if err != nil {
		log.Fatal(err)
	}

	n, err := port.Write([]byte("1\n\r"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	buff := make([]byte, 100)

	n1, err := port.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	if n1 != 0 {
		fmt.Println("\nEOF")

	}
	fmt.Printf("%v", string(buff[:n1]))
}
