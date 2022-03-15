package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	value, err := json.Marshal(map[string]string{
		"username": "kiizuha",
		"password": "KznF2ksd!320A.",
	})
	if err != nil {

	}

	fmt.Println(string(value))
}
