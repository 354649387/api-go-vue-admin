package login

import (
	"api-go-vue-admin/mysqli"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func Login(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	db := mysqli.GormConnect()

	tx := db.Table("admin")

	var admin Admin

	tx.Where("username = ?", username).Find(&admin)

	if tx.RowsAffected <= 0 {
		fmt.Println("用户名不存在")
		c.String(201, "用户名不存在")
		return
	}

	if password != admin.Password {
		fmt.Println("密码错误")
		c.String(201, "密码错误")
		return
	}

	c.JSON(200, gin.H{
		"username": username,
	})

}
