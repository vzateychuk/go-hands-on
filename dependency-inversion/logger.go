package main

// MyLogger - интерфейс функции журналирования.
// Предназначен чтобы бизнес-логика заносила в журнал сведения о времени вызова.
type MyLogger interface {
	MyLog(msg string)
}

// LoggerAdapter - адаптер чтобы сделать функцию LogOutput (пакет loggin).
// соответствующей интерфейсу MyLogger определим функциональный тип с требуемым методом
type LoggerAdapter func(msg string)

func (l LoggerAdapter) MyLog(msg string) {
	l(msg)
}
