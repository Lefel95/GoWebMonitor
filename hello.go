package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntro()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing Logs")
		case 0:
			fmt.Println("Exiting...")
			error := false
			exitProgram(error)
		default:
			fmt.Println("Unknown Command")
			error := true
			exitProgram(error)
		}
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
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("The Website: ", site, " was loaded succesfully")
	} else {
		fmt.Println("The WebSite: ", site, "had a error! StatusCode: ", resp.StatusCode)
	}
}

func exitProgram(error bool) {
	if error == true {
		os.Exit(-1)
	} else {
		os.Exit(0)
	}
}
