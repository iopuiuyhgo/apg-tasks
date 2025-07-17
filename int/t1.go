package main

import "fmt"

//Условие задачи
//Дан массив целых чисел nums и целое число k. Нужно написать функцию,
//которая вынимает из массива nums k наиболее часто встречающихся элементов.

//Пример
//# ввод
//nums = [1,1,1,2,2,3]
//k = 2
//# вывод (в любом порядке)
//[1, 2]

func topKFrequentElements(nums []int, k int) []int {
	m := make(map[int]int)
	l := make([][]int, len(nums)+1)
	for _, v := range nums {
		m[v] += 1
	}

	for k, v := range m {
		l[v] = append(l[v], k)
	}

	var res []int

	for i := len(nums); i > 0; i -= 1 {
		for _, v := range l[i] {
			if len(res) >= k {
				return res
			}
			res = append(res, v)
		}
	}
	return res
}

func main() {
	inp := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(inp, 2)
	fmt.Println(topKFrequentElements(inp, 2))
}
