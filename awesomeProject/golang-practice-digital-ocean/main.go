package main

import (
	"fmt"
	"sync"
)

func GenerateNumbers(total int, wg *sync.WaitGroup) {
	// decreases the count stored in waitGroup by 1
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
	}
}

func PrintNumbers(wg *sync.WaitGroup) {
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
	go PrintNumbers(&wg)

	// go routine to handle generating numbers
	go GenerateNumbers(3, &wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}

// Result: Normally, numbers are generated and printed in order
