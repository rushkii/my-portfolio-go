package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	value, err := json.Marshal(map[string]string{
		"username": "kiizuha",
		"password": "KznF2ksd!320A.",
	})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(value))
}
