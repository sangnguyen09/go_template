package route

import (
	"github.com/labstack/echo"
	"github.com/letanthang/my_framework/handlers"
)

func Public(e *echo.Echo) {
	publicRoute := e.Group("/v1/public")
	publicRoute.GET("/student", handlers.GetStudent)
}

func Staff(e *echo.Echo) {
	publicRoute := e.Group("/v1/staff")
	publicRoute.POST("/student", handlers.AddStudent)
}
