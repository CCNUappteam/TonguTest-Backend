package forum

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/pkg/errno"
	service "tongue/service/forum"
)

func GetComment(c *gin.Context) {
	id := c.Query("postid")
	comments, err := service.GetComments(id)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, gin.H{
		"comments": comments,
	})
}
