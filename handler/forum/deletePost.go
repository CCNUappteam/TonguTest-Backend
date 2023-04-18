package forum

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/pkg/errno"
	service "tongue/service/forum"
)

func DeletePost(c *gin.Context) {
	id := c.Query("postid")
	email := c.MustGet("email").(string)
	if err := service.DeletePost(email, id); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error(), handler.GetLine())
	}
	handler.SendResponse(c, nil, "success")
}
