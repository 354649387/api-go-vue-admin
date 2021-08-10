package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
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

	//第几页
	page := c.Query("page")
	//每页显示多少
	pageSize := c.Query("pageSize")

	//字符串转成数字类型
	page1, _ := strconv.Atoi(page)
	pageSize1, _ := strconv.Atoi(pageSize)

	offsetNum := (page1 - 1) * pageSize1

	//总条数
	var total int64

	err := Db.QueryRow("select count(*) from admin").Scan(&total)

	if err != nil {
		fmt.Println("获取总条数失败")
	}

	rows, _ := Db.Query("select * from admin limit ?,?", offsetNum, pageSize1)

	//遍历
	var admin Admin

	for rows.Next() {

		rows.Scan(&admin.Id, &admin.Username, &admin.Password)

		admins = append(admins, admin)

	}

	c.JSON(200, gin.H{"admins": admins, "total": total})

}

func adminAdd(c *gin.Context) {

	username := c.Query("username")

	password := c.Query("password")

	res, err := Db.Exec("insert into admin(username,password) values(?,?)", username, password)

	if err != nil {
		fmt.Println("insert into admin failed", err)
	}

	rows, err := res.RowsAffected()

	if rows > 0 {

		c.JSON(200, "success")

	} else {

		c.JSON(201, err)

	}

}
