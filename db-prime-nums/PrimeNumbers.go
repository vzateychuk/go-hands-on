package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"math/big"
)

var number int64
var prop string
var primeSum int64
var newNumber int64

func main() {
	// connect to DB
	db, err := sql.Open("postgres",
		"user=postgres password=root host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println("The connection to the DB was successfully initialized!")
	}
	defer db.Close()

	// get all nums
	selectNums := "SELECT * FROM test"
	stmt, err := db.Prepare(selectNums)
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()

	allRows, err := stmt.Query()
	if err != nil {
		log.Panic(err)
	}
	log.Println("List of Prime numbers:")
	for allRows.Next() {
		err := allRows.Scan(&number, &prop)
		if err != nil {
			log.Panic(err)
		}
		// assert is prime
		if big.NewInt(number).ProbablyPrime(0) {
			primeSum += number
			log.Println("Prime: ", number)
		}
	}
	allRows.Close()
	log.Println("\nThe total sum of prime numbers in this range is:", primeSum)
}
