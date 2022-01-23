package main

import (
	"Rest-Api"
	"Rest-Api/pkg/repository"
	"Rest-Api/pkg/service"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	//for i := 0; i < 100; i++ {
	//	fileName := fmt.Sprintf("storage/posts/%d.txt", i+1)
	//	file, err := os.Create(fileName)
	//	result := Request(i)
	//	_, err = file.Write(result)
	//	if err != nil {
	//		log.Fatalf("failed to write: %v\n", err)
	//	}
	//	file.Close()
	//}
	post1 := Rest_Api.Post{
		UserId: 13,
		Title:  "131 post",
		Body:   "131 post",
	}
	db := repository.NewMysqlDb(repository.Config{
		Name:     "root",
		Password: "11111111",
		DBName:   "NIXdb",
	})
	//if err != nil {
	//	log.Fatalf("failed init db: %s ", err.Error())
	//}
	repos := repository.NewRepository(db)
	newService := service.NewService(repos)
	p, err := newService.AddP(post1)
	if err != nil {
		return
	}
	comment := Rest_Api.Comment{

		Name:  "perviy comment",
		Email: "fucktthemall@mail.com",
		Body:  "my first comment",
	}
	c, err := newService.AddC(p, comment)
	if err != nil {
		return
	}
	fmt.Println(c)
	//service.repository.NewRepository(db).Add(post1)

	//_, err := repos.Add(post1)
	//if err != nil {
	//	log.Fatal("in db no info", err.Error())
	//}

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
