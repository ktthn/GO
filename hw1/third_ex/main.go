package main

import (
	"fmt"
	"sort"
)

func main() {

	/* 3.Написать функцию для сравнения слайса с цело-числовыми значениями
	a. сравнивает два слайса
	b. возвращает булево значение: совпало или нет
	c. порядок не важен
	d. порядок важен
	*/

	a := []int{5, 7, 8}
	b := []int{8, 7, 5}
	fmt.Println(equal(a, b))
	fmt.Println(matters(a, b))
	fmt.Println(withOrder_2(a, b))
}

// order matters
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, num := range a {
		if num != b[i] {
			return false
		}
	}
	return true
}

// order doesn't matter

func matters(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for i, num := range a {

		if num != b[i] {
			return false
		}
	}
	return true
}

// реализация с мапой, практис с Мадьяром
func withOrder_2(a, b []int) bool {

	aMap := make(map[int]int, len(a))
	bMap := make(map[int]int, len(b))

	for _, aVal := range a {
		aMap[aVal]++
	}

	for _, bVal := range b {
		bMap[bVal]++
	}

	for aKey, aVal := range aMap {
		if bMap[aKey] != aVal { //
			return false
		}
	}

	return true
}
