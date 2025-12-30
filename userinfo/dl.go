package userinfo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录结构体
type Login_info struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 登录函数
func Login(c *gin.Context) {
	var login Login_info
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
