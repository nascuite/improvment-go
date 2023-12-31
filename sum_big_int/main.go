package main

import "fmt"

//-------------------------------------------------------------------------------------
//	Условие задачи на алгоритмы
//-------------------------------------------------------------------------------------
// Мы хотим складывать очень большие числа, которые превышают емкость базовых типов,
// поэтому мы храним их в виде массива неотрицательных чисел.
// Нужно написать функцию, которая примет на вход два таких массива,
// вычислит сумму чисел, представленных массивами, и вернет результат в виде такого же массива.
// Пример:
// Пример 1
// ввод
// arr1 = [1, 2, 3] # число 123
// arr2 = [4, 5, 6] # число 456
// вывод
// res = [5, 7, 9] # число 579. Допустим ответ с первым незначимым нулем [0, 5, 7, 9]
// Пример 2
// ввод
// arr1 = [5, 4, 4] # число 544
// arr2 = [4, 5, 6] # число 456
// вывод
// res = [1, 0, 0, 0] # число 1000

// arr1 = [5, 4, 4] # число 544
// arr2 = [5, 6] # число 56
// # то вывод будет валидный:
// res = [ 6, 0, 0] # число 600
// или
// res = [0, 6, 0, 0] # число 600
func main() {
	arr1 := []int{5, 4, 4}
	arr2 := []int{5, 6}
	res := sum(arr1, arr2)
	fmt.Println(res)
}

func sum(arr1, arr2 []int) []int {
	sdvg := 0
	lenArr1 := len(arr1)
	lenArr2 := len(arr2)

	maxLen := lenArr1
	if lenArr1 < lenArr2 {
		maxLen = lenArr2
	}

	result := make([]int, maxLen+1)
	for i := 1; i <= maxLen+1; i++ {
		arg1 := 0
		if lenArr1 < i {
			arg1 = 0
		} else {
			arg1 = arr1[lenArr1-i]
		}

		arg2 := 0
		if lenArr2 < i {
			arg2 = 0
		} else {
			arg2 = arr2[lenArr2-i]
		}

		res := arg1 + arg2 + sdvg

		sdvg = 0
		if res > 9 {
			ost := res - 10
			sdvg = 1
			result[maxLen-i+1] = ost
		} else {
			result[maxLen-i+1] = res
		}
	}

	return result
}
