package forum

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/pkg/errno"
	service "tongue/service/forum"
)

func DeleteLike(c *gin.Context) {
	email := c.MustGet("email").(string)
	id := c.Query("postid")
	if err := service.DeleteLike(email, id); err != nil {
		handler.SendError(c, err, errno.ErrDatabase, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, "success")
}
