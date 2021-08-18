package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"net/http"
	"path"
)

func saveImg(c *gin.Context) {
	//获取文件
	f, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, "接收文件失败")
	}
	//f.Filename 原本的文件名称
	//获取文件后缀
	prefix := path.Ext(f.Filename)

	name := uuid.New()

	//指定保存路径  不能用绝对路径，否则传不上去
	dst := path.Join("./upload/images/", name+prefix)

	if err := c.SaveUploadedFile(f, dst); err != nil {
		c.String(http.StatusBadRequest, "保存文件失败")
	}

	c.JSON(200, gin.H{
		"img": "/upload/images/" + name + prefix,
	})

}
