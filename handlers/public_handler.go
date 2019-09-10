package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/letanthang/my_framework/db"
)

func GetStudent(c echo.Context) error {
	student, _ := db.GetStudent()
	return c.JSON(http.StatusOK, student)
}

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
