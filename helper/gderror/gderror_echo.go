package gderror

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		code                = http.StatusInternalServerError
		message interface{} = map[string]string{"message": "Internal Server Error"}
	)

	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
		if he.Message != nil {
			message = map[string]interface{}{
				"message": he.Message,
			}
		}
	}

	if !c.Response().Committed {
		if err = c.JSON(code, message); err != nil {
			return
		}
	}
}
