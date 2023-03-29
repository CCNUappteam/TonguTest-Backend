package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "tongue/handler"
	"tongue/log"
	"tongue/model/user"
	"tongue/pkg/errno"
	"tongue/util"
)

// @Summary 修改密码
// @Description 修改用户密码
// @Tags user
// @Accept  json/application
// @Produce  json/application
// @Param Authorization header string true  "token 用户令牌"
// @Param req body updatePassword true  "用户的修改的密码"
// @Success 200  "Success"
// @Failure 400 {string} json  "{"Code":400, "Message":"Error occurred while binding the request body to the struct","Data":nil}"
// @Failure 500 {string} json  "{"Code":500,"Message":"Database error","Data":nil}"
// @Router /user/password [put]
func ChangePassword(c *gin.Context) {
	log.Info("student login function called.", zap.String("X-Request-Id", util.GetReqID(c)))
	var req updatePassword
	email := c.MustGet("email").(string)
	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	if err := user.UpdatePassword(email, req.OriginalPassword, req.NewPassword); err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "Success")
}
