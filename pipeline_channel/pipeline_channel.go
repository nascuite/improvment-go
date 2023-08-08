package main

import (
	"fmt"
	"sync"
	"time"
)
import "math/rand"

func main() {
	now := time.Now()
	defer func() { fmt.Printf("duration: %v \n", time.Since(now)) }()
	var i int
	for val := range middle(getNum()) {
		fmt.Printf("№%v, %v \n", i, val)
		i++
	}
}

var batchLen int = 100
var cnt int = 10000

func getNum() (<-chan int, <-chan []int) {
	out := make(chan int)

	go func() {
		for i := 0; i < cnt; i++ {
			out <- i
		}
		close(out)
	}()

	outBatchNums := make(chan []int)
	batchCnt := cnt/batchLen + 1
	go func() {
		mu := sync.WaitGroup{}
		for j := 0; j < batchCnt; j++ {
			mu.Add(1)
			go func() {
				res := getNums(batchLen)
				outBatchNums <- res
				mu.Done()
			}()
		}
		mu.Wait()
		close(outBatchNums)
	}()

	return out, outBatchNums
}

func middle(in <-chan int, inNums <-chan []int) chan int {
	out := make(chan int)

	go func() {
	LOOP:
		for nums := range inNums {
			for _, v := range nums {
				_, exists := <-in
				if !exists {
					break LOOP
				}
				out <- v
			}
		}

		close(out)
	}()

	return out
}

func getNums(len int) []int {
	time.Sleep(time.Second)

	result := make([]int, len)
	for i, _ := range result {
		result[i] = rand.Int()
	}
	return result
}
