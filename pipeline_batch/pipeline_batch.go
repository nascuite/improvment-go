package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	defer func() { fmt.Printf("time: %v", time.Since(now)) }()

	b := NewBatch()
	wg := sync.WaitGroup{}
	for i := 0; i < b.maxLen; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := strconv.Itoa(i)

			fmt.Println("№", i, "result:", b.getNum(key))
		}(i)
	}
	wg.Wait()
}

type batch struct {
	maxLen   int
	batchLen int
	keyMap   map[string]chan int
	mu       sync.Mutex
}

func NewBatch() *batch {
	return &batch{
		maxLen:   10000,
		batchLen: 100,
		keyMap:   make(map[string]chan int),
	}
}

func (b *batch) getNum(key string) int {
	ch := make(chan int, 1)
	b.mu.Lock()
	b.keyMap[key] = ch
	if len(b.keyMap) >= b.batchLen {
		data := b.keyMap
		go b.processing(data)
		b.keyMap = make(map[string]chan int)
	}
	b.mu.Unlock()

	return <-ch
}

func (b *batch) processing(data map[string]chan int) {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}

	results := getNums(keys)

	for i, key := range keys {
		data[key] <- results[i]
	}
}

func getNums(keys []string) []int {
	time.Sleep(time.Second)

	var result []int
	for i := 0; i < len(keys); i++ {
		result = append(result, rand.Int())
	}

	return result
}
