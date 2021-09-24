package work

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Doable interface {
	DoWork(wrk *WorkItem)
}

type Worker struct {
	input      chan *WorkItem
	wrksAmount int
}

func (w *Worker) DoWork(wrk *WorkItem) error {
	go func() {
		w.input <- wrk
	}()
	return nil
}

var worker *Worker
var once sync.Once

func GetWorkerInstance(amount int) *Worker {

	// Use the sync.Once to enforce goroutine safety create singleton
	once.Do(func() {
		fmt.Println("Singleton : Creating Worker instance")

		// канал которым будут пользоваться worker-ы для передачи заданий
		inputChan := make(chan *WorkItem, 2)

		// создаем экземпляр worker, со ссылкой на channel через который будем передавать задания
		worker = &Worker{
			input:      inputChan,
			wrksAmount: amount,
		}

		// стартует pool worker-в зачитывающих значения из канала
		for i := 0; i < amount; i++ {
			go startWorker(i, inputChan)
		}

	})
	return worker
}

//region Private

func startWorker(workerId int, inputChan chan *WorkItem) {
	// сообщает о запуске worker
	fmt.Printf("===> Worker: %d Started\n", workerId)
	// пока в канале есть значения вычитываем и выполняем задания
	for input := range inputChan {
		doWork(workerId, input)
		// передаем выполнение следующей goroutine
		runtime.Gosched()
	}
	fmt.Printf("<=== Worker: %v Stoped\n", workerId) // сообщает о завершении worker
}

func doWork(workerId int, wrk *WorkItem) string {
	log.Printf("Worker#[%d] Started task [%s]\n", workerId, wrk.WorkId)
	waitTime := time.Duration(rand.Intn(3000)+10) * time.Millisecond // эмулируем случайную задержку
	time.Sleep(waitTime)
	log.Printf("Worker#[%d] Finished task [%s] in [%d]\n", workerId, wrk.WorkId, int(waitTime.Milliseconds()))
	return fmt.Sprintf("Worker#[%d] completed task: [%s]\n", workerId, wrk.WorkId)
}

//endregion
