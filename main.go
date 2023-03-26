package main

import (
	"hello-world/load"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	for true {
		bar := load.New()
		bar.Loading(2)
	}
	wg.Wait()
}
