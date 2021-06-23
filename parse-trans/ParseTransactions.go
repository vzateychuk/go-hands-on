package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"parse-trans/budget"
)

var FileNotFoundError = errors.New("file not found.")

// This will take the msg, err, and data strings and write them to the log file.
func writeToLog(logfile string, msg string) error {
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte(msg)); err != nil {
		return err
	}
	return nil
}

func readRecord(record []string) (id, payee, spent, category string) {
	id = record[0]
	payee = record[1]
	spent = record[2]
	category = record[3]
	return
}

func parseBankFile(bankTransactions io.Reader, logFile string) []budget.Transaction {
	result := []budget.Transaction{}
	r := csv.NewReader(bankTransactions) // creates a reader
	header := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if !header {
			id, payee, spent, scat := readRecord(record)
			category, err := budget.GetByName(scat)
			if err != nil {
				writeToLog(logFile, err.Error())
			}
			tran := budget.Transaction{
				ID:       id,
				Payee:    payee,
				Spent:    spent,
				Category: category,
			}
			result = append(result, tran)
		}
		header = false
	}
	return result
}

func main() {
	csvFileName := "trans.csv"
	logFileName := "parse.log"

	_, err := os.Stat(csvFileName)
	if err != nil {
		if os.IsNotExist(err) {
			writeToLog(logFileName, FileNotFoundError.Error())
			log.Panic(FileNotFoundError)
		}
	}

	csvFile, err := os.Open(csvFileName)
	if err != nil {
		writeToLog(logFileName, FileNotFoundError.Error())
		log.Panic(FileNotFoundError)
	}
	defer csvFile.Close()
	transactions := parseBankFile(csvFile, logFileName)
	for _, tran := range transactions {
		fmt.Println(tran.String())
	}
}
