package userinfo

import (
	"errors"

	"aws_rds/usermod"

	"gorm.io/gorm"
)

// 用户注册请求结构体
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone"`
}

// 用户注册函数
func (r *RegisterRequest) Register(db *gorm.DB) error {

	// 检查用户名是否存在
	var user usermod.User
	result := db.Where("username = ?", r.Username).First(&user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		return errors.New("用户名已存在")
	}

	var email usermod.User
	resultmail := db.Where("email = ?", r.Email).First(&email)
	if resultmail.Error != nil && !errors.Is(resultmail.Error, gorm.ErrRecordNotFound) {
		return resultmail.Error
	}

	if resultmail.RowsAffected > 0 {
		return errors.New("邮箱已被注册")
	}

	user = usermod.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
		Phone:    r.Phone,
	}
	return db.Create(&user).Error
}
