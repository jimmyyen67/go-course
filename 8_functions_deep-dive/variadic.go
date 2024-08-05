package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(2, 14, 56)
	fmt.Println(sum) // output: 72

	anotherSum := sumup(7, numbers...)
	fmt.Println(anotherSum) // output: 33
}

func sumup(firstValue int, numbers ...int) int {
	sum := firstValue
	for _, value := range numbers {
		sum += value
	}
	return sum
}
