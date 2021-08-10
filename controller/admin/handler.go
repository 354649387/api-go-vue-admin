package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

type Admin struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func init() {

	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_vue_admin")

	if err != nil {
		fmt.Println("mysql connect failed", err)
	}

	Db = db

}

func adminList(c *gin.Context) {

	//创建一个切片存放一条条的结构体
	var admins []Admin

	rows, _ := Db.Query("select * from admin")

	//遍历
	var admin Admin

	for rows.Next() {

		rows.Scan(&admin.Id, &admin.Username, &admin.Password)

		admins = append(admins, admin)

	}

	c.JSON(200, admins)

}

func adminAdd(c *gin.Context) {

	c.JSON(200, "adminAdd")

}
