package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sangnguyen09/go_template/db/mongo"
	"github.com/sangnguyen09/go_template/router"
)

func main() {
	mongo := &mongo.Mongo{}
	mongo.Connect()

	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, //1KB
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	// if config.Config.Profiler.StatsdAddress != "" {
	// 	e.Use(profiler.ProfilerWithConfig(profiler.ProfilerConfig{Address: config.Config.Profiler.StatsdAddress, Service: config.Config.Profiler.Service}))
	// }

	// e.File("/form", "form.html")
	router.UserRouter(e, mongo)

	fmt.Println("Server listening at 9090")

	port := "9090"
	err := e.Start(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
