package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Server *echo.Echo

// init echo web server
func Init() {
	Server = echo.New()
	// Middleware
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())

}
