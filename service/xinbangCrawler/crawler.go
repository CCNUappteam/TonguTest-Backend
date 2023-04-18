package xinbangCrawler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type requestData struct {
	Date      string `json:"date"`
	DateType  string `json:"date_type"`
	FirstType string `json:"firstType"`
	Size      int    `json:"size"`
	Start     int    `json:"start"`
}
type AccountInfo struct {
	AccountClassifyFirst       string      `json:"account_classify_first"`
	AccountClassifySecond      string      `json:"account_classify_second"`
	AddFollowerCount           string      `json:"add_follower_count"`
	AddMplatformFollowersCount string      `json:"add_mplatform_followers_count"`
	AwemeCount                 string      `json:"aweme_count"`
	City                       string      `json:"city"`
	CollectionStatus           string      `json:"collection_status"`
	CommentCount               string      `json:"comment_count"`
	CustomVerify               string      `json:"custom_verify"`
	DailyStatus                interface{} `json:"dailyStatus"`
	DiggCount                  string      `json:"digg_count"`
	EnterpriseVerifyReason     string      `json:"enterprise_verify_reason"`
	FollowerCount              string      `json:"follower_count"`
	ImageURL                   string      `json:"image_url"`
	McnDesc                    string      `json:"mcn_desc"`
	McnID                      string      `json:"mcn_id"`
	McnName                    string      `json:"mcn_name"`
	MplatformFollowersCount    string      `json:"mplatform_followers_count"`
	NewrankIndex               string      `json:"newrank_index"`
	Nickname                   string      `json:"nickname"`
	Province                   string      `json:"province"`
	RankPosition               string      `json:"rank_position"`
	SecUid                     string      `json:"sec_uid"`
	ShareCount                 string      `json:"share_count"`
	ShortID                    string      `json:"short_id"`
	TotalFavorited             string      `json:"total_favorited"`
	Type                       string      `json:"type"`
	Uid                        string      `json:"uid"`
	UniqueID                   string      `json:"unique_id"`
	UpdateTime                 string      `json:"update_time"`
	VerifyLabel                string      `json:"verify_label"`
	Vid                        interface{} `json:"vid"`
	WorksCount                 string      `json:"works_count"`
}

type ResponseData struct {
	Msg  string `json:"msg"`
	Data struct {
		Count int           `json:"count"`
		List  []AccountInfo `json:"list"`
	} `json:"data"`
}

func SimpleCrawler(token, date, datetype string, size int, start int) (*ResponseData, error) {
	url := "https://gw.newrank.cn/api/xd/xdnphb/nr/cloud/douyin/rank/mainHotAccountAllRankList"

	data := requestData{
		Date:      date,
		DateType:  datetype,
		FirstType: "医疗健康",
		Size:      size,
		Start:     start,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("N-Token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	var responseData ResponseData
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return nil, err
	}
	return &responseData, nil
}
