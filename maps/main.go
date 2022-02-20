package main

import "fmt"

func main() {
	emails := make(map[string]string)
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@hotmail.fr"
	emails["Mike"] = "mike@yahoo.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	delete(emails, "Bob")
	fmt.Println(emails)

	// Declare map and add kv
	savedEmails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@hotmail.fr"}
	fmt.Println(savedEmails)
}
