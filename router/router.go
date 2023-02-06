package router

import (
	"errors"
	"log"
	"net/http"
)

type HandleRequestFunction = func(request *http.Request) error

type RequestHandler struct {
	HandlerFunction HandleRequestFunction
}

func validateRequest(request *http.Request) error {
	return errors.New("Hello")
}

func (handler *RequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := validateRequest(request)
	if err != nil {
		log.Println("An validation error occurred.")
	}

	err = handler.HandlerFunction(request)
}
