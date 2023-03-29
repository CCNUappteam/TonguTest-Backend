package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "tongue/handler"
	"tongue/log"
	"tongue/pkg/errno"
	service "tongue/service/user"
	"tongue/util"
)

// Login ... 登录
// @Summary login api
// @Description
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param object body loginRequest true "login_request"
// @Success 200 {object} loginResponse
// @Router /user/login [post]
func Login(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	log.Info("student login function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	var req loginRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	token, err := service.Login(req.Email, req.Password)
	if err != nil {
		SendError(c, errno.ErrBadRequest, nil, err.Error(), GetLine())
		return
	}

	resp := loginResponse{
		Token: token,
	}

	SendResponse(c, nil, resp)
}
