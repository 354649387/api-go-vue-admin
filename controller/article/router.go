package article

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	article := e.Group("/article")

	{
		article.GET("/list", articleList)
		article.GET("/add", add)
		article.GET("/update", update)
		article.GET("/delete", delete)
		article.GET("/getArticleById", getArticleById)
	}

}
