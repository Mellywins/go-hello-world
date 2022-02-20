package main

import "fmt"

func main() {
	// main types
	// string
	// bool
	// int
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//byte // alias for uint8
	//rune // alias for int32
	//float32 float64
	//complex64 complex128

	// Using var
	var age int = 37
	const isCool = true
	name := "Alex"
	size := 1.3
	name = "Alex²"
	fmt.Println("Hello "+name+" Your age is : ", age, isCool, " Your size is: ", size)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)

}
