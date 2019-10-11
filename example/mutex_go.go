package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mutex sync.Mutex
	wait := sync.WaitGroup{}

	fmt.Println("Locked")
	mutex.Lock()

	for i := 1; i <= 3; i++ {
		wait.Add(1)

		go func(i int) {
			fmt.Println("Not lock:", i)

			mutex.Lock()
			fmt.Println("Lock:", i)

			time.Sleep(time.Second)

			fmt.Println("Unlock:", i)
			mutex.Unlock()

			defer wait.Done()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock()

	wait.Wait()

	lockAndR()
}

func lockAndR() {
	mutex := &sync.RWMutex{}
	var a int
	mutex.Lock()
	a = 4
	fmt.Println("Lock:", a)
	go func() {
		mutex.RLock()
		a = 10
		fmt.Println("RLock:", a)
		mutex.RUnlock()
	}()
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	time.Sleep(time.Second * 1)
}
