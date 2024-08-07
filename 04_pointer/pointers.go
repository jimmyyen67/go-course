package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int
	agePointer = &age

	//fmt.Println("Age:", age)
	fmt.Println("Age:", *agePointer) // output 32
	fmt.Println("Age:", agePointer)  // output 0x1400009e008

	editAgeToAdultsYears(&age) // or getAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultsYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
