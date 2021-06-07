package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	// мапа хранит значения которые мы уже видели
	alreadySeen := make(map[string]bool)

	for in.Scan() {
		txt := in.Text()

		// если мы такое значение уже видели, просто пропускаем строку
		if _, found := alreadySeen[txt]; found {
			continue
		}

		alreadySeen[txt] = true
		
		fmt.Println(txt)
	}
}
