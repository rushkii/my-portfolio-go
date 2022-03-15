package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data map[string]interface{}
	content := []byte(`{
		"username": "kiizuha",
		"password": "KznF2ksd!320A."
	}`)

	if err := json.Unmarshal(content, &data); err != nil {
		log.Println(err)
	}

	fmt.Println(data)
}
