package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, value := range nums {
		for j, value1 := range nums {
			if i != j && (value+value1) == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
