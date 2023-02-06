package main

import (
	"awesomeProject/router"
	"net/http"
)

func handler(request *http.Request) error {
	return nil
}

func main() {
	requestHandler := &router.RequestHandler{
		HandlerFunction: handler,
	}
	_ = http.ListenAndServe(":9001", requestHandler)
}
