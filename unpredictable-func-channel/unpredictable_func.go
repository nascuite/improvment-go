package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

// UnpredictableFunc ...
// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
func predictableFunc() (int64, error) {
	start := time.Now()
	defer func() { fmt.Println("duration: ", time.Since(start)) }()

	ch := make(chan int64)
	go func() {
		ch <- unpredictableFunc()
	}()

	select {
	case <-time.After(time.Second):
		return 0, fmt.Errorf("%v", "timeout reached")
	case res := <-ch:
		return res, nil
	}
}

func main() {
	fmt.Println("started")

	fmt.Println(predictableFunc())
}
