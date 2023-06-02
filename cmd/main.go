package main

import (
	"concurrency/concurrency"
	"log"
	"sync"
)

var list = []string{"Hello", "Coucou", "Bonjour"}

func main() {
	printWithChannel()
	printWithWaitGroup()
}

func printWithChannel() {
	result := make(chan string)

	for _, value := range list {
		go concurrency.FromChannel(result, value)
	}

	for range list {
		toPrint := <-result
		log.Println(toPrint)
	}
}

func printWithWaitGroup() {
	wg := sync.WaitGroup{}

	for _, value := range list {
		wg.Add(1)

		go func(value string) {
			defer wg.Done()
			toPrint := concurrency.FromWaitGroup(value)
			log.Println(toPrint)
		}(value)
	}

	wg.Wait()
}
