package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	ErrWorkingFileNotFound = errors.New("The file was not found.")
)

func createBackup(fileName, backup string) error {

	_, err := os.Stat(fileName) // Проверяем что файл создан и доступен
	if err != nil {
		if os.IsNotExist(err) {
			return ErrWorkingFileNotFound
		} else {
			return err
		}
	}

	file, err := os.Open(fileName) // открыть файл
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(file) // прочитали файл
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(backup, content, 0644) // перезаписали backup файл
	if err != nil {
		return err
	}

	return nil
}

func addNotes(fileName, notes string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close() // Обязательно закрываем файл после завершения логики

	notes = "\n" + notes

	_, err = file.Write([]byte(notes)) // добавляем запись в файл
	return err
}

func main() {
	backupFile := "backupFile.txt"
	workingFile := "note.txt"
	data := "note"

	err := createBackup(workingFile, backupFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		note := data + strconv.Itoa(i)
		err := addNotes(backupFile, note)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("FINISH")
}
