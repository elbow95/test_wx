package handler

import (
	"log"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/middleware"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/service"
	"wxcloudrun-golang/util"

	"github.com/gin-gonic/gin"
)

func ListStation(c *gin.Context, req *models.ListStationParam) (*models.ListStationData, error) {
	// 获取用户油站，管理员全部、加油员自己、司机全部
	loginUser := middleware.GetUser(c)

	if len(req.Ids) == 0 && loginUser.Type == int(common.UserType_Staff) {
		req.Ids = []string{util.Int642Str(loginUser.StationId)}
	}
	stations, total, err := service.ListStation(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &models.ListStationData{Stations: stations, Total: total}, nil
}

func AddStation(c *gin.Context, req *models.AddStationParam) (interface{}, error) {
	err := service.AddStation(req.Station)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return nil, nil
}

func UpdateStation(c *gin.Context, req *models.UpdateStationParam) (interface{}, error) {

	err := service.UpdateStation(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return nil, nil
}

func DeleteStation(c *gin.Context, req *models.DeleteStationParam) (interface{}, error) {

	err := service.DeleteStation(req.StationId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return nil, nil
}
