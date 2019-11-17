package handlers

import (
	"context"
	"fmt"
	"github.com/sangnguyen09/go_template/lang"
	"github.com/sangnguyen09/go_template/validator"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"

	"github.com/sangnguyen09/go_template/helpers"
	"github.com/sangnguyen09/go_template/middleware"
	"github.com/sangnguyen09/go_template/models"
	"github.com/sangnguyen09/go_template/repository"

)

type UserHandler struct {
	UserRepo repository.UserRespository
}

//----------- Ham xử lý Đăng ký
func (u *UserHandler) Register(c echo.Context) error {
	req := models.User{}

	defer c.Request().Body.Close()

	req.Avatar = "https://lh3.googleusercontent.com/KDOZi-KREdJv1SGLcDuAwrzik8BT5gQWaQCcNp-V6J336HgqAdIq5FuWllkMekap12PoSelf9qmP6cdqB_pLY2g075fNTfUCypGN_L2Vi1kuJK9AD4Uc8J85AHxujzV3-3S-prUTd4_8dhIFFD8hlMmQYucnVf70gPlCk7GOSuKY7sssIYWfy1JHHnpDXFsAV_28eQElkDounDDW8ZWs1bg46MJkCFR6rZdxwnaP55-cZd1qeqlIXDpYqjCcFjMyATI-tv2H3h22OdB8dhZAM83nGMOZSw6Ionf39EEId6lnuYbvGKWYRmnGnRgmvTFwpdd4Sbj_KdDzBp-rHMuKkgfh7zKA6uCnpj1XP82Y6QzdMrjlODcRpn_nyJkLSb0tffq8GB69y4LASiGvmxExPHuMOm6sC7uSKETdG-IUvzNcTA88PrRbFHSv_vgbb_nIf7WFV8aUL-j4StnSpZOlE-EYgNP1EZbL-8pMjR4rA9GeeY881ZhwED4II5smouwVb4IRck9sj4P3E8gcc1daZjW8SNNMJp78AL0K0b_IIPGu3Dc0xRCHSx9hPbtBn9b4Icx4H6xnEcgKb1yKYxUGIzC1L7nM1qv1kRtdNnWwCWlW_J0AVtpU2OyubcjXnncByzTstKZZRh49Uudhr9zz1CXs0WTCcOVaHHFDh3s5Zr3rXiEi9nb3Yg=s200-no"

	if err := c.Bind(&req); err != nil {
		return helpers.ResponseErr(c, http.StatusBadRequest)
	}

	//--- Validate thông tin req ----
	if _, err := govalidator.ValidateStruct(req); err != nil {
		return helpers.ResponseErr(c, http.StatusBadRequest, err.Error())
	}
	if validPass := validator.ValidPassword(req.Password); validPass == false{
		return helpers.ResponseErr(c, http.StatusBadRequest,lang.Password_incorect)
	}
	if validUsername := validator.MatchRegex(req.Username,`^[a-z0-9_]{5,15}$`); validUsername == false{
		return helpers.ResponseErr(c, http.StatusBadRequest,lang.Username_incorect)
	}
	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)

	// ---- Kiểm tra tồn tại tài khoản -----
	if  checkExist :=u.UserRepo.CheckExist(ctx,req.Username) ; checkExist == true{
		return helpers.ResponseErr(c,http.StatusBadRequest,lang.User_exist)
	}

	//--- Mã hoá mật khẩu
	req.Password = helpers.EncryptPass(req.Password)
	req.Role = "member"

	//
	 err := u.UserRepo.Register(ctx, req)
	if err !=nil {
		return helpers.ResponseErr(c, http.StatusBadRequest,err.Error())
	}
	return helpers.ResponseData(c,nil)

}

//----------- Hàm xử lý đăng nhập
//response trả về kèm token để truy cập các api về sau

func (u *UserHandler) Login(c echo.Context) error {
	// Lấy thông tin dữ liệu từ người dùng gửi lên
	req := models.LoginRequest{}

	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return helpers.ResponseErr(c, http.StatusBadRequest)
	}
	//----- ket thuc------

	//--- Validate thông tin req ----
	if _, err := govalidator.ValidateStruct(req); err != nil {
		return helpers.ResponseErr(c, http.StatusBadRequest, err.Error())
	}
	if validPass := validator.ValidPassword(req.Password); validPass == false{
		return helpers.ResponseErr(c, http.StatusBadRequest,lang.Password_incorect)
	}
	if validUsername := validator.MatchRegex(req.Username,`^[a-z0-9_]{5,15}$`); validUsername == false{
		return helpers.ResponseErr(c, http.StatusBadRequest,lang.Username_incorect)
	}
	//--- convert pass to md5 ----/
	req.Password = helpers.EncryptPass(req.Password)

	//------ Check Database----/
	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	user, err := u.UserRepo.CheckLogin(ctx, req)

	if err != nil {
		return helpers.ResponseErr(c, http.StatusUnauthorized, lang.Login_fail)
	}
	//------ Tao ma token va refresh token-----
	token, err := middleware.GenToken(user)

	if err != nil {
		return helpers.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}
	refreshToken, err := middleware.GenTokenRefresh(user)
	if err != nil {
		return helpers.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.ResponseData(c, models.UserResponse{user.Username, user.UserId, user.Role, user.Avatar, token, refreshToken})

}

func (u *UserHandler) ChangePassword(c echo.Context) error {
	defer c.Request().Body.Close()

	//---- Lấy thông thông tin user từ token
		user :=  c.Get("user").(*jwt.Token)
		claims := user.Claims.(*models.JWTCustomClaims)

		req := models.ChangePassword{}
		if err := c.Bind(&req); err != nil {
			return helpers.ResponseErr(c, http.StatusBadRequest)
		}
		//--- Validate thông tin req ----
		if _, err := govalidator.ValidateStruct(req); err != nil {
			fmt.Println( err.Error())
			return helpers.ResponseErr(c, http.StatusBadRequest, err.Error())
		}
		if validPassCurrent := validator.ValidPassword(req.PasswordCurrent); validPassCurrent == false{
			return helpers.ResponseErr(c, http.StatusBadRequest,lang.Password_incorect)
		}
		if validPassNew := validator.ValidPassword(req.PasswordNew); validPassNew == false{
			return helpers.ResponseErr(c, http.StatusBadRequest,lang.Password_incorect)
		}

	ctx,_ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	//--- convert pass to md5 ----/
		req.PasswordCurrent = helpers.EncryptPass(req.PasswordCurrent)
		req.PasswordNew = helpers.EncryptPass(req.PasswordNew)

	//----- Check pass cũ -----
	if checkPass := u.UserRepo.ComparePassword(ctx, req.PasswordCurrent, claims.UserId); checkPass == false{
		return helpers.ResponseErr(c,http.StatusBadRequest,lang.Password_compare) // không trùng khớp
	}

	//---- Update Pass to DB ----
	err :=	u.UserRepo.UpdatePass(ctx,req.PasswordNew,claims.UserId)
	if err != nil {
		  	return helpers.ResponseErr(c,http.StatusInternalServerError,err.Error())
	}

	return helpers.ResponseData(c,nil)

}