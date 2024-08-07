package funcsarevalues

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	otherNumbers := []int{5, 1, 2, 6}

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&otherNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	otherTransformNumbers := transformNumbers(&otherNumbers, transformerFn2)
	fmt.Println(transformedNumbers)
	fmt.Println(otherTransformNumbers)
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	var doubled []int
	for _, value := range *numbers {
		doubled = append(doubled, transform(value))
	}
	return doubled
}

func getTransformerFunction(numbers *[]int) transformFunc {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
