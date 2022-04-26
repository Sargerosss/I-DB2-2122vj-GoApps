package main

// GENERAL CONCEPT
// When using this script, it will send a message to the Arduino connected to the Serial Port (COM(X)).
// Then the Arduino takes it over and executes his own code

// Imports to use
import (
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"
)

// Check error
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	serialCommunicationArduino()
}

func serialCommunicationArduino() {
	// Check ports that are connected/used
	ports, err := serial.GetPortsList()
	checkError(err)

	// If no ports found, log error
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	// Print found ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	// Wait 1 sec
	time.Sleep(1 * time.Second)

	// Settings to use, 9600 baudrate (IMPORTANT)
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	time.Sleep(1 * time.Second)

	// Open port, with above settings
	port, err := serial.Open(ports[0], mode)

	checkError(err)

	// Wait 2 seconds
	time.Sleep(2 * time.Second)

	// Use the port to start script2
	script2(port)

	// Wait to close the port
	defer port.Close()
}

// Takes in port as argument
func script2(port serial.Port) {
	// Write 1 to Arduino
	n, err := port.Write([]byte("100;1"))
	checkError(err)
	// Print bytes sent
	fmt.Printf("Sent %v bytes\n", n)

}
