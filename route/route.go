package route

import (
	"github.com/labstack/echo"
	"github.com/letanthang/my_framework/handlers"
)

func Public(e *echo.Echo) {
	publicRoute := e.Group("/v1/public")
	publicRoute.GET("/student", handlers.GetAllStudent)
	publicRoute.PATCH("/student/simple", handlers.GetStudent)
	publicRoute.PATCH("/student", handlers.SearchStudent)
	publicRoute.GET("/health", handlers.CheckHealth)
	publicRoute.GET("/student/group/last_name", handlers.GroupStudent)
}

func Staff(e *echo.Echo) {
	staffRoute := e.Group("/v1/staff")
	staffRoute.POST("/student", handlers.AddStudent)
	staffRoute.DELETE("/student", handlers.DeleteStudent)
}
