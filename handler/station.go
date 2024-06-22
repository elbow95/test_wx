package handler

import (
	"log"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func ListStation(c *gin.Context, req *models.ListStationParam) (*models.ListStationData, error) {
	stations, err := service.ListStation(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &models.ListStationData{Stations: stations}, nil
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
