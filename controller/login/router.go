package login

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	login := e.Group("/login")

	{
		login.GET("/login", Login)
	}

}
