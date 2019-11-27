package main

import "fmt"
import "log"
import "github.com/tarm/serial"

import "os"

func main() {
        config := &serial.Config{
                Name: "COM32",
                Baud: 9600,
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

/*
func Append(){
	file, err := os.OpenFile("./data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
			panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString("text to append\n"); err != nil {
			log.Println(err)
	}


}

*/