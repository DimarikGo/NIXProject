package handler

import (
	"Rest-Api/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

// SignUp godoc
// @Summary Add user in DB.
// @Description Add user in DB table user.
// @Tags SignUp
// @Accept json
// @Produce json
// @Produce xml
// @Param user body models.User true "List User"
// @Success 200 {object} map[string]interface{}
// @Router /auth/Sign-up [post]
func (h *Handler) SignUp(ctx echo.Context) error {
	var user models.User

	err := ctx.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode user data"))
	}
	id, err := h.services.Authorization.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create user"))
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type SignInUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignIn godoc
// @Summary  Authorization user in Api.
// @Description  Authorization user in Rest-Api.
// @Tags SignIn
// @Accept json
// @Produce json
// @Produce xml
// @Param user body SignInUser true "Auth user"
// @Success 200 {object} map[string]interface{}
// @Router /auth/Sign-in [post]
func (h *Handler) SignIn(ctx echo.Context) error {
	var user SignInUser
	err := ctx.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode user data"))
	}
	token, _ := h.services.Authorization.GenerateToken(user.Username, user.Password)

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}
