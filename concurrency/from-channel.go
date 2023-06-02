package concurrency

import (
	"log"
)

func FromChannel(ch chan string) {
	defer close(ch)

	log.Println("execute FromChannel")

	ch <- "from channel"
}
