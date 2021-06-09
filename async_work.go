package main

import (
	"fmt"
	"time"
)

// асинхронная загрузка комментариев
func loadCommentsAsync() chan string {
	// используем буферизированный канал
	result := make(chan string, 1)
	go func(out chan<- string) {
		time.Sleep(2 * time.Second) // эмуляция получения комментариев
		fmt.Println("loadCommentsAsync completed, return")
		out <- "32 comments"
	}(result)
	return result
}

// загрузка страницы вместе с комментариями
func getPage() string {
	commentsCh := loadCommentsAsync() // инициируем канал с асинх. загрузкой комментов
	time.Sleep(2 * time.Second)       // эмуляция загрузки страницы
	fmt.Println("get Page")
	comments := <-commentsCh
	fmt.Println("getPage goroutine, loaded:", comments)
	return "getPage() goroutine, with: " + comments
}

func main() {
	fmt.Println("Main start loading page")
	page := getPage()
	fmt.Println("Main loaded:", page)
}
