package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/letanthang/my_framework/db"
	types "github.com/letanthang/my_framework/db/types"
	"github.com/letanthang/validator"
)

func GetStudent(c echo.Context) error {
	var req types.StudentReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: "Bad request"})
	}
	if err := validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: "Bad request"})
	}
	student, _ := db.GetStudent(req)
	return c.JSON(http.StatusOK, student)
}

func SearchStudent(c echo.Context) error {
	var req types.StudentSearchReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: "Bad request"})
	}
	if err := validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: "Bad request"})
	}
	student, _ := db.SearchStudent(req)
	return c.JSON(http.StatusOK, student)
}

func GetAllStudent(c echo.Context) error {
	student, _ := db.GetAllStudent()
	return c.JSON(http.StatusOK, student)
}

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GroupStudent(c echo.Context) error {
	student, _ := db.GroupStudent()
	return c.JSON(http.StatusOK, student)
}
