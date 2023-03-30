package forum

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/pkg/errno"
	service "tongue/service/forum"
)

func PublishPost(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var req PostRequest
	if err := c.ShouldBind(&req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}
	email := c.MustGet("email").(string)
	id, err := service.UploadPost(email, req.Title, req.Content)
	if err != nil {
		handler.SendError(c, errno.ErrServerWrong, nil, err.Error(), handler.GetLine())
		return
	}

	handler.SendResponse(c, nil, gin.H{
		"result":  "Success",
		"post_id": id,
	})

}
