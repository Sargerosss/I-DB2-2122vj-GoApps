package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.bug.st/serial"
)

func getArg() string {
	args := os.Args
	arg := args[1]
	return arg
}

func main() {

	serialCommunication(getArg())
}

func serialCommunication(arg string) {
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
	time.Sleep(1 * time.Second)
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	fmt.Println("Baudrate check")
	time.Sleep(1 * time.Second)

	port, err := serial.Open(ports[0], mode)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)

	script1(port, arg)

	defer port.Close()
}

func script1(port serial.Port, arg string) {
	if arg == "Golang" {
		n, err := port.Write([]byte(arg))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", n)
	} else if arg == "Arduino" {
		n2, err := port.Write([]byte(arg))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", n2)
	} else if arg == "Hello" {
		n3, err := port.Write([]byte(arg))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", n3)
	} else if arg == "ABCDEF" {
		n3, err := port.Write([]byte(arg))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", n3)
	}
}
