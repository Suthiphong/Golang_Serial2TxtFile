package main

import (
	"fmt"
)

func main() {
	var firstname string
	var surname string

	fmt.Print("COM PORT  : ")
	fmt.Scanln(&firstname)

	fmt.Print("BaudRate : ")
	fmt.Scanln(&surname)

	fmt.Print(firstname)
	fmt.Print(surname)
}

