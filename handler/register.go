package handler

import (
	"fmt"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/service"
	"wxcloudrun-golang/util"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"gorm.io/gorm"
)

type RegisterParam struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	PlateNumber string `json:"plate_number"`
	Company     string `json:"company"`
	License     string `json:"license"`
}

type RegisterData struct {
	User *models.User `json:"user"`
}

func Register(c *gin.Context, req *RegisterParam) (*RegisterData, error) {

	existUser := &db.User{}
	// 根据手机号查询用户
	found, err := db.FindOneWithWrite(&existUser, func(db *gorm.DB) *gorm.DB {
		return db.Where("phone = ?", req.Phone).Where("type = ?", int(common.UserType_Driver)).
			Where("is_delete = 0")
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("found: %v, exist user: %s\n", found, util.MarshalJsonIgnoreError(existUser))
	if found {
		// 存在已经录入的司机，更新司机的open_id
		existUser.OpenId = c.Request.Header.Get("X-Wx-Openid")
		if existUser.Name == "" {
			existUser.Name = req.Name
		}
		if existUser.PlateNumber == "" {
			existUser.PlateNumber = req.PlateNumber
		}
		if gjson.Get(existUser.Extra, "company").String() == "" {
			existUser.Extra, _ = sjson.Set(existUser.Extra, "company", req.Company)
		}
		if gjson.Get(existUser.Extra, "license").String() == "" {
			existUser.Extra, _ = sjson.Set(existUser.Extra, "license", req.License)
		}
	} else {
		// 不存在司机，新建
		existUser = &db.User{
			OpenId:      c.Request.Header.Get("X-Wx-Openid"),
			Name:        req.Name,
			Type:        int(common.UserType_Driver),
			Phone:       req.Phone,
			PlateNumber: req.PlateNumber,
			Extra: util.MarshalJsonIgnoreError(&db.UserExtra{
				License: req.License,
				Company: req.Company,
			}),
			Status: int(common.UserStatus_WaitVerify),
		}
	}
	fmt.Printf("exist user after process: %s", util.MarshalJsonIgnoreError(existUser))
	err = db.GetWrite().Save(existUser).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &RegisterData{User: service.UserVo2Dto(existUser, nil)}, nil
}
