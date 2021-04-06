package main

import (
	"sync"

	"github.com/java-akademie/myutils"
)

func main() {
	myutils.H1("start the webservers")

	var wg sync.WaitGroup

	wg.Add(1)
	go webserver1()

	wg.Add(1)
	go webserver2()

	wg.Add(1)
	go webserver3()

	wg.Add(1)
	go webserver4()

	wg.Wait()
}
