package main

import "fmt"

func isUgly(m int) bool {
	if m <= 0 {
		return false
	}
	if m == 1 {
		return true
	}
	for m%2 == 0 {
		m /= 2
	}
	for m%3 == 0 {
		m /= 3
	}
	for m%5 == 0 {
		m /= 5
	}
	return m==1
	return false
}
func main() {
	fmt.Println("solved leetcode 263 ugly number")
	num := 6
	fmt.Printf("%v is a ugly number or not? %v", num, isUgly(num))
}
