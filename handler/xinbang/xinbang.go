package xinbang

import (
	"github.com/gin-gonic/gin"
	"tongue/handler"
	"tongue/pkg/errno"
	"tongue/service/xinbangCrawler"
)

type GetRankRequest struct {
	Token    string `json:"token"`
	Date     string `json:"date"`
	DataType string `json:"data_type"`
	Size     int    `json:"size"`
	Start    int    `json:"start"`
}

func GetRank(c *gin.Context) {
	var req GetRankRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}
	resp, err := xinbangCrawler.SimpleCrawler(req.Token, req.Date, req.DataType, req.Size, req.Start)
	if err != nil {
		handler.SendError(c, err, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, gin.H{
		"result": resp,
	})
}
