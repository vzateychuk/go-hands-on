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

type Executor struct {
	input      chan *WorkItem
	wrksAmount int
}

func (w *Executor) DoWork(wrk *WorkItem) error {
	// записываем задания в очередь задач
	go func() {
		w.input <- wrk
	}()
	return nil
}

var executor *Executor
var once sync.Once

func GetExecutorInstance(parallelism int) *Executor {

	// Создаем singleton
	once.Do(func() {
		fmt.Println("Singleton : Creating Executor instance")

		// канал которым будут пользоваться executor-ы для передачи заданий
		inputChan := make(chan *WorkItem, 2)

		// создаем экземпляр executor, со ссылкой на input
		// channel через который будем передавать задания
		executor = &Executor{
			input:      inputChan,
			wrksAmount: parallelism,
		}

		// стартует pool executor-в зачитывающих значения из канала
		for i := 0; i < parallelism; i++ {
			go startWorker(i, inputChan)
		}

	})
	return executor
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
	fmt.Printf("<=== Worker: %v Stoped\n", workerId) // сообщает о завершении executor
}

func doWork(workerId int, wrk *WorkItem) string {
	log.Printf("Worker#[%d] Started task [%s]\n", workerId, wrk.WorkId)
	waitTime := time.Duration(rand.Intn(3000)+10) * time.Millisecond // эмулируем случайную задержку
	time.Sleep(waitTime)
	log.Printf("Worker#[%d] Finished task [%s] in [%d]\n", workerId, wrk.WorkId, int(waitTime.Milliseconds()))
	return fmt.Sprintf("Worker#[%d] completed task: [%s]\n", workerId, wrk.WorkId)
}

//endregion
