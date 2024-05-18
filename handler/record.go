package handler

import (
	"fmt"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func ListRecord(c *gin.Context) {
	var param models.ListRecordParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}

	users, err := service.ListRecord(&param)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "记录信息查询失败")
		return
	}

	common.Success(c, &models.ListRecordData{Records: users})
}

func AddRecord(c *gin.Context) {
	var param models.AddRecordParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}
	if param.Record == nil {
		fmt.Printf("[AddRecord] param is nil, param: %+v", param)
		common.Failed(c, common.ParamInvalid)
		return
	}
	err := service.AddRecord(param.Record)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "记录信息保存失败")
		return
	}
	common.Success(c, nil)
}

func DeleteRecord(c *gin.Context) {
	var param models.DeleteRecordParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		common.Failed(c, common.ParamInvalid)
		return
	}
	if param.RecordId == 0 {
		fmt.Printf("[DeleteRecord] user id is nil, param: %+v", param)
		common.Failed(c, common.ParamInvalid)
		return
	}
	err := service.DeleteRecord(param.RecordId)
	if err != nil {
		fmt.Println(err)
		common.Failed(c, "记录信息删除失败")
		return
	}
	common.Success(c, nil)
}
