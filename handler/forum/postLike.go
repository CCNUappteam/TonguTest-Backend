package forum

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/pkg/errno"
	service "tongue/service/forum"
)

func PostLike(c *gin.Context) {
	id := c.Query("postid")
	email := c.MustGet("email").(string)
	if err := service.Like(email, id); err != nil {
		handler.SendError(c, err, errno.ErrDatabase, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, "success")
}
