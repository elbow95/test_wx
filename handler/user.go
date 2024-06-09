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
		return nil, errors.New("人员信息查询失败")
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
		return nil, errors.New("人员信息保存失败")
	}
	return nil, nil
}

func UpdateUser(c *gin.Context, req *models.UpdateUserParam) (interface{}, error) {
	if req.User == nil || req.User.Id == 0 {
		log.Printf("[UpdateUser] req is nil, req: %+v", req)
		return nil, errors.New(common.ParamInvalid)
	}
	err := service.UpdateUser(req.User)
	if err != nil {
		log.Println(err)
		return nil, errors.New("人员信息更新失败")
	}
	return nil, nil
}

func DeleteUser(c *gin.Context, req *models.DeleteUserParam) (interface{}, error) {
	if req.UserId == 0 {
		log.Printf("[DeleteUser] user id is nil, req: %+v", req)
		return nil, errors.New(common.ParamInvalid)
	}
	err := service.DeleteUser(req.UserId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("人员信息删除失败")
	}
	return nil, nil
}
