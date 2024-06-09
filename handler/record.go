package handler

import (
	"errors"
	"log"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func ListRecord(c *gin.Context, req *models.ListRecordParam) (*models.ListRecordData, error) {

	users, err := service.ListRecord(req)
	if err != nil {
		log.Println(err)
		return nil, errors.New("记录信息查询失败")
	}

	return &models.ListRecordData{Records: users}, nil
}

func AddRecord(c *gin.Context, req *models.AddRecordParam) (interface{}, error) {

	if req.Record == nil {
		log.Printf("[AddRecord] req is nil, req: %+v", req)
		return nil, errors.New(common.ParamInvalid)
	}
	err := service.AddRecord(req.Record)
	if err != nil {
		log.Println(err)
		return nil, errors.New("记录信息保存失败")

	}
	return nil, nil
}

func DeleteRecord(c *gin.Context, req *models.DeleteRecordParam) (interface{}, error) {
	if req.RecordId == 0 {
		log.Printf("[DeleteRecord] user id is nil, req: %+v", req)
		return nil, errors.New(common.ParamInvalid)
	}
	err := service.DeleteRecord(req.RecordId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("记录信息删除失败")

	}
	return nil, nil
}
