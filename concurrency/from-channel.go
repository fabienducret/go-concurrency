package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func FromChannel(result chan<- string, value string) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	result <- fmt.Sprintf("%s from channel", value)
}
