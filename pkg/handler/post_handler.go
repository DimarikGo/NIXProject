package handler

import (
	"Rest-Api/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (h *Handler) AddPost(ctx echo.Context) error {
	var post models.Post
	err := ctx.Bind(&post)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode post data"))
	}
	addPost, err := h.services.Post.Add(&post)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create post"))
	}
	return ctx.JSON(http.StatusCreated, addPost)
}

func (h *Handler) GetPost(ctx echo.Context) error {
	param, _ := strconv.Atoi(ctx.Param("id"))
	get, err := h.services.Post.Get(param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, get)
}

func (h *Handler) DelPost(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse post"))
	}

	delId, err := h.services.Post.Del(userID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, delId)
}

func (h *Handler) UpdatePost(ctx echo.Context) error {
	var post models.Post
	err := ctx.Bind(&post)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode post data"))
	}
	postId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse post"))
	}

	update := h.services.Post.Update(&post, postId)

	return ctx.JSON(http.StatusOK, update)
}
