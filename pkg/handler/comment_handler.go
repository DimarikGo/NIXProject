package handler

import (
	"Rest-Api/models"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	s := path[len(path)-1]
	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	get := h.services.Comment.Get(i)
	for _, comment := range get {
		bytes, err := xml.Marshal(comment)
		if err != nil {
			log.Fatal(err.Error())
		}
		_, err = w.Write(bytes)
		if err != nil {
			log.Fatal(err.Error())
		}
		marshal, _ := json.Marshal(comment)
		_, err = w.Write(marshal)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err := json.Unmarshal(b, &comment)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = h.services.Comment.Add(comment)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = w.Write(b)
	if err != nil {
		log.Fatal(err.Error())
	}
	bytes, err := xml.Marshal(comment)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = w.Write(bytes)

	if err != nil {
		log.Fatal(err.Error())
	}
}
func (h *Handler) DelComment(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	s := path[len(path)-1]
	i, _ := strconv.Atoi(s)
	id, err := h.services.Comment.Del(i)
	if err != nil {
		log.Fatal(err.Error())
	}
	ii := strconv.Itoa(id)
	_, err = w.Write([]byte(r.Method + " id:" + ii))

	if err != nil {
		log.Fatal(err.Error())
	}
}
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &comment)
	if err != nil {
		log.Fatal(err.Error())
	}

	update := h.services.Comment.Update(comment)
	marshal, err := json.Marshal(update)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = w.Write(marshal)
	if err != nil {
		log.Fatal(err.Error())
	}

	bytes, err := xml.Marshal(comment)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Write(bytes)
}
