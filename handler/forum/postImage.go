package forum

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"tongue/handler"
	"tongue/log"
	"tongue/pkg/errno"
	service2 "tongue/service"
	service "tongue/service/forum"
	"tongue/util"
)

func PostImage(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))
	file, err := c.FormFile("image")
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetFile, nil, err.Error(), handler.GetLine())
		return
	}
	id, _ := strconv.Atoi(c.PostForm("post_id"))
	err, url := service2.UploadFile(file)
	if err != nil || url == "" {
		handler.SendError(c, errno.ErrUploadFile, nil, err.Error(), handler.GetLine())
		return
	}
	err = service.UploadPostImage(id, url)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, "Upload avatar success!")
}
