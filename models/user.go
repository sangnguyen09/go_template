package models

type User struct {
	UserId   int    `json:"user_id" bson:"user_id" gorm:"column:user_id"`
	Username string `json:"username,omitempty" bson:"username,omitempty"  gorm:"column:username" valid:"required~Username không được để trống,stringlength(5|15)~Username không được chứa kí tự đặc biệt và giới hạn từ 5 đến 15 kí tự"`
	Role     string `json:"role,omitempty" bson:"role,omitempty"  gorm:"column:role" `
	Password string `json:"password,omitempty" bson:"password,omitempty"  gorm:"column:password" valid:"required~Password không được để trống"`
	Avatar   string `json:"avatar,omitempty" bson:"avatar,omitempty"  gorm:"column:avatar"`
}
type LoginRequest struct {
	Username string `json:"username" gorm:"column:password" `
	Password string `json:"password" gorm:"column:username" `
}

type UserResponse struct {
	Username     string `json:"username"`
	UserId       int `json:"user_id"`
	Role         string `json:"role"`
	Avatar       string `json:"avatar"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}