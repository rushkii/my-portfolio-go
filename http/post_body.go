package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var client = &fasthttp.Client{}

func main() {
	var data map[string]interface{}

	res := fasthttp.AcquireResponse()
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.SetBody([]byte(`{
		"username": "kiizuha",
		"password": "kiizuha"
	}`))
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetRequestURI("https://httpbin.org/post")
	if err := client.Do(req, res); err != nil {
		log.Panicln(err)
	}

	if err := json.Unmarshal(res.Body(), &data); err != nil {
		log.Println(err)
	}

	value, err := json.MarshalIndent(&data, "", "    ")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(value))
}
