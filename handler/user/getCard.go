package user

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/model/user"
)

func GetCard(c *gin.Context) {
	year := c.Query("year")
	month := c.Query("month")
	email := c.MustGet("email").(string)
	cards, err := user.GetCard(email, year, month)
	if err != nil {
		handler.SendError(c, err, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, cards)
}
