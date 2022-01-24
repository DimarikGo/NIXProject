package main

import (
 	"Rest-Api"
	"Rest-Api/pkg/repository"
	"Rest-Api/pkg/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	num := 2

	db, err := repository.NewMysqlDb(repository.Config{
		Name:     "root",
		Password: "11111111",
		DBName:   "NIXdb",
	})
	if err != nil {
		log.Fatalf("failed init db: %s ", err.Error())
	}
	repos := repository.NewRepository(db)

	var postJ Rest_Api.Post
	request := RequestPost(num)

	if err = json.Unmarshal(request, &postJ); err != nil {
		log.Fatalf("cant Unmarshal JSON: %s", err.Error())
	}

	service := service.NewService(repos)

	post1 := Rest_Api.Post{
		UserId: postJ.UserId,
		Id:     postJ.Id,
		Title:  postJ.Title,
		Body:   postJ.Body,
	}

	d, err := service.AddP(post1)
	

	if err != nil {
		log.Fatalf("failed add Post to db: %s ", err.Error())
	}
	RequestComment := RequestComment(d)
	var comments Rest_Api.Comments

	json.Unmarshal(RequestComment, &comments)

	for _, com := range comments {
		comment := Rest_Api.Comment{
			PostId: com.PostId,
			Id:     com.Id,
			Name:   com.Name,
			Email:  com.Email,
			Body:   com.Body,
		}
		go func(c Rest_Api.Comment) {
			_, err := newService.AddC(comment)
			if err != nil {
				return
			}
		}(comment)
	}
	time.Sleep(1 * time.Second)
}

func RequestPost(i int) []byte {
	fmt.Println(i)
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
	return readAll
}

func RequestComment(i int) []byte {
	fmt.Println(i)
	request := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%v", i)
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
