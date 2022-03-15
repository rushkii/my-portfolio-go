package main

import "fmt"

type Users struct {
	Name string
	Age  int
}

func (u *Users) AddAge(n int) {
	u.Age += n

}

func main() {
	user := &Users{
		Name: "Kiizuha",
		Age:  21,
	}
	fmt.Printf(
		"%s's age is %d years old right now, next year his age will be ",
		user.Name, user.Age,
	)

	user.AddAge(1)
	fmt.Printf("%d years old", user.Age)
}
