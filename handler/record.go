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

	users, total, err := service.ListRecord(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &models.ListRecordData{Records: users, Total: total}, nil
}

func AddRecord(c *gin.Context, req *models.AddRecordParam) (interface{}, error) {

	if req.Record == nil {
		log.Printf("[AddRecord] req is nil, req: %+v", req)
		return nil, errors.New(common.ParamInvalid)
	}
	err := service.AddRecord(req.Record)
	if err != nil {
		log.Println(err)
		return nil, err

	}
	return nil, nil
}

func UpdateRecord(c *gin.Context, req *models.UpdateRecrodParam) (interface{}, error) {
	err := service.UpdateRecord(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}

func DeleteRecord(c *gin.Context, req *models.DeleteRecordParam) (interface{}, error) {
	err := service.DeleteRecord(req.RecordId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}
