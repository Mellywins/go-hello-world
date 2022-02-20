package main

import "fmt"

func main() {
	x := 5
	y := 10

	// if else
	if x <= y {
		fmt.Printf("%d is less than or equal to  %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	//else if
	color := "blue"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is neither red or blue")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is neither red or blue")
	}
}
