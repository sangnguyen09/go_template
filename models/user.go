package models

type User struct {
	UserId   int    `json:"user_id" bson:"user_id" gorm:"column:user_id"`
	Username string `json:"username,omitempty" bson:"username,omitempty"  gorm:"column:username" valid:"required~Username không được để trống,stringlength(5|15)~Username không được chứa kí tự đặc biệt và giới hạn từ 5 đến 15 kí tự"`
	Role     Role `json:"role,omitempty" bson:"role,omitempty"  gorm:"column:role" `
	Password string `json:"password,omitempty" bson:"password,omitempty"  gorm:"column:password" valid:"required~Password không được để trống"`
	Avatar   string `json:"avatar,omitempty" bson:"avatar,omitempty"  gorm:"column:avatar"`
}
type LoginRequest struct {
	Username string `json:"username" gorm:"column:username" valid:"required~Tài khoản hoặc mật khẩu không chính xác,stringlength(5|15)~Tài khoản hoặc mật khẩu không chính xác" `
	Password string `json:"password" gorm:"column:password" valid:"required~Tài khoản hoặc mật khẩu không chính xác"`
}
type ChangePassword struct {
	PasswordCurrent string `json:"password_current" valid:"required~Vui lòng nhập mật khẩu hiện tại"`
	PasswordNew string `json:"password_new" bson:"password" gorm:"column:password" valid:"required~Vui lòng nhập mật khẩu mới"`
}
type UserResponse struct {
	Username     string `json:"username"`
	UserId       int `json:"user_id"`
	Role         Role `json:"role"`
	Avatar       string `json:"avatar"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
type Delete struct {
	UserId int `json:"user_id" bson:"user_id" gorm:"column:user_id" valid:"required~Bạn là hacker ?,int~Bạn là hacker ?"`
}