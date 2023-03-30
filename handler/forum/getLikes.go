package forum

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/pkg/errno"
	service "tongue/service/forum"
)

func GetLikes(c *gin.Context) {
	id := c.Query("postid")
	count, err := service.GetLikes(id)
	if err != nil {
		handler.SendError(c, err, errno.ErrDatabase, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, gin.H{
		"count": count,
	})
}
