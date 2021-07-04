package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

/*
Read numbers from a file and send them to a channel.
This function will run as two Goroutines, one per input file. Inside this function, we read from the file line by line.
	fileName = a filename for the input file to open;
	out = a channel to pipe messages in;
	wg = a WaitGroup to notify the end of the process.
*/
func readFromFile(fileName string, out chan<- int, wg *sync.WaitGroup) {

	// Open file for read
	file, err := os.Open(fileName)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	// creating a buffered reader on the file f and then looping the ReadString function 
	// with the newline character '\n' as the delimiter. Be mindful that it has to
	// be with single quotes and not "\n" because the delimiter is a character and not a string
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			// if EndOfFile, exit
			if err.Error() == "EOF" {
				wg.Done()
				return
			} else {
				log.Panicln(err)
			}
		}
		// We also need to strip the line so that we just have the number
		str = strings.ReplaceAll(str, "\n", "")
		i, err := strconv.Atoi(str)
		if err != nil {
			log.Panicln(err)
		}
		// finally put the int to output channel
		out <- i
	}
}

/* Receive the numbers and then send them to two different channels, one for odd numbers and one for even numbers
	in = channel to get numbers from the sources;
	odd, even = two channels to pipe numbers to, one for the even numbers and one for the odd numbers;
	wg = Waitgroup is used to notify the main routine of completion. 
The purpose of this function is to split the numbers so we can loop over the channel
*/
func splitEvenOdd(in <-chan int, odd, even chan<- int, wg *sync.WaitGroup) {

	for i:= range in {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	wg.Done()
}

// function to sum the numbers coming in and send the sum to an outbound channel
func sum(in <- chan int, out chan<- int, wg *sync.WaitGroup) {

	result := 0
	for i := range in {
		result += i
	}
	out <- result
	wg.Done()
}

// output the sum of the even and odd
func outputAll(even, odd <-chan int, wg *sync.WaitGroup, resultFile string) {
	rsfile, err := os.Create(resultFile)
	if err != nil {
		log.Panicln(err)
	}
	rsfile.WriteString(fmt.Sprintf("Even: %d\n", <-even))
	rsfile.WriteString(fmt.Sprintf("Odd: %d\n", <-odd))
	wg.Done()
}

/*
Please note that here, we could have used more than two files. You could have even
used an arbitrary number of files. Hence, there is no way for the splitter to know
how to terminate the execution, so we close the channel after the sources have
finished piping numbers in.
*/
func main() {
	// We used two Waitgroups here; one for the sources and one for the rest of the routines
	wg2 := &sync.WaitGroup{}
	wg2.Add(2)

	wg4 := &sync.WaitGroup{}
	wg4.Add(4)
	// odd and even are the ones where the numbers are piped for being summed
	odd := make(chan int)
	even := make(chan int)
	// out is the channel used by the source functions to pipe the messages to the splitter (splitEvenOdd)
	out := make(chan int)
	// holding a single number with the sum
	sumodd := make(chan int)
	sumeven := make(chan int)

	go readFromFile("./data/input1.dat", out, wg2)
	go readFromFile("./data/input2.dat", out, wg2)

	go splitEvenOdd(out, even, odd, wg4)
	go sum(even, sumeven, wg4)
	go sum(odd, sumodd, wg4)
	go outputAll(sumeven, sumodd, wg4, "./result.txt")

	wg2.Wait()
	close(out)
	wg4.Wait()
}