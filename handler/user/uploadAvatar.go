package user

import (
	. "tongue/handler"
	"tongue/log"
	U "tongue/service/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tongue/util"
)

// UploadAvatar ... 上传头像
// @Summary Get Qiniuyun token
// @Description
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {string} json {"Code":200,"Token":"token"}
// @Router /user/qiniu_token [get]
func GetQiniuToken(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))
	Token := U.GetToken()
	SendResponse(c, nil, map[string]string{
		"Token": Token,
	})

}
