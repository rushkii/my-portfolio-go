package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var client = &fasthttp.Client{}

func main() {
	res := fasthttp.AcquireResponse()
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.Header.SetMethod("GET") // GET, POST, HEAD, OPTIONS, DELETE, PUT, PATCH
	req.SetRequestURI("https://httpbin.org/get")
	if err := client.Do(req, res); err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(res.Body()))
}
