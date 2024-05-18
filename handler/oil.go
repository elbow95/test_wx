package handler

import (
	"fmt"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func ListOil(c *gin.Context) {
	var param models.ListOilParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}

	users, err := service.ListOil(&param)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "油品信息查询失败")
		return
	}

	common.Success(c, &models.ListOilData{Oils: users})
}

func AddOil(c *gin.Context) {
	var param models.AddOilParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}
	if param.Oil == nil {
		fmt.Printf("[AddOil] param is nil, param: %+v", param)
		common.Failed(c, common.ParamInvalid)
		return
	}
	err := service.AddOil(param.Oil)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "油品信息保存失败")
		return
	}
	common.Success(c, nil)
}

func UpdateOil(c *gin.Context) {
	var param models.UpdateOilParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}
	if param.Oil == nil || param.Oil.Id == 0 {
		fmt.Printf("[UpdateOil] param is nil, param: %+v", param)
		common.Failed(c, common.ParamInvalid)
		return
	}
	err := service.UpdateOil(param.Oil)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "油品信息更新失败")
		return
	}
	common.Success(c, nil)
}

func DeleteOil(c *gin.Context) {
	var param models.DeleteOilParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}
	if param.OilId == 0 {
		fmt.Printf("[DeleteOil] user id is nil, param: %+v", param)
		common.Failed(c, common.ParamInvalid)
		return
	}
	err := service.DeleteOil(param.OilId)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "油品信息删除失败")
		return
	}
	common.Success(c, nil)
}
