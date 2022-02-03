package handler

import (
	"Rest-Api/pkg/services"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *services.Service
	echo     *echo.Echo
}

func NewHandler(services *services.Service, echo *echo.Echo) *Handler {
	return &Handler{
		services: services,
		echo:     echo}
}

//
//func (h *Handler) InitRoutes() *http.ServeMux {
//
//	serveMux := http.NewServeMux()
//
//	serveMux.HandleFunc("/post/add", h.AddPost)
//	serveMux.HandleFunc("/post/", h.GetPost)
//	serveMux.HandleFunc("/post/del/", h.DelPost)
//	serveMux.HandleFunc("/post/update/", h.UpdatePost)
//
//	serveMux.HandleFunc("/comments/add", h.AddComment)
//	serveMux.HandleFunc("/comments/", h.GetComment)
//	serveMux.HandleFunc("/comments/del/", h.DelComment)
//	serveMux.HandleFunc("/comments/update/", h.UpdateComment)
//
//	return serveMux
//}

func (h *Handler) InitRoutes() *echo.Echo {

	auth := h.echo.Group("/auth")

	auth.POST("/Sign-up", h.SignUp)
	auth.POST("/Sign-in", h.SignIn)

	post := h.echo.Group("/post", h.identifiedUser)
	post.Use(h.identifiedUser)
	post.POST("/add", h.AddPost)
	post.PATCH("/patch/:id", h.UpdatePost)
	post.DELETE("/del/:id", h.DelPost)
	post.GET("/:id", h.GetPost)

	comment := h.echo.Group("/comment", h.identifiedUser)
	comment.Use(h.identifiedUser)
	comment.POST("/add", h.AddComment)
	comment.PATCH("/update/:id", h.UpdateComment)
	comment.DELETE("/del/:id", h.DelComment)
	comment.GET("/:id", h.GetComment)

	return h.echo
}
