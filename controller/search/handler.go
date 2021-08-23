package search

import (
	"api-go-vue-admin/mysqli"
	"github.com/gin-gonic/gin"
)

type Article struct {
	Id     int    `json:"id" form:"id"`
	Title  string `json:"title" form:"title"`
	Cid    int    `json:"cid" form:"cid"`
	Aid    int    `json:"aid" form:"aid"`
	Status int    `json:"status" form:"status"`
}

type Admin struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func searchArticle(c *gin.Context) {

	//获取查询条件
	title := c.Query("title")
	cid := c.Query("cid")
	aid := c.Query("aid")

	//连接数据库
	db := mysqli.GormConnect()

	var article []Article

	tx := db.Table("article")

	//判断查询条件
	if title != "" {
		tx.Where("title = ?", title)
	}
	if cid != "" {
		tx.Where("cid = ?", cid)
	}
	if aid != "" {
		tx.Where("aid = ?", aid)
	}

	//存入链式查询结果
	tx.Find(&article)

	//结果条数
	rows := tx.RowsAffected

	c.JSON(200, gin.H{
		"searchList": article,
		"total":      rows,
	})

}

func searchAdmin(c *gin.Context) {

	//获取查询条件
	username := c.Query("username")

	//连接数据库
	db := mysqli.GormConnect()

	var admin []Admin

	tx := db.Table("admin")

	//判断查询条件
	if username != "" {
		tx.Where("username = ?", username)
	}

	//存入链式查询结果
	tx.Find(&admin)

	//结果条数
	rows := tx.RowsAffected

	c.JSON(200, gin.H{
		"searchList": admin,
		"total":      rows,
	})

}
