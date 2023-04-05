package test

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/pkg/errno"
	service "tongue/service/test"
)

func TongueTest(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}
	resp, err := service.UploadTest(file)
	if err != nil {
		handler.SendError(c, err, errno.ErrUploadFile, err.Error(), handler.GetLine())
		return
	}

	handler.SendResponse(c, nil, gin.H{
		"result": resp,
	})
}
