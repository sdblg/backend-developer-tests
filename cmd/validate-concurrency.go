package main

import (
	"github.com/stackpath/backend-developer-tests/concurrency"
	"math/rand"
	"time"
)

func main() {

	sp := concurrency.NewSimplePool(4)

	var task = func() {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(5) // random n will be 0 to 5
		time.Sleep(time.Duration(n) * time.Second)

	}

	sp.Submit(task)

}
