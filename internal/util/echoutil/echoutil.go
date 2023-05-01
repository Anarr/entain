package echoutil

import (
	"github.com/Anarr/entain/internal/model"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	if ae, ok := err.(model.APIError); ok {
		c.JSON(ae.StatusCode, ae)
		return
	}
}
