package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/handler/user"
	"tongue/pkg/errno"
	service "tongue/service/user"
)

// @Summary 发送邮箱验证码
// @Tags auth
// @Description 邮箱验证码
// @Accept application/json
// @Produce application/json
// @Param object body user.VerificationRequest true "注册用户信息"
// @Success 200 {object} handler.Response "{"msg":"将student_id作为token保留"}"
// @Failure 401 {object} errno.Errno "{"error_code":"10001", "message":"The email address has been registered"} "
// @Failure 400 {object} errno.Errno "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} errno.Errno "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /auth/code [post]
func VerificationCode(c *gin.Context) {
	var req user.VerificationRequest
	if err := c.ShouldBind(&req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}
	code, err := service.VerificationCode(req.Email)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrRegister, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, gin.H{
		"VerificationCode": code,
	})
}

// @Summary Register
// @Tags auth
// @Description 邮箱注册登录
// @Accept application/json
// @Produce application/json
// @Param object body user.RegisterRequest true "注册用户信息"
// @Success 200 {object} handler.Response "{"msg":"将student_id作为token保留"}"
// @Failure 401 {object} errno.Errno "{"error_code":"10001", "message":"The email address has been registered"} "
// @Failure 400 {object} errno.Errno "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} errno.Errno "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var req user.RegisterRequest

	if err := c.ShouldBind(&req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine()) // , errno.ErrBind)
		fmt.Println(req)
		return
	}

	err := service.Register(req.Email, req.Name, req.Password, req.Code)

	if err != nil {
		handler.SendBadRequest(c, errno.ErrDatabase, nil, err.Error(), handler.GetLine())
		return
	}

	handler.SendResponse(c, nil, "succeed in registration")

}
