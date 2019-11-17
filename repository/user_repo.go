package repository

import (
	"context"

	"github.com/sangnguyen09/go_template/models"
)

type UserRespository interface {
	CheckLogin(context context.Context, loginReq models.LoginRequest) (models.User, error)
	Register(context context.Context, user models.User) error
	CheckExist(context context.Context, username string) (bool)
	ComparePassword(context context.Context, pwdcurrent string, userId int) (bool)
	UpdatePass(context context.Context,pwdnew string, userId int)  error
	//VerifyAccount(context context.Context, registerReq *models.User)(int64, error)
}
