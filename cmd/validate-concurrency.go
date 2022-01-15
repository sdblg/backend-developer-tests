package main

import (
	"context"
	"fmt"
	"github.com/stackpath/backend-developer-tests/concurrency"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Concurrency - SimplePool")
	fmt.Println()

	sp := concurrency.NewSimplePool(4)

	var task = func() {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(5) // random n will be 0 to 5
		time.Sleep(time.Duration(n) * time.Second)

	}

	sp.Submit(task)
	fmt.Println()

	//---------- Advanced Pool ------------
	fmt.Println("SP// Backend Developer Test - Concurrency - AdvancedPool")
	fmt.Println()

	var taskWithContext = func(ctx context.Context) {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(5) // random n will be 0 to 5
		time.Sleep(time.Duration(n) * time.Second)

	}

	asp, _ := concurrency.NewAdvancedPool(10, 4)
	//ctx, cancelContext := context.WithCancel(context.Background())
	//cancelContext()
	ctx, _ := context.WithCancel(context.Background())

	err := asp.Submit(ctx, taskWithContext)
	if err != nil {
		fmt.Println("ERROR on Submit: ", err)
		return
	}

	err = asp.Close(ctx)
	if err != nil {
		fmt.Println("ERROR on Close: ", err)
		return
	}

}
