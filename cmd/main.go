package main

import (
	"Rest-Api/models"
	"Rest-Api/pkg/repository"
	"Rest-Api/pkg/services"
	"encoding/json"
	//"Rest-Api/pkg/repository"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	num := 2
	gorm := repository.Init()
	err := gorm.AutoMigrate(&models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("models no migrate")
	}
	repo := repository.NewRepository(*gorm)
	service := services.NewService(repo)

	var postJS models.Post
	request := RequestPost(num)

	if err = json.Unmarshal(request, &postJS); err != nil {
		log.Fatalf("cant Unmarshal JSON: %s", err.Error())
	}

	post := models.Post{
		Id:     postJS.Id,
		UserId: postJS.UserId,
		Title:  postJS.Title,
		Body:   postJS.Body,
	}
	id, err := service.Post.Add(post)

	if err != nil {
		log.Fatal(err.Error())
	}
	requestComment := RequestComment(id)
	var commentsJS models.Comments

	if err = json.Unmarshal(requestComment, &commentsJS); err != nil {
		log.Fatalf("cant Unmarshal JSON: %s", err.Error())
	}
	for i, _ := range commentsJS {
		comment := models.Comment{
			PostID: commentsJS[i].PostId,
			Name:   commentsJS[i].Name,
			Email:  commentsJS[i].Email,
			Body:   commentsJS[i].Body,
		}
		_, err := service.Comment.Add(comment)
		if err != nil {
			log.Fatalf("No comments add to db:%s", err.Error())
		}
	}
}
func RequestPost(i int) []byte {
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
