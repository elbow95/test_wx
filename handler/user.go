package handler

import (
	"fmt"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	var param models.ListUserParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}

	users, err := service.ListUser(&param)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "人员信息查询失败")
		return
	}

	common.Success(c, &models.ListUserData{Users: users})
}

func AddUser(c *gin.Context) {
	var param models.AddUserParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}
	if param.User == nil {
		fmt.Printf("[AddUser] param is nil, param: %+v", param)
		common.Failed(c, common.ParamInvalid)
		return
	}
	err := service.AddUser(param.User)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "人员信息保存失败")
		return
	}
	common.Success(c, nil)
}

func UpdateUser(c *gin.Context) {
	var param models.UpdateUserParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}
	if param.User == nil || param.User.Id == 0 {
		fmt.Printf("[UpdateUser] param is nil, param: %+v", param)
		common.Failed(c, common.ParamInvalid)
		return
	}
	err := service.UpdateUser(param.User)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "人员信息更新失败")
		return
	}
	common.Success(c, nil)
}

func DeleteUser(c *gin.Context) {
	var param models.DeleteUserParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}
	if param.UserId == 0 {
		fmt.Printf("[DeleteUser] user id is nil, param: %+v", param)
		common.Failed(c, common.ParamInvalid)
		return
	}
	err := service.DeleteUser(param.UserId)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "人员信息删除失败")
		return
	}
	common.Success(c, nil)
}
