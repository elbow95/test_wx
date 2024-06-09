package handler

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
	"wxcloudrun-golang/util"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type GetPhoneParam struct {
	CloudId string `json:"cloud_id"`
}

type GetPhoneData struct {
	Phone string `json:"phone"`
}

func GetPhone(c *gin.Context, req *GetPhoneParam) (*GetPhoneData, error) {
	if req == nil || req.CloudId == "" {
		return nil, errors.New("获取手机号参数错误，请重试")
	}
	reqData := map[string]interface{}{
		"cloudid_list": []string{req.CloudId},
	}
	url := fmt.Sprintf("http://api.weixin.qq.com/wxa/getopendata?openid=%s", c.Request.Header.Get("X-Wx-Openid"))
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(util.MarshalJsonIgnoreError(reqData))))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("获取用户手机号失败，请重试")
	}
	if resp == nil || resp.Body == nil || resp.StatusCode != 200 {
		fmt.Println("获取用户手机号返回信息有误")
		return nil, errors.New("获取用户手机号返回失败，请重试")
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取返回信息失败")
		return nil, errors.New("获取返回信息失败，请重试")
	}
	fmt.Printf("get phone resp body: %s", string(respBody))
	phone := gjson.Get(string(respBody), "phoneNumber").String()
	fmt.Println(phone)
	return &GetPhoneData{Phone: phone}, nil
}
