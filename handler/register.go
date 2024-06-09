package handler

import "github.com/gin-gonic/gin"

type RegisterParam struct {
	ExistUserId string `json:"exist_user_id"`
}

func Register(c *gin.Context, req *GetUserInfoParam) (*GetUserInfoData, error) {
	return nil, nil
}
