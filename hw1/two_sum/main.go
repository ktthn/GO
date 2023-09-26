package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := 5
	fmt.Print(twoSum(a, b))
}

// O(N)^2
//func twoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return []int{}
//}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums)) //обыяное объявлние мапы, ассоциативный массив хранит по ключу, как дикшинари в питоне

	for i := 0; i < len(nums); i++ {

		match := target - nums[i]
		if idx, ok := m[match]; ok {
			return []int{i, idx}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}
