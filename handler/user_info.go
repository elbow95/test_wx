package handler

import (
	"errors"
	"fmt"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetUserInfoParam struct{}

type GetUserInfoData struct {
	User *models.User `json:"user"`
}

func GetUserInfo(c *gin.Context, req *GetUserInfoParam) (*GetUserInfoData, error) {

	openId := c.Request.Header.Get("X-Wx-Openid")
	fmt.Printf("open_id: %s\n", openId)
	if openId == "" {
		return nil, errors.New("用户未登录")
	}
	existUser := &db.User{}
	found, err := db.FindOne(&existUser, func(db *gorm.DB) *gorm.DB {
		return db.Where("open_id = ?", openId).Where("is_delete = 0").Order("type")
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if found {
		// 用户存在，直接返回
		station := &db.Station{}
		_, _ = db.FindOne(&station, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", existUser.StationId).Where("is_delete = 0")
		})
		userData := service.UserVo2Dto(existUser, map[int64]*db.Station{existUser.StationId: station})
		return &GetUserInfoData{User: userData}, nil
	} else {
		// 用户不存在，走登录环节
		return &GetUserInfoData{User: &models.User{}}, nil
	}
}
