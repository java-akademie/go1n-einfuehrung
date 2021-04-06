package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/java-akademie/myutils"
)

func doMutex() {
	myutils.H2("mutex")
	test(false, 20)
	test(true, 20)
}

func test(lock bool, max int) {
	myutils.H2(fmt.Sprintf("%v Millionen Durchlaeufe, Lock: %v", max, lock))

	start := time.Now().UnixNano()

	var account = 0
	var wg sync.WaitGroup
	var mutex sync.Mutex

	max *= 1_000_000
	for i := 1; i <= max; i++ {
		wg.Add(1)
		go changeAccount(lock, i, &account, &wg, &mutex,
			myutils.GetRandom2(1, 1000))
	}
	wg.Wait()
	fmt.Printf("\nAccount: %d \n", account)
	duration := time.Now().UnixNano() - start
	fmt.Println("Duration:", duration/1_000_000_000)
}

func changeAccount(lock bool, i int, account *int, wg *sync.WaitGroup, mutex *sync.Mutex, value int) {
	defer wg.Done()

	if lock {
		mutex.Lock()
	}

	*account = *account + value
	fmt.Print("")
	*account = *account - value

	if lock {
		mutex.Unlock()
	}
}
