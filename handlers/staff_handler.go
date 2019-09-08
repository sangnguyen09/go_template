package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/letanthang/my_framework/db"
	"github.com/letanthang/my_framework/db/types"
	"github.com/letanthang/validator"
)

func AddStudent(c echo.Context) error {

	var req types.StudentAddReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "Bad request", Message: "bad request"})
	}

	if err := validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "Bad request", Message: err.Error()})
	}

	id, err := db.InsertStudent(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "Bad request", Message: err.Error()})
	}

	data := map[string]interface{}{
		"student_id": id,
	}
	return c.JSON(http.StatusOK, data)
}
