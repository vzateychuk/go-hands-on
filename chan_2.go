package main

import "fmt"

func main() {

	in := make(chan int)

	// стрелка показыват что func в только писать
	go func(out chan<- int) {
		for i := 0; i <= 4; i++ {
			fmt.Println("Generate before: ", i)
			out <- i
			fmt.Println("Generate after: ", i)
		}
		// закрытие канала помечает range канала как завершенный
		close(out)
		fmt.Println("Generator finish")
	}(in)
	// итерируемся по значениям канала пока тот не закрыт
	for i := range in {
		fmt.Println("\tMain get: ", i)
	}
}
