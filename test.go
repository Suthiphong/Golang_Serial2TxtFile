package main

import "log"

func main(){
	log.Printf("hello")
	defer log.Printf("end of ")
}