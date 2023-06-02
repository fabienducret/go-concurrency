package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func FromWaitGroup(value string) string {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	return fmt.Sprintf("%s from wait group", value)
}
