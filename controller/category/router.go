package category

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	category := e.Group("/category")

	{
		//栏目列表
		category.GET("/list", categoryList)
	}

}
