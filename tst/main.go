package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//
//func getNum(key string) int {
//	res, err := strconv.Atoi(key)
//	if err != nil {
//		panic(err)
//	}
//
//	return res
//}

func getNums(keys []string) []int {
	time.Sleep(time.Second)
	res := make([]int, len(keys))
	for i, key := range keys {
		converted, err := strconv.Atoi(key)
		if err != nil {
			panic(err)
		}
		res[i] = converted
	}
	fmt.Printf("processed batch: %d\n", len(keys))

	return res
}

type BatchBuilder struct {
	limit int

	data map[string]chan int
	f    func([]string) []int
	mu   sync.Mutex
}

func NewBatchBuilder(limit int, onAggregate func([]string) []int) *BatchBuilder {
	return &BatchBuilder{
		limit: limit,
		data:  make(map[string]chan int),
		f:     onAggregate,
	}
}

func (b *BatchBuilder) GetNum(key string) int {
	resCh := make(chan int, 1)
	b.mu.Lock()
	b.data[key] = resCh
	if len(b.data) >= b.limit {
		data := b.data
		go b.processOperations(data)
		b.data = make(map[string]chan int)
	}
	b.mu.Unlock()

	return <-resCh
}

func (b *BatchBuilder) processOperations(operations map[string]chan int) {
	keys := make([]string, 0, len(operations))
	for key := range operations {
		keys = append(keys, key)
	}
	results := b.f(keys)

	for i, res := range results {
		key := keys[i]
		operations[key] <- res
	}
}

func main() {
	b := NewBatchBuilder(100, getNums)
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			key := strconv.Itoa(i)
			_ = b.GetNum(key)
		}()
	}
	wg.Wait()
}
