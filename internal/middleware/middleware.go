package middleware

import (
	"fmt"
	"github.com/Anarr/entain/internal/model"
	"github.com/Anarr/entain/internal/util/sliceutil"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

var supportedSourceHeaders = []string{"game", "server", "payment"}

func CheckSourceHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		source := c.Request().Header.Get("Source-Type")

		if strings.TrimSpace(source) == "" {
			return model.APIError{
				StatusCode: http.StatusUnprocessableEntity,
				Err:        "Source-Type header required",
			}
		}

		if !sliceutil.Exists(supportedSourceHeaders, source) {
			return model.APIError{
				StatusCode: http.StatusUnprocessableEntity,
				Err:        fmt.Sprintf("unsupported source-type %s", source),
			}
		}

		return next(c)
	}
}
