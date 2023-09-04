package main

import (
	"fmt"
	"sync"
)

// Находим максимальное четное число
// Сделай code review, найди проблемы
func main() {
	var max int
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 1000; i > 0; i-- {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			wg.Done()
			if i%2 == 0 && i > max {
				max = i
			}
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	fmt.Printf("Maximum is %d", max)
}

// 1. нужно явно передать i в функцию, иначе у i значение на момент исполнения, а не запуска функции
// 2. гонка при обращении к i, нужно добавить Mutex
// go run -race ./maximum_even_number
