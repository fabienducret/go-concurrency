package main

import (
	"concurrency/concurrency"
	"fmt"
	"sync"
)

const hello = "Hello, %s\n"

func main() {
	printWithChannel()
	printWithWaitGroup()
}

func printWithChannel() {
	ch := make(chan string)
	go concurrency.FromChannel(ch)

	toPrint := <-ch
	fmt.Println(fmt.Sprintf(hello, toPrint))
}

func printWithWaitGroup() {
	var wg sync.WaitGroup
	var toPrint string
	wg.Add(1)

	go func() {
		defer wg.Done()
		toPrint = concurrency.FromWaitGroup()
	}()

	wg.Wait()
	fmt.Println(fmt.Sprintf(hello, toPrint))
}
