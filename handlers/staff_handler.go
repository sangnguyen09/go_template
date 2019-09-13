package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/letanthang/my_framework/db"
	types "github.com/letanthang/my_framework/db/types"
	"github.com/letanthang/validator"
)

func AddStudent(c echo.Context) error {
	var req types.StudentAddReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "Bad request", Message: "Bad parameter"})
	}

	if err := validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "Bad request", Message: err.Error()})
	}

	i, err := db.InsertStudent(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "bad request", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, i)
}

func DeleteStudent(c echo.Context) error {
	var req types.DeleteReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "Bad request", Message: "Bad parameter"})
	}

	if err := validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "Bad request", Message: err.Error()})
	}

	res, err := db.DeleteStudent(req.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "bad request", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
