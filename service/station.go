package service

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/models"
)

func StationVo2Dto(s *db.Station) *models.Station {
	if s == nil {
		return nil
	}
	return &models.Station{
		Id:         s.Id,
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
		Id:        s.Id,
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
			query = query.Where("id in (?)", param.Ids)
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

func UpdateStation(s *models.Station) error {
	sVo := StationDto2Vo(s)
	return db.Get().Save(sVo).Error
}

func DeleteStation(stationId int64) error {
	updateFields := map[string]interface{}{
		"is_delete": 1,
	}
	return db.Get().Model(&db.Station{}).Where("id = ?", stationId).Updates(updateFields).Error
}
