package main

import (
	"flag"
	"http-db/db"
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
		err := db.DBInit("People")
		if err != nil {
			log.Panic(err)
		}
	}

	myhttp.RunServer(8080)
}
