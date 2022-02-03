package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) identifiedUser(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		header := ctx.Request().Header.Get(authorizationHeader)
		if header == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "empty auth header")
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid auth header")
		}

		userId, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		ctx.Set(userCtx, userId)
		return handlerFunc(ctx)
	}
}
