package category

import (
	"api-go-vue-admin/mysqli"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Category struct {
	Id   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Pid  int    `form:"pid" json:"pid"`
}

//栏目列表
func categoryList(c *gin.Context) {

	Db := mysqli.Connect()

	//实例化一个结构体类型的数组
	var category []Category

	err := Db.Select(&category, "select * from category")

	if err != nil {
		fmt.Println("select * from category 失败", err)
	}

	c.JSON(200, gin.H{
		"categoryList": category,
	})
}
