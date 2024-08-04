package main

import "fmt"

func main() {
	var productName [4]string = [4]string{"Book"}
	prices := [4]float64{110.99, 12.99, 16.99, 20.50}
	fmt.Println(prices)

	productName[2] = "Carpet"

	fmt.Println(productName)
	fmt.Println(prices[2])
}
