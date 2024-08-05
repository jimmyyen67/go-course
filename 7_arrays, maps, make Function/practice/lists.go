package practice

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	// 1). create a new array with three of your hobbies
	hobbies := [3]string{"jogging", "lifting", "running"}
	fmt.Println(hobbies)

	// 2). output first alone, output second and third as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3). create a slice based on first element that contains the first and second elements
	mainHobbies := hobbies[:2] // or hobbies[0:2]
	fmt.Println(mainHobbies)

	// 4). re-slice the slice from (3) and change it to contain the second and last element of original array
	newSlice := mainHobbies[1:3]
	//newSlice = append(newSlice, hobbies[2])
	fmt.Println(newSlice)

	// 5). create a dynamic array that contains your goals (at least 2 of your goals)
	goals := []string{"being smart", "eating less"}
	fmt.Println(goals)

	// 6). set the second goal to a different and then add a third goal to the existing dynamic array
	goals[1] = "losing weight"
	goals = append(goals, "helping other people")
	fmt.Println(goals)

	// 7). create a Product struct with title, id, price and create a dynamic list of products (at least 2 produtcs)
	//     then add a third product to the existing list of products
	products := []Product{
		{
			title: "Book",
			id:    "1",
			price: 12.50,
		},
		{
			title: "Pen",
			id:    "2",
			price: 19.99,
		},
	}
	newProduct := Product{
		title: "Pillow",
		id:    "3",
		price: 18.15,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}
