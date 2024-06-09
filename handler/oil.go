package handler

import (
	"errors"
	"log"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func ListOil(c *gin.Context, req *models.ListOilParam) (*models.ListOilData, error) {
	users, err := service.ListOil(req)
	if err != nil {
		log.Println(err)
		return nil, errors.New("油品信息查询失败")
	}

	return &models.ListOilData{Oils: users}, nil
}

func AddOil(c *gin.Context, req *models.AddOilParam) (interface{}, error) {

	if req.Oil == nil {
		log.Printf("[AddOil] req is nil, req: %+v", req)
		return nil, errors.New(common.ParamInvalid)
	}
	err := service.AddOil(req.Oil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}

func UpdateOil(c *gin.Context, req *models.UpdateOilParam) (interface{}, error) {
	if req.Oil == nil || req.Oil.Id == 0 {
		log.Printf("[UpdateOil] req is nil, req: %+v", req)
		return nil, errors.New(common.ParamInvalid)
	}
	err := service.UpdateOil(req.Oil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("油品信息更新失败")
	}
	return nil, nil
}

func DeleteOil(c *gin.Context, req *models.DeleteOilParam) (interface{}, error) {

	if req.OilId == 0 {
		log.Printf("[DeleteOil] user id is nil, req: %+v", req)
		return nil, errors.New(common.ParamInvalid)
	}
	err := service.DeleteOil(req.OilId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("油品信息删除失败")
	}
	return nil, nil
}
