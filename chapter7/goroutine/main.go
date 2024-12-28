package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	printLater("Hello\n", time.Millisecond*100)
	printLater("World\n", time.Millisecond*100)
	printLater(os.Getenv("USER")+"\n", time.Millisecond*100)
}
func printLater(msg string, duration time.Duration) {
	time.Sleep(duration)

	fmt.Printf(msg)
}
