package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		fileName := fmt.Sprintf("storage/posts/%d.txt", i+1)
		file, err := os.Create(fileName)
		result := Request(i)
		_, err = file.Write(result)
		if err != nil {
			log.Fatalf("failed to write: %v\n", err)
		}
		file.Close()
	}

	time.Sleep(time.Second * 10)
}
func Request(i int) []byte {
	request := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i+1)
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
	return readAll
}
