package upload

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	upload := e.Group("/upload")

	{
		upload.POST("/saveImg", saveImg)
	}

}
