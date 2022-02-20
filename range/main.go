package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, id := range ids {
		fmt.Printf("index %d has %d \n", i, id)
	}
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum is", sum)

	// Range with maps
	savedEmails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@hotmail.fr"}
	for k, v := range savedEmails {
		fmt.Println(k, v)
	}
}
