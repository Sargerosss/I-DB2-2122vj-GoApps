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
	time.Sleep(2 * time.Second)
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	fmt.Println("Baudrate check")
	time.Sleep(2 * time.Second)

	port, err := serial.Open(ports[0], mode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Almost sending")
	time.Sleep(2 * time.Second)

	script1(port, arg)

	defer port.Close()
}

func script1(port serial.Port, arg string) {
	if arg == "golang" {
		n, err := port.Write([]byte("Golang"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", n)
	} else if arg == "arduino" {
		n2, err := port.Write([]byte("Arduino"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", n2)
	} else if arg == "hello" {
		n3, err := port.Write([]byte("Hello"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", n3)
	}
}
