package main

import (
	"fmt"
)

func main() {
	name := "Hugo"
	version := 1.1
	fmt.Println("Hello, Mr.", name)
	fmt.Println("This pragram is in version", version)

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - View logs")
	fmt.Println("0 - Exit")

	var command int
	fmt.Scan(&command)
	fmt.Println("The command chosen was", command)

	if command == 1 {
		fmt.Println("Monitoring...")
	} else if command == 2 {
		fmt.Println("Viewing logs...")
	} else if command == 0 {
		fmt.Println("Exiting...")
	} else {
		fmt.Println("Invalid code")
	}
}
