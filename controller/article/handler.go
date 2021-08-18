package article

import (
	"api-go-vue-admin/mysqli"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Article struct {
	Id      int    `json:"id" form:"id"`
	Title   string `json:"title" form:"title"`
	Img     string `json:"img" form:"img"`
	Content string `json:"content" form:"content"`
	Cid     int    `json:"cid" form:"cid"`
	Aid     int    `json:"aid" form:"aid"`
	Status  int    `json:"status" form:"status"`
}

//文章新增
func add(c *gin.Context) {

	title := c.Query("title")
	img := c.Query("img")
	content := c.Query("content")

	article := Article{Title: title, Img: img, Content: content}

	db := mysqli.GormConnect()
	tx := db.Table("article")
	res := tx.Create(&article)

	if res.RowsAffected <= 0 {

		c.JSON(201, res.Error)

	}

	c.JSON(200, "success")

}

//文章更新
func update(c *gin.Context) {

}

//文章列表
func articleList(c *gin.Context) {

	db := mysqli.GormConnect()

	var article []Article

	//第几页
	page := c.DefaultQuery("page", "1")
	//每页显示多少
	pageSize := c.DefaultQuery("pageSize", "10")

	//字符串转成数字类型
	page1, _ := strconv.Atoi(page)
	pageSize1, _ := strconv.Atoi(pageSize)

	offsetNum := (page1 - 1) * pageSize1

	//limit：多少个，offset从哪个开始
	tx := db.Table("article")

	tx.Limit(pageSize1).Offset(offsetNum)

	tx.Find(&article)

	//article表总记录数
	tx1 := db.Table("article")

	var count int64

	tx1.Count(&count)

	c.JSON(200, gin.H{
		"articleList": article,
		"total":       count,
	})

}

func articleList1(c *gin.Context) {

	Db := mysqli.Connect()

	//第几页
	page := c.DefaultQuery("page", "1")
	//每页显示多少
	pageSize := c.DefaultQuery("pageSize", "10")

	//字符串转成数字类型
	page1, _ := strconv.Atoi(page)
	pageSize1, _ := strconv.Atoi(pageSize)

	offsetNum := (page1 - 1) * pageSize1

	var total int64

	err := Db.QueryRow("select count(*) from article").Scan(&total)

	if err != nil {
		fmt.Println("select count(*) from article失败", err)
	}

	rows, _ := Db.Query("select * from article limit ?,?", offsetNum, pageSize1)

	var articles []Article

	var article Article

	for rows.Next() {

		err := rows.Scan(&article.Id, &article.Title, &article.Cid, &article.Aid, &article.Status)

		if err != nil {
			fmt.Println("rows.Next()->rows.Scan失败", err)
		}

		articles = append(articles, article)

	}

	c.JSON(200, gin.H{
		"articleList": articles,
		"total":       total,
	})

}
