package forum

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tongue/handler"
	"tongue/pkg/errno"
	service "tongue/service/forum"
)

func GetPosts(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var err error
	var limit, page int
	limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}
	post, count, err := service.GetPosts("", limit*page, limit)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, gin.H{
		"posts": post,
		"count": count,
	})
}

func GetMyPosts(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var err error
	var limit, page int
	email := c.MustGet("email").(string)
	limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}
	post, count, err := service.GetPosts(email, limit*page, limit)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, gin.H{
		"posts": post,
		"count": count,
	})
}
