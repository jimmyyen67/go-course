package arrays

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println("Original prices: ", prices)
	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 29.50, 100.0)
	prices = prices[1:]
	fmt.Println("Updated prices", prices)

	discountPrices := []float64{100.99, 80.99, 20.50}
	prices = append(prices, discountPrices...)
	fmt.Println("Latest prices: ", prices)
}

// introduce of array and slice
//func main() {
//	var productName [4]string = [4]string{"Book"}
//	prices := [4]float64{110.99, 12.99, 16.99, 20.50}
//	fmt.Println(prices)
//
//	productName[2] = "Carpet"
//
//	fmt.Println(productName)
//	fmt.Println(prices[2])
//
//	featuredPrices := prices[1:]
//	featuredPrices[0] = 199.99
//	highlightedPrices := featuredPrices[:1]
//	fmt.Println(featuredPrices)
//	fmt.Println(prices)
//	fmt.Println(len(prices))
//	fmt.Println(len(featuredPrices), cap(featuredPrices))
//	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
//
//	highlightedPrices = highlightedPrices[:3]
//	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
//}
