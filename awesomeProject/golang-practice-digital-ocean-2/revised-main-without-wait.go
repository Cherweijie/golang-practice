package main

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, wg *sync.WaitGroup) {
	// decreases the count stored in waitGroup by 1
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
	}
}

func printNumbers(wg *sync.WaitGroup) {
	// decreases the count stored in waitGroup by 1
	defer wg.Done()

	for idx := 1; idx <= 3; idx++ {
		fmt.Printf("Printing number %d\n", idx)
	}
}

func main() {
	var wg sync.WaitGroup

	// indicate 2 Done calls are required before considering that the group is finished
	wg.Add(2)

	// go routine to handle printing numbers
	go printNumbers(&wg)

	// go routine to handle generating numbers
	go generateNumbers(3, &wg)

	fmt.Println("Waiting for goroutines to finish...")
	fmt.Println("Done!")
}

// Result: Normally only the last 2 lines are printed
