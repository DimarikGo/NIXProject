package handler

import (
	"Rest-Api/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

// AddPost godoc
// @Summary Add post in DB.
// @Description Add post in DB table post.
// @Tags AddPost
// @Accept json
// @Produce json
// @Produce xml
// @Param post body models.Post true "List Post"
// @Success 200 {object} map[string]interface{}
// @Router /post/add [post]
func (h *Handler) AddPost(ctx echo.Context) error {
	id := ctx.Get(userCtx)
	var post models.Post
	err := ctx.Bind(&post)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode post data"))
	}
	addPost, err := h.services.Post.Add(&post, id.(int))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create post"))
	}
	return ctx.JSON(http.StatusCreated, addPost)
}

// GetPost godoc
// @Summary Get post from DB.
// @Description Get post from DB table post.
// @Tags GetPost
// @Accept json
// @Produce json
// @Produce xml
// @Param  id path int true "Id"
// @Success 200 {object} map[string]interface{}
// @Router /post/{id} [get]
func (h *Handler) GetPost(ctx echo.Context) error {
	param, _ := strconv.Atoi(ctx.Param("id"))
	get, err := h.services.Post.Get(param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, get)
}

// DelPost godoc
// @Summary Delete post from DB.
// @Description Delete post from DB table post.
// @Tags DelPost
// @Accept json
// @Produce json
// @Produce xml
// @Param  id path int true "Id"
// @Success 200 {object} map[string]interface{}
// @Router /post/del/{id} [delete]
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

// UpdatePost godoc
// @Summary Update post from DB.
// @Description Update post from DB table post.
// @Tags UpdatePost
// @Accept json
// @Produce json
// @Produce xml
// @Param post body models.Post true "Update List Post"
// @Success 200 {object} map[string]interface{}
// @Router /post/patch/{id} [patch]
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
