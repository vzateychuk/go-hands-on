package main

import "fmt"

func main() {
	finished := make(chan bool)
	names := []string{"Packt"}

	go func() {
		names = append(names, "Electric")
		names = append(names, "Boogaloo")
		finished <- true
	}()
	
	// чтение данных из массива происходит параллельно изменению массива
	// решением проблемы было бы сделать чтение ПОСЛЕ завершения обновления, 
	// т.е. после <-finished
	for _, name := range names {
		fmt.Println(name)
	}
	<-finished	// так будет race-condition
}
