package main

import (
	"fmt"
	"sync"
)

func onlyOnce() {
	var once sync.Once
	done := make(chan bool)
	onceBody := func() {
		fmt.Println("Only once")
	}
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

func waitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

var data = struct {
	lock    *sync.RWMutex
	payload int
}{
	lock:    new(sync.RWMutex),
	payload: 0,
}

func lockData() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			data.lock.Lock()
			defer data.lock.Unlock()
			data.payload = i
			wg.Done()
		}(i)
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			data.lock.RLock()
			defer data.lock.RUnlock()
			fmt.Println(data.payload)
			wg.Done()
		}(i)
	}
}

func main() {
	lockData()
}
