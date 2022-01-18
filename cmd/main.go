package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	Request()
}

func Request() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		//todo logger
		log.Fatal(err)
	}
	readAll, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//todo logger
		log.Fatal(err)
	}
	fmt.Println(readAll)
}
