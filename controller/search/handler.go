package search

import (
	"api-go-vue-admin/mysqli"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Article struct {
	Id     int    `json:"id" form:"id"`
	Title  string `json:"title" form:"title"`
	Cid    int    `json:"cid" form:"cid"`
	Aid    int    `json:"aid" form:"aid"`
	Status int    `json:"status" form:"status"`
}

func searchArticle(c *gin.Context) {

	Db := mysqli.Connect()

	var article []Article

	title := c.Query("title")

	err := Db.Select(&article, "select * from article where title=?", title)

	if err != nil {
		fmt.Println("select * from article失败")
	}

	c.JSON(200, gin.H{
		"articleList": article,
	})

}
