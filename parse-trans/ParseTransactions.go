package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"parse-trans/budget"
	"strconv"
	"strings"
)

var FileNotFoundError = errors.New("file not found.")

// Записывает msg, err, в log file.
func writeToLog(msg string, err error, logfile string) error {
	msg += "\n" + err.Error()
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(msg); err != nil {
		return err
	}
	return nil
}

// Читает запись из record и мапинг в Transaction
func parseRecordAndMapToTransaction(record []string, logFile string) budget.Transaction {

	value := strings.TrimSpace(record[0])
	id, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}

	payee := strings.TrimSpace(record[1])

	value = strings.TrimSpace(record[2])
	spent, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal(err)
	}

	value = strings.TrimSpace(record[3])
	category, err := budget.GetCategoryBy(value)
	if err != nil {
		s := strings.Join(record, ", ")
		writeToLog("Unable to find category: '"+value+"', record: '"+s+"'", err, logFile)
	}
	return budget.Transaction{
		ID:       id,
		Payee:    payee,
		Spent:    spent,
		Category: category,
	}
}

func parseBankFile(bankTransactions io.Reader, logFile string) []budget.Transaction {
	trxs := []budget.Transaction{}
	r := csv.NewReader(bankTransactions) // Создаем csv-reader для чтения
	header := true                       // пометка чтобы пропустить первую строку

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if !header {
			trx := parseRecordAndMapToTransaction(record, logFile)
			trxs = append(trxs, trx)
		}
		header = false
	}
	return trxs
}

func main() {
	// Получаем параметры командной строки - имя csv файла с данными и лог файла
	bankFile := flag.String("c", "", "location of the bank transaction csv file")
	logFile := flag.String("l", "", "logging of errors")
	flag.Parse()
	if *bankFile == "" || *logFile == "" {
		fmt.Println("csvFile and logFile is required.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// проверяем что csv файл существует. Если нет, выходим
	_, err := os.Stat(*bankFile)
	if os.IsNotExist((err)) {
		fmt.Println("BankFile does not exist: ", *bankFile)
		os.Exit(1)
	}
	// проверяем что лог файл не существует. Если существует, удалим
	_, err = os.Stat(*logFile)
	if !os.IsNotExist(err) {
		os.Remove(*logFile)
	}

	// Пробуем открыть на чтение файл bankFile c помощью os.Open, возвращает *os.File который имплемент io.Reader interface
	csvFile, err := os.Open(*bankFile)
	if err != nil {
		fmt.Println("Error opening file: ", *bankFile)
		os.Exit(1)
	}
	defer csvFile.Close()

	// Парсим файл получая массив транзакций
	trxs := parseBankFile(csvFile, *logFile)
	fmt.Println()
	for _, trx := range trxs {
		fmt.Printf("%v\n", trx)
	}

}
