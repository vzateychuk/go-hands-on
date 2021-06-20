package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	name := "Thanos"
	log.Println("Demo app")
	log.Printf("%s is here!", name)
	log.Print("Run")
}
