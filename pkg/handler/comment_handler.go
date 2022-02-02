package handler

import (
	"Rest-Api/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

// GetComment godoc
// @Summary Get comment from DB.
// @Description Get comment from DB table comment.
// @Tags GetComment
// @Accept json
// @Produce json
// @Produce xml
// @Param  id path int true "Id"
// @Success 200 {object} map[string]interface{}
// @Router /comment/{id} [get]
func (h *Handler) GetComment(ctx echo.Context) error {
	commentId, _ := strconv.Atoi(ctx.Param("id"))
	get, err := h.services.Comment.Get(commentId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, get)
}

// AddComment godoc
// @Summary Add comments in DB.
// @Description Add comments in DB table comments.
// @Tags AddComment
// @Accept json
// @Produce json
// @Produce xml
// @Param comment body models.Comment true "List Comment"
// @Success 200 {object} map[string]interface{}
// @Router /comment/add [post]
func (h *Handler) AddComment(ctx echo.Context) error {
	var comment models.Comment
	err := ctx.Bind(&comment)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode comment data"))
	}
	addComment, err := h.services.Comment.Add(&comment)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create comment"))
	}
	return ctx.JSON(http.StatusCreated, addComment)
}

// DelComment godoc
// @Summary Delete comment from DB.
// @Description Delete comment from DB table comment.
// @Tags DelComment
// @Accept json
// @Produce json
// @Produce xml
// @Param  id path int true "Id"
// @Success 200 {object} map[string]interface{}
// @Router /comment/del/{id} [delete]
func (h *Handler) DelComment(ctx echo.Context) error {
	commentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse comment"))
	}

	delId, err := h.services.Comment.Del(commentId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, delId)
}

// UpdateComment godoc
// @Summary Update comment from DB.
// @Description Update comment from DB table comment.
// @Tags UpdateComment
// @Accept json
// @Produce json
// @Produce xml
// @Param comment body models.Comment true "Update List Comment"
// @Success 200 {object} map[string]interface{}
// @Router /comment/patch/{id} [patch]
func (h *Handler) UpdateComment(ctx echo.Context) error {
	var comment models.Comment
	err := ctx.Bind(&comment)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode comment data"))
	}
	commentId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse comment"))
	}

	update := h.services.Comment.Update(&comment, commentId)

	return ctx.JSON(http.StatusOK, update)
}
