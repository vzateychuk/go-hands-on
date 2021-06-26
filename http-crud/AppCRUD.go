package main

import (
	"flag"
	"fmt"
	"http-db/dbinit"
	"http-db/myhttp"
	"log"
)

/*
Создает таблицу в БД, заполняет данными
*/
func main() {
	dbInit := flag.Bool("init", false, "Run DBInit?")
	flag.Parse()
	if *dbInit {
		err := dbinit.DBInit("People")
		if err != nil {
			log.Panic(err)
		}
	}
	fmt.Println("Finish")

	myhttp.RunServer(8080)
}
