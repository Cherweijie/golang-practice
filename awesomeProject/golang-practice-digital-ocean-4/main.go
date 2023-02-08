package main

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, ch chan<- int, wg *sync.WaitGroup) {
	// wg.Done() removed so that it won't reduce the WaitGroup count when it finishes.
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("sending %d to channel\n", idx)
		ch <- idx
	}
}

func printNumbers(idx int, ch <-chan int, wg *sync.WaitGroup) {
	// only reduce count when
	defer wg.Done()

	for num := range ch {
		fmt.Printf("%d: read %d from channel\n", idx, num)
	}
}

func main() {
	var wg sync.WaitGroup
	numberChan := make(chan int)

	for idx := 1; idx <= 3; idx++ {
		wg.Add(1)
		// 3 go routines for printing numbers
		go printNumbers(idx, numberChan, &wg)
	}

	generateNumbers(5, numberChan, &wg)

	close(numberChan)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}

// result:
/*
Since there are three printNumbers goroutines running, there’s an element of chance
determining which one receives a specific number. When one printNumbers goroutine receives a number,
it spends a small amount of time printing that number to the screen, while another goroutine
reads the next number from the channel and does the same thing. When a goroutine has finished
its work of printing the number and is ready to read another number, it will go back and read the channel again
to print the next one. If there are no more numbers to be read from the channel,
it will start to block until the next number can be read. Once generateNumbers has finished
and close() is called on the channel, all three of the
printNumbers goroutines will finish their range loops and exit. When all three goroutines have exited
and called Done on the WaitGroup, the WaitGroup’s count will reach zero and the program will exit.
*/
