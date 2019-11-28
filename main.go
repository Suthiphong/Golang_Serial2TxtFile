package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"os"
)

func main() {
	var sp string
	var baudrate int
	fmt.Println("########################################")
	fmt.Println("#      Serial to text                  #")
	fmt.Println("########################################")
	fmt.Print("COM PORT  : ")
	fmt.Scanln(&sp)
	fmt.Print("Baudrate : ")
	fmt.Scanln(&baudrate)
	fmt.Print(sp,  " ")
	fmt.Print(baudrate)
    config := &serial.Config{
                Name: sp,
                Baud: baudrate,
                ReadTimeout: 5,
                Size: 8,
		}
    stream, err := serial.OpenPort(config)
        if err != nil {
                log.Fatal(err)
        }

		buf := make([]byte, 1024)
		file, err := os.OpenFile("./data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Println("wait data")
		//if _, err := file.WriteString()
        for {
                n, err := stream.Read(buf)
                if err != nil {
                        log.Fatal(err)
                }
				s := string(buf[:n])
				
				fmt.Println(s)
				if _, err := file.WriteString(s); err != nil {
					log.Println(err)
				}
				
        }
        fmt.Scanln()

}

