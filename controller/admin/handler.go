package admin

import (
	"api-go-vue-admin/mysqli"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"strconv"
)

var Db *sqlx.DB

type Admin struct {
	Id        int            `json:"id" form:"id"`
	Username  string         `json:"username" form:"username"`
	Password  string         `json:"password" form:"password"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" form:"deleted_at" gorm:"column:deleted_at"`
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

func (a Admin) TableName() string {
	return "admin"
}

func adminList(c *gin.Context) {

	db := mysqli.GormConnect()

	var admin []Admin

	//第几页
	page := c.DefaultQuery("page", "1")
	//每页显示多少
	pageSize := c.DefaultQuery("pageSize", "10")

	//字符串转成数字类型
	page1, _ := strconv.Atoi(page)
	pageSize1, _ := strconv.Atoi(pageSize)

	offsetNum := (page1 - 1) * pageSize1

	tx := db.Offset(offsetNum).Limit(pageSize1).Order("id DESC").Find(&admin)

	var total int64 = 0

	db.Model(Admin{}).Count(&total)

	//total := tx.RowsAffected

	if tx.Error != nil {
		fmt.Println("查询所有文章列表失败", tx.Error)
	}

	c.JSON(200, gin.H{
		"adminList": admin,
		"total":     total,
	})

}

func adminAdd(c *gin.Context) {

	//原生写法，暂未改成gorm写法

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

func adminDelete(c *gin.Context) {

	id := c.Query("id")
	fmt.Println(id)
	admin := Admin{}

	db1 := mysqli.GormConnect()

	db1.Where("id = ?", id).Take(&admin)

	db1.Delete(&admin)

	c.JSON(200, "success")

}

//通过id获取管理员详情
func getAdminById(c *gin.Context) {

	id := c.Query("id")

	db := mysqli.GormConnect()

	admin := Admin{}

	res := db.Where("id = ?", id).Take(&admin)

	if res.Error != nil {

		fmt.Println(res.Error)

	}

	c.JSON(200, gin.H{
		"adminList": admin,
	})

}

//管理员更新
func update(c *gin.Context) {

	id := c.Query("id")
	username := c.Query("username")
	password := c.Query("password")

	//通过结构体变量设置更新字段
	admin := Admin{
		Username: username,
		Password: password,
	}

	db := mysqli.GormConnect()

	re := db.Model(&Admin{}).Where("id = ?", id).Updates(&admin)

	if re.Error != nil {

		fmt.Println(re.Error)

	}

	c.JSON(200, "success")
}
