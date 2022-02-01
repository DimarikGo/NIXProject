package handler

import (
	"Rest-Api/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (h *Handler) GetComment(ctx echo.Context) error {
	commentId, _ := strconv.Atoi(ctx.Param("id"))
	get, err := h.services.Comment.Get(commentId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, get)
}

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
