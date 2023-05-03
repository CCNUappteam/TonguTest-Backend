package user

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/model/record"
	"tongue/pkg/errno"
)

func GetRecord(c *gin.Context) {
	email := c.MustGet("email").(string)
	rec, err := record.GetRecord(email)
	if err != nil {
		handler.SendError(c, errno.ErrUploadFile, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, gin.H{
		"Record": rec,
	})
}
