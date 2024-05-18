package handler

import (
	"fmt"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func ListStation(c *gin.Context) {
	var param models.ListStationParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}

	stations, err := service.ListStation(&param)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "车站信息查询失败")
		return
	}

	common.Success(c, &models.ListStationData{Stations: stations})
}

func AddStation(c *gin.Context) {
	var param models.AddStationParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}

	err := service.AddStation(param.Station)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "车站信息添加失败")
		return
	}

	common.Success(c, nil)
}

func UpdateStation(c *gin.Context) {
	var param models.UpdateStationParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}

	err := service.UpdateStation(param.Station)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "车站信息更新失败")
		return
	}

	common.Success(c, nil)
}

func DeleteStation(c *gin.Context) {
	var param models.DeleteStationParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}

	err := service.DeleteStation(param.StationId)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "车站信息删除失败")
		return
	}

	common.Success(c, nil)
}
