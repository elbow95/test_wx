package service

import (
	"errors"
	"fmt"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/util"

	"gorm.io/gorm"
)

func StationVo2Dto(s *db.Station) *models.Station {
	if s == nil {
		return nil
	}
	return &models.Station{
		Id:         util.Int642Str(s.Id),
		Name:       s.Name,
		Address:    s.Address,
		Longitude:  s.Longitude,
		Latitude:   s.Latitude,
		CreateTime: s.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: s.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}

func StationDto2Vo(s *models.Station) *db.Station {
	if s == nil {
		return nil
	}
	return &db.Station{
		Id:        util.Str2Int64(s.Id),
		Name:      s.Name,
		Address:   s.Address,
		Longitude: s.Longitude,
		Latitude:  s.Latitude,
	}
}

func ListStation(param *models.ListStationParam) ([]*models.Station, error) {
	query := db.Get()
	if param != nil {
		if len(param.Ids) > 0 {
			query = query.Where("id in (?)", util.StrSliceToInt64(param.Ids))
		}
		if param.Name != "" {
			query = query.Where("name like %?%", param.Name)
		}
	}
	query = query.Where("is_delete = 0").Order("id desc")
	stations := make([]*db.Station, 0)
	err := query.Find(&stations).Error
	if err != nil {
		return nil, err
	}
	results := make([]*models.Station, 0, len(stations))
	for _, s := range stations {
		results = append(results, StationVo2Dto(s))
	}
	return results, nil
}

func AddStation(s *models.Station) error {
	sVo := StationDto2Vo(s)
	return db.Get().Create(sVo).Error
}

func UpdateStation(param *models.UpdateStationParam) error {
	stationId := util.Str2Int64(param.StationId)
	if stationId == 0 {
		return errors.New("未指定加油站")
	}
	existStation := &db.Station{}
	found, err := db.FindOne(&existStation, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", stationId).Where("is_delete = 0")
	})
	if err != nil {
		fmt.Printf("查询加油站失败,err: %v", err)
		return errors.New("查询加油站失败")
	}
	if !found {
		return errors.New("指定加油站不存在")
	}
	if param.Name != nil && *param.Name != "" {
		existStation.Name = *param.Name
	}
	if param.Address != nil && *param.Address != "" {
		existStation.Address = *param.Address
	}
	if param.Longitude != nil && *param.Longitude > 0.0 {
		existStation.Longitude = *param.Longitude
	}
	if param.Latitude != nil && *param.Latitude > 0.0 {
		existStation.Latitude = *param.Latitude
	}
	return db.Get().Save(existStation).Error
}

func DeleteStation(stationIdStr string) error {
	stationId := util.Str2Int64(stationIdStr)
	if stationId == 0 {
		return errors.New("未指定加油站")
	}
	updateFields := map[string]interface{}{
		"is_delete": 1,
	}
	return db.Get().Model(&db.Station{}).Where("id = ?", stationId).Updates(updateFields).Error
}
