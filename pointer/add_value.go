package main

import "fmt"

func Add(x int, y int) int {
	return x + y
}

func Add2(x *int, y int) {
	*x = *x + y
}

func main() {
	// Non-pointer method
	number := 10
	add := 5
	fmt.Printf("Add number without pointer method: %d + %d = ", number, add)
	number = Add(number, add)
	fmt.Printf("%d\n", number)

	// Pointer method
	add = 10
	fmt.Printf("Add number with pointer method: %d + %d = ", number, add)
	Add2(&number, add)
	fmt.Printf("%d", number)
}
