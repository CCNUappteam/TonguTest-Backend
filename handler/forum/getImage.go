package forum

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/pkg/errno"
	service "tongue/service/forum"
)

func GetImage(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	postid := c.Query("postid")
	images, err := service.GetPostImage(postid)
	if err != nil {
		handler.SendError(c, errno.ErrGetFile, nil, err.Error(), handler.GetLine())
		return
	}

	handler.SendResponse(c, nil, gin.H{
		"images": images,
	})
}
