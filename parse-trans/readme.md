Parsing Bank Transaction Files
---

Ingesting a transaction file from the bank. The file is a CSV file. 
Our bank also includes budget categories for the transactions in the file.

The aim of this activity is to create a command-line program that will accept two flags:
the location of the CSV bank transaction file and the location of a log file. We will check
that the log and bank file location is valid before the application starts parsing the CSV
file. The program will parse the CSV file and log any errors it encounters to the log.
Upon each restart of the program, it will also delete the previous log file.

Запуск из командной строки
---
`go run main.go -c bank.csv -l log.log`