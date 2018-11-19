package main

import (
	"data-service/app"
	"data-service/db/redis"
	"data-service/handler"
)

func main() {
	app.Init()
	redis.Init()

	app.Server.GET("/", handler.HelloWorldRequest)

	app.Server.Logger.Fatal(app.Server.Start("127.0.0.1:1323"))
}
