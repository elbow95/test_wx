package handler

import (
	"errors"
	"log"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context, req *models.ListUserParam) (*models.ListUserData, error) {

	users, total, err := service.ListUser(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &models.ListUserData{Users: users, Total: total}, nil
}

func AddUser(c *gin.Context, req *models.AddUserParam) (interface{}, error) {
	if req.User == nil {
		log.Printf("[AddUser] req is nil, req: %+v", req)
		return nil, errors.New(common.ParamInvalid)
	}
	err := service.AddUser(req.User)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}

func UpdateUser(c *gin.Context, req *models.UpdateUserParam) (interface{}, error) {
	err := service.UpdateUser(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}

func DeleteUser(c *gin.Context, req *models.DeleteUserParam) (interface{}, error) {
	err := service.DeleteUser(req.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}
