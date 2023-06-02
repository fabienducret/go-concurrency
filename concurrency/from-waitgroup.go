package concurrency

import (
	"log"
	"time"
)

func FromWaitGroup() string {
	log.Println("execute FromWaitGroup")
	time.Sleep(1 * time.Second)

	return "from WaitGroup"
}
