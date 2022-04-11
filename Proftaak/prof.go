package main

// Imports to use
import (
	"fmt"
	"log"
	"os"
	"time"

	"go.bug.st/serial"
)

// Check error
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Take argument
func getArg() string {
	args := os.Args
	arg := args[1]
	return arg
}

// Use arg from getArg and use it in serialCommunication
func main() {

	serialCommunication(getArg())
}

// Take arg to execute function
func serialCommunication(arg string) {
	// Check ports that are connected/used
	ports, err := serial.GetPortsList()
	checkErr(err)
	// If no ports found, error
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	// Print found ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
	// Wait 1 second
	time.Sleep(1 * time.Second)
	// Settings to use, 9600 baudrate (IMPORTANT)
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	// Small check
	fmt.Println("Baudrate check")
	// Wait 1 second
	time.Sleep(1 * time.Second)

	// Start communication, open the port with settings
	port, err := serial.Open(ports[0], mode)
	// Check error
	checkErr(err)
	// Wait 2 seconds
	time.Sleep(2 * time.Second)

	// Use the port and argument to start script1
	script1(port, arg)

	// Wait to close the port
	defer port.Close()
}

// Takes in port and arg as argument
func script1(port serial.Port, arg string) {
	// Check arg = Golang
	if arg == "Golang" {
		// Write Golang to the Arduino
		n, err := port.Write([]byte(arg))
		//Check error
		checkErr(err)
		// Print bytes sent
		fmt.Printf("Sent %v bytes\n", n)
		// Check arg = Arduino
	} else if arg == "Arduino" {
		// Write Arduino to the Arduino
		n2, err := port.Write([]byte(arg))
		// Error check
		checkErr(err)
		// Print bytes sent
		fmt.Printf("Sent %v bytes\n", n2)
		// Check arg = Hello
	} else if arg == "Hello" {
		// Write Hello to the Arduino
		n3, err := port.Write([]byte(arg))
		// Error check
		checkErr(err)
		// Print bytes sent
		fmt.Printf("Sent %v bytes\n", n3)
		// Check arg = ABCDEF
	} else if arg == "ABCDEF" {
		// Write ABCDEF to the Arduino
		n3, err := port.Write([]byte(arg))
		// Error check
		checkErr(err)
		// Print bytes sent
		fmt.Printf("Sent %v bytes\n", n3)
	}
}
