package forum

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/pkg/errno"
	service "tongue/service/forum"
)

func PostComment(c *gin.Context) {
	var req CommentRequest
	email := c.MustGet("email").(string)
	if err := c.ShouldBind(&req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}
	if err := service.PostComment(email, req.PostId, req.Content); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, "success")
}
