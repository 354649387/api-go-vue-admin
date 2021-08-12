package article

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	article := e.Group("/article")

	{
		article.GET("/list", articleList)
	}

}
