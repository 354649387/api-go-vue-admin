package search

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	article := e.Group("/search")

	{
		article.GET("/article", searchArticle)
		article.GET("/admin", searchAdmin)
	}

}
