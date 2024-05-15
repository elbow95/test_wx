package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"wxcloudrun-golang/common"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IndexHandler 计数器接口
func IndexHandler(ctx *gin.Context) {
	data, err := getIndex()
	if err != nil {
		common.Failed(ctx, err.Error())
		return
	}
	common.Success(ctx, data)
}

// CounterHandler 计数器接口
func CounterHandler(ctx *gin.Context) {

	if ctx.Request.Method == http.MethodGet {
		counter, err := getCurrentCounter()
		if err != nil {
			common.Failed(ctx, err.Error())
			return
		} else {
			common.Success(ctx, counter.Count)
			return
		}
	} else if ctx.Request.Method == http.MethodPost {
		count, err := modifyCounter(ctx.Request)
		if err != nil {
			common.Failed(ctx, err.Error())
			return
		} else {
			common.Success(ctx, count)
			return
		}
	} else {
		common.Failed(ctx, fmt.Sprintf("请求方法 %s 不支持", ctx.Request.Method))
		return
	}
}

// modifyCounter 更新计数，自增或者清零
func modifyCounter(r *http.Request) (int32, error) {
	action, err := getAction(r)
	if err != nil {
		return 0, err
	}

	var count int32
	if action == "inc" {
		count, err = upsertCounter(r)
		if err != nil {
			return 0, err
		}
	} else if action == "clear" {
		err = clearCounter()
		if err != nil {
			return 0, err
		}
		count = 0
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return count, err
}

// upsertCounter 更新或修改计数器
func upsertCounter(r *http.Request) (int32, error) {
	currentCounter, err := getCurrentCounter()
	var count int32
	createdAt := time.Now()
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	} else if err == gorm.ErrRecordNotFound {
		count = 1
		createdAt = time.Now()
	} else {
		count = currentCounter.Count + 1
		createdAt = currentCounter.CreatedAt
	}

	counter := &model.CounterModel{
		Id:        1,
		Count:     count,
		CreatedAt: createdAt,
		UpdatedAt: time.Now(),
	}
	err = dao.Imp.UpsertCounter(counter)
	if err != nil {
		return 0, err
	}
	return counter.Count, nil
}

func clearCounter() error {
	return dao.Imp.ClearCounter(1)
}

// getCurrentCounter 查询当前计数器
func getCurrentCounter() (*model.CounterModel, error) {
	counter, err := dao.Imp.GetCounter(1)
	if err != nil {
		return nil, err
	}

	return counter, nil
}

// getAction 获取action
func getAction(r *http.Request) (string, error) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return "", err
	}
	defer r.Body.Close()

	action, ok := body["action"]
	if !ok {
		return "", fmt.Errorf("缺少 action 参数")
	}

	return action.(string), nil
}

// getIndex 获取主页
func getIndex() (string, error) {
	b, err := ioutil.ReadFile("./index.html")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
