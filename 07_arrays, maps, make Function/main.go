package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// this make() code will create a slice with 2 empty slots with default value nil and "" for string, 0 for int
	// and the 5 is capacity of the slice with initial in memory
	// so the slice won't need to be re-place in the memory if the actual length of slice is under 5
	userNames := make([]string, 2, 5)

	userNames[0] = "Kent"

	userNames = append(userNames, "John")
	userNames = append(userNames, "Paul")

	fmt.Println(userNames)

	courseRating := make(floatMap, 3)
	courseRating["go"] = 4.8
	courseRating["react"] = 4.9
	courseRating["php"] = 4.5
	courseRating.output()

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for index, value := range courseRating {
		fmt.Println(index, value)
	}
}
