package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Можешь ли ты предложить оптимизацию работы с памятью
// Еще можно вынести bytes за пределы цикла, тогда будет меньше аллокаций
func main() {
	one()

	fmt.Println("finished")
}

const (
	sentenceCount = 10_000_000
	path          = "text.txt"
)

func one() {

	for i := 0; i < 10000; i++ {
		bytes, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			//panic(err)
		}

		if strings.Contains(string(bytes), "malware") {
			fmt.Println(i, "contains malware")
		}
	}
}

func second() {
	for i := 0; i < 10000; i++ {
		f, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}

		sc := bufio.NewScanner(f)
		for sc.Scan() {
			if strings.Contains(sc.Text(), "malware") {
				fmt.Println(i, "contains malware")
			}
		}
		f.Close()
	}
}
