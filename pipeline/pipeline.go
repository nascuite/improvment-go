package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("№%d, %d \n", i, getNum())
			wg.Done()
		}(i)
	}
	ch <- 0

	wg.Wait()

}

func getNums(len int) []int {
	result := make([]int, len)
	for i, _ := range result {
		result[i] = rand.Int()
	}
	return result
}

var getNumsResult int
var batchLen int32 = 100
var ch = make(chan int32, 1)

func getNum() int {
	i := <-ch
	if i%batchLen == 0 {
		getNumsResult = getNums(1)[0]
	}
	i++
	ch <- i

	return getNumsResult //getNumsResult
}
