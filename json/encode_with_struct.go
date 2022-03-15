package main

import (
	"encoding/json"
	"fmt"
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

	}

	fmt.Println(string(value))
}
