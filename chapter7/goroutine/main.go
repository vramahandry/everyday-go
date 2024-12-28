package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	workItems := 5
	wg.Add(workItems)
	for i := 1; i <= workItems; i++ {
		j := i
		go func() {
			defer wg.Done()
			printLater(
				fmt.Sprintf("Hello from %d\n", j),
				time.Millisecond*100)
		}()
	}
	wg.Wait()
}

func printLater(msg string, duration time.Duration) {
	time.Sleep(duration)

	fmt.Printf(msg)
}
