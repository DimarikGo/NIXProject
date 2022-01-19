package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go Request(i)
	}
	time.Sleep(time.Second * 4)
}

func Request(i int) {
	request := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i)
	response, err := http.Get(request)
	if err != nil {
		//todo logger
		log.Fatal(err)
	}
	readAll, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//todo logger
		log.Fatal(err)
	}
	fmt.Println(string(readAll))
}
