package handler

import (
	"Rest-Api/pkg/services"
	"net/http"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/post/add", h.AddPost)
	serveMux.HandleFunc("/post/", h.GetPost)
	serveMux.HandleFunc("/post/del/", h.DelPost)
	serveMux.HandleFunc("/post/update/", h.UpdatePost)

	serveMux.HandleFunc("/comments/add", h.AddComment)
	serveMux.HandleFunc("/comments/", h.GetComment)
	serveMux.HandleFunc("/comments/del/", h.DelComment)
	serveMux.HandleFunc("/comments/update/", h.UpdateComment)

	return serveMux
}
