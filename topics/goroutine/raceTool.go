package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	fmt.Println("race condition")

	// mut.RLock()
	var score = []int{0}
	// mut.Unlock()

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {

		fmt.Println("one r")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {

		fmt.Println("two r")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {

		fmt.Println("three r")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {

		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()

	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
