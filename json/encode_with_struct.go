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
	value, err := json.Marshal(&Data{
		Username: "kiizuha",
		Password: "KznF2ksd!320A.",
	})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(value))
}
