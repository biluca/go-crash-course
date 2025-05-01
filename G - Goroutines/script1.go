package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var waitGroup = sync.WaitGroup{}

var cacheData = []int{10, 20, 30, 40, 50}
var resultCache = []int{}

func main() {

	t0 := time.Now()

	for i := 0; i < len(cacheData); i++ {
		waitGroup.Add(1)
		go cacheCall(i)
	}
	waitGroup.Wait()
	fmt.Println("Time taken to fetch data from cache: ", time.Since(t0))
	fmt.Println("The Result from Cache is ", resultCache)
}

func cacheCall(i int) {
	var delay float32 = 3000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The Result from Cache is ", cacheData[i])

	mutex.Lock()
	resultCache = append(resultCache, cacheData[i])
	mutex.Unlock()

	waitGroup.Done()
}
