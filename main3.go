package main

import (
	"fmt"
	"log"
	"net/url"
)

const (
	PATH_TEMPLATE = "/leads/%s/specification"
)

func main() {
	leadId := "123"
	basePath := "http://example.com"
	path := fmt.Sprintf(PATH_TEMPLATE, leadId)
	requestPath, err := url.JoinPath(basePath, path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(requestPath)
}
