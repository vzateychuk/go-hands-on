package main

import (
	"database/sql"
	"fmt"
	"log"
)
import _ "github.com/lib/pq"

func main() {

	db, err := sql.Open("postgres",
		"user=postgres password=root host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO public.test VALUES ($1, $2)")
	if err != nil {
		log.Panic(err)
	}
	defer insert.Close()

	prop := ""
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			prop = "Even"
		} else {
			prop = "Odd"
		}
		_, err = insert.Exec(i, prop)
		if err != nil {
			log.Panic(err)
		} else {
			log.Println("The number:", i, "is:", prop)
		}
	}
	fmt.Println("The numbers are ready.")
}
