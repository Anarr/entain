package cmd

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func getHTTPServer() *echo.Echo {
	return echo.New()
}

func initSwagger(echoServer *echo.Echo) {
	echoServer.GET("/swagger/*", echoSwagger.EchoWrapHandler())
}
