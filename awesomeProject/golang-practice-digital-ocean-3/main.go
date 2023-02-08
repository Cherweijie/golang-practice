package main

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, ch chan<- int, wg *sync.WaitGroup) {
	// write only channel -> prevent deadlock
	defer wg.Done()

	// loop to write numbers into the channel
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("sending %d to channel\n", idx)
		ch <- idx
	}
}

func printNumbers(ch <-chan int, wg *sync.WaitGroup) {
	// read only channel -> prevent deadlock
	defer wg.Done()

	// loop to print numbers from channel
	for num := range ch {
		fmt.Printf("read %d from channel\n", num)
	}
}

func main() {
	var wg sync.WaitGroup
	numberChan := make(chan int)

	wg.Add(2)
	go printNumbers(numberChan, &wg)

	generateNumbers(3, numberChan, &wg)

	// important to close channel at the end
	close(numberChan)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}

// Logic for deadlock - pretty straightforward
/*
Even though these types could be a chan int, which would allow both reading and writing,
it can be helpful to constrain them to only what the function needs to avoid accidentally causing
your program to stop running from something known as a deadlock. A deadlock can happen when one part
of a program is waiting for another part of the program to do something, but that other part of the
program is also waiting for the first part of the program to finish. Since both parts of the program
are waiting on each other, the program will never continue running, almost like when two gears seize.
*/
