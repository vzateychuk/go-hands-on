package main

import "fmt"

func DeleteElement(values []string) []string {
	/*	// Альтернативное решение:
		// Важно объявить slice для 2-х значений, куда будут скопированы 2 из values. Если result := []string{},
		// то будет создан slice len=0,cap=0 и тогда команда 'copy' не скопирует ничего потомучто некуда.
		result := make([]string, 2, 4)	// slice len=2, cap=4
		var copied int = copy(result, values[0:2])	// будет скопировано 2 элемента
		result = append(result, values[3:]...) // добавлены 2 элемента в конец slice result
	*/
	return append(values[0:2], values[3:]...)
}

/* Нужно удалить из ["1", "2", "Bad", "4", "5"] значение "Bad" */
func main() {
	values := []string{"1", "2", "Bad", "4", "5"}
	result := DeleteElement(values)
	fmt.Println(result)
}
