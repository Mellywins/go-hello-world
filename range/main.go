package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, id := range ids {
		fmt.Printf("index %d has %d \n", i, id)
	}
}
