package main

import "fmt"

func main() {
	var fruits [2]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fmt.Println(fruits[0], fruits[1])

	fruitArr := [2]string{"Apple", "Orange"}
	fmt.Println(fruitArr)

	// slices: Arrays with unfixed length
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[2:])
}
