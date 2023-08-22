package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// 2) В func main() написать цикл 10000 раз вызывающий в горутине функцию getNum() и печатающий индекс вызова и результат в консоль
func main() {
	requests := make(chan request)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go batching(ctx, requests)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			r := request{
				key:      strconv.Itoa(i),
				response: make(chan result),
			}
			var res result
			select {
			case requests <- r:
				res = <-r.response
			case <-ctx.Done():
				return
			}

			fmt.Printf("№: %v = %v \n", i, res.value)
		}()
	}
	wg.Wait()
}

func batching(ctx context.Context, requests chan request) {
	batch := make([]request, 0, 5)

	send := func() {
		keys := make([]string, 0, len(batch))
		for i := 0; i < len(batch); i++ {
			keys = append(keys, batch[i].key)
		}
		results := getNums(keys)

		for i := 0; i < len(results); i++ {
			batch[i].response <- result{value: results[i]}
		}
		batch = make([]request, 0, 5)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		case req := <-requests:
			batch = append(batch, req)
			if len(batch) >= 5 {
				send()
			}
		}
	}
}

type request struct {
	key      string
	response chan result
}

type result struct {
	value int
}

// 1) Написать функцию getNum() int с реализацией return getNums(1)[0]
// 3) Переписать getNum не меняя интерфейс так, чтобы getNums вызывалась 1 раз на 100 вызовов getNum.
// (называется batch processing или pipeline call, можно создавать и использовать переменные
// уровня файла общие между функциями)
func getNum(key string) int {
	time.Sleep(time.Second)
	return getNums([]string{"1", "2", "3"})[0]
}

// 0) Написать и реализовать функцию getNums(keys []string]) []int, возвращающую массив рандомных чисел длинны len(keys)
func getNums(keys []string) []int {
	time.Sleep(time.Second)
	len := len(keys)
	result := make([]int, len)

	for i := 0; i < len; i++ {
		result[i] = rand.Int()
	}

	return result
}
