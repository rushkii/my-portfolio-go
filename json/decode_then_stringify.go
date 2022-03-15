package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Data struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	var data Data
	content := []byte(`{
		"username": "kiizuha",
		"password": "KznF2ksd!320A."
	}`)

	if err := json.Unmarshal(content, &data); err != nil {
		log.Println(err)
	}

	value, err := json.MarshalIndent(&data, "", "    ")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(value))
}
