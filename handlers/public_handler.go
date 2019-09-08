package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetStudent(c echo.Context) error {
	data := map[string]interface{}{
		"class_name": "Golang course",
		"students":   []string{"A", "B", "C"},
	}
	return c.JSON(http.StatusOK, data)
}
