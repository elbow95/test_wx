package middleware

import (
	"errors"
	"wxcloudrun-golang/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserInfo struct {
	Id        int64
	OpenId    string
	Type      int
	Name      string
	StationId int64
	Phone     string
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		openId := c.Request.Header.Get("X-Wx-Openid")
		c.Set("open_id", openId)

		userInfo := &db.User{}
		found, err := db.FindOne(&userInfo, func(db *gorm.DB) *gorm.DB {
			return db.Where("open_id = ?", openId).Where("is_delete = 0").Order("type")
		})

		if err != nil {
			c.AbortWithError(500, errors.New("服务错误，请稍后重试"))
			return
		} else if !found {
			c.AbortWithError(403, errors.New("用户未登录"))
			return
		}
		c.Set("user_info", &UserInfo{
			Id:        userInfo.Id,
			OpenId:    userInfo.OpenId,
			Type:      userInfo.Type,
			Name:      userInfo.Name,
			StationId: userInfo.StationId,
			Phone:     userInfo.Phone,
		})

		c.Next()
	}
}

func GetOpenId(c *gin.Context) string {
	if val, ok := c.Get("open_id"); ok {
		return val.(string)
	}
	return ""
}

func GetUser(c *gin.Context) *UserInfo {
	if val, ok := c.Get("user_info"); ok {
		return val.(*UserInfo)
	}
	return nil
}
