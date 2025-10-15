package main

import "fmt"

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	if n == 1 {
		return true
	}
	return false
}
func main() {
	fmt.Println("solved leetcode 263 ugly number")
	num := 6
	fmt.Printf("%v is a ugly number or not? %v", num, isUgly(num))
}
