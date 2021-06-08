package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
* Функция считывае из входного потока отсортированные строки, и сохраняет в выходной поток уникальный поток входящих строк.
* Если строки не отсортированы, генерируется ошибка.
 */
func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	var prev string

	for in.Scan() {
		txt := in.Text()

		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("data not sorted")
		}

		prev = txt

		// позволяет писать в поток output
		fmt.Fprintln(output, txt)
	}
	// если все выполнилось, возвращаем пустую ошибку
	return nil
}

func main() {

	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
