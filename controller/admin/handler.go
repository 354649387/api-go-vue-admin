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

//搜索条件结构体
type SearchList struct {
	Username string `json:"username" form:"username"`
	//Category string `json:"category" form:"category"`
	//AdminId string `json:"admin" form:"admin"`
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
	page := c.DefaultQuery("page", "1")
	//每页显示多少
	pageSize := c.DefaultQuery("pageSize", "10")

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

func adminSearch(c *gin.Context) {

	var searchLists SearchList

	//获取到form参数并映射到结构体，结构体后面的标注必须得有form
	if err := c.Bind(&searchLists); err != nil {
		c.JSON(201, "获取请求参数绑定到结构体失败")
	}

	username := searchLists.Username
	//categoryId := searchLists.Category
	//adminId := searchLists.AdminId

	var admin []Admin

	//fmt.Println(username)
	//总条数
	var total int64
	if username == "" {
		err := Db.QueryRow("select count(*) from admin").Scan(&total)

		if err != nil {
			fmt.Println("获取总条数失败")
		}

		err1 := Db.Select(&admin, "select * from admin")

		if err1 != nil {
			fmt.Println("空搜索条件搜索失败")
		}

	}

	err2 := Db.QueryRow("select count(*) from admin where username = ?", username).Scan(&total)

	if err2 != nil {
		fmt.Println("获取总条数失败")
	}

	err := Db.Select(&admin, "select * from admin where username = ?", username)

	if err != nil {
		fmt.Println("非空搜索条件搜索失败")
	}

	c.JSON(200, gin.H{
		"searchResult": admin,
		//用户名不会那么的，直接传1
		"total": total,
	})

}
