package main

import "fmt"

func main() {
	a := 5
	b := &a
	fmt.Println(*b)
	fmt.Printf("%T\n", b)

	// change val with pointer
	*b = 10
	fmt.Println(a)
}
