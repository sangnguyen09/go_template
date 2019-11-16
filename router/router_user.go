package router

import (
	"github.com/labstack/echo"
	"github.com/sangnguyen09/go_template/db/mongo"
	"github.com/sangnguyen09/go_template/handlers"
	"github.com/sangnguyen09/go_template/repository/repo_impl"
)

func UserRouter(e *echo.Echo, mongo *mongo.Mongo) {
	handler := handlers.UserHandler{
		UserRepo: repo_impl.NewUserRepo(mongo),
	}

	public := e.Group("/api/user/v1/public")
	// public.GET("/health", handler.CheckHealth)
	// public.GET("/users", handler.GetAllStudent)
	public.POST("/login", handler.Login)
	public.POST("/register", handler.Register)

	//Router Private JWT
	// private := e.Group("/api/student/v1/staff")
	// private.POST("/student", handler.AddStudent)
}
