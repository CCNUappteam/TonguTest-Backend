package user

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"

	"tongue/model/record"
	"tongue/pkg/errno"
)

type RecordRequest struct {
	Health string `json:"health"`
}

func PostRecord(c *gin.Context) {
	email := c.MustGet("email").(string)
	var req RecordRequest
	if err := c.ShouldBind(&req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}
	if err := record.CreateRecord(email, req.Health); err != nil {
		handler.SendError(c, errno.ErrUploadFile, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, "Post Health State success!")
}
