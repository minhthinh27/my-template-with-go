package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		apiKeyHeader := ctx.Request().Header.Get("api-key")
		if apiKeyHeader == "" {
			return ctx.JSON(http.StatusUnauthorized, "api-key empty")
		}

		value := "1234567"
		if apiKeyHeader != value {
			return ctx.JSON(http.StatusUnauthorized, "api-key invalid")
		}

		ctx.Set("CustomerID", 1)
		return next(ctx)
	}
}
