package main

import (
	"echo-sample/app"
	"echo-sample/db/redis"
	"echo-sample/handler"
)

func main() {
	app.Init()
	redis.Init()

	app.Server.GET("/", handler.HelloWorldRequest)

	app.Server.Logger.Fatal(app.Server.Start("127.0.0.1:1323"))
}
