package handler

import (
	"Rest-Api/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) AddPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &post)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = h.services.Post.Add(post)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = w.Write(b)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	s := path[len(path)-1]
	i, _ := strconv.Atoi(s)
	get := h.services.Post.Get(i)
	marshal, _ := json.Marshal(get)
	_, err := w.Write(marshal)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func (h *Handler) DelPost(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	s := path[len(path)-1]
	i, _ := strconv.Atoi(s)
	id, err := h.services.Post.Del(i)
	if err != nil {
		log.Fatal(err.Error())
	}
	ii := strconv.Itoa(int(id))
	_, err = w.Write([]byte(r.Method + " id: " + ii))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	b, _ := ioutil.ReadAll(r.Body) // do error checking!
	defer r.Body.Close()

	err := json.Unmarshal(b, &post)
	if err != nil {
		log.Fatal(err.Error())
	}
	update := h.services.Post.Update(post)
	marshal, _ := json.Marshal(update)
	_, err = w.Write(marshal)
	if err != nil {
		log.Fatal(err.Error())
	}

}
