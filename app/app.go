package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var Server *echo.Echo

// init echo web server
func Init() {
	Server = echo.New()

	Server.Logger.SetLevel(log.INFO)
	// Middleware
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())

}
