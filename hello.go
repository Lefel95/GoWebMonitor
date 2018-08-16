package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntro()
	showMenu()
	command := readCommand()

	switch command {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Showing Logs")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Unknown Command")
		os.Exit(-1)
	}
}

func showIntro() {
	name := "Felipe"
	version := 1.1

	fmt.Println("Hello mr.", name)
	fmt.Println("Software Version:", version)
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	fmt.Println("The Chosen Command is", command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring")
	site := "http://google.com"
	fmt.Println(http.Get(site))
}

// func exitProgram(bool error) {
// 	var error = error
// 	if error == true {

// 	} else {

// 	}
// }
