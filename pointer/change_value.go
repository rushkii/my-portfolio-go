package main

import "fmt"

func main() {
	var alpha *int

	beta := 10
	alpha = &beta // *alpha 10

	fmt.Println(*alpha)
	fmt.Println(beta)
}
