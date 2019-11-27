package main

import (
        "os"
        "log"
        //"io"
)

func main() {
        
        /*err := WriteToFile("./data.txt", "1234")
        if err != nil {
                log.Fatal(err)
        }*/
        Append()
        defer log.Printf("finish close file")

} 

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

/*
func WriteToFile(filename string, data string) error {
        file,err := os.Create(filename)
        if err != nil {
                panic(err)
        }
        defer file.Close()
        _, err = io.WriteString(file, data)
        if err != nil {
                panic(err)
        }
        return file.Sync()

}*/