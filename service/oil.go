package service

import (
	"errors"
	"log"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/models"
)

func ListOil(param *models.ListOilParam) ([]*models.Oil, error) {
	query := db.Get()
	if param != nil {
		if len(param.Ids) > 0 {
			query = query.Where("id in (?)", param.Ids)
		}
		if len(param.StationIds) > 0 {
			query = query.Where("station_id in (?)", param.StationIds)
		}
		if param.Name != "" {
			query = query.Where("name like %?%", param.Name)
		}
	}
	query = query.Where("is_delete = 0").Order("id desc")
	oils := make([]*db.Oil, 0)
	err := query.Find(&oils).Error
	if err != nil {
		return nil, errors.New("查询油品信息失败")
	}

	stationIds := make([]int64, 0)
	for _, oil := range oils {
		stationIds = append(stationIds, oil.StationId)
	}
	stations := make([]*db.Station, 0)
	err = db.Get().Where("id in (?)", stationIds).Find(&stations).Error
	if err != nil {
		return nil, errors.New("查询油站信息失败")
	}
	stationMap := make(map[int64]*db.Station)
	for _, s := range stations {
		stationMap[s.Id] = s
	}

	results := make([]*models.Oil, 0, len(oils))
	for _, u := range oils {
		results = append(results, OilVo2Dto(u, stationMap))
	}

	return results, nil
}

func OilVo2Dto(o *db.Oil, sMap map[int64]*db.Station) *models.Oil {
	if o == nil {
		return nil
	}
	oil := &models.Oil{
		Id:         o.Id,
		Name:       o.Name,
		StationId:  o.StationId,
		Price:      o.Price,
		CreateTime: o.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: o.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	if s, ok := sMap[oil.StationId]; ok {
		oil.StationName = s.Name
	}
	return oil
}

func OilDto2Vo(o *models.Oil) *db.Oil {
	if o == nil {
		return nil
	}
	return &db.Oil{
		Id:        o.Id,
		Name:      o.Name,
		StationId: o.StationId,
		Price:     o.Price,
	}
}

func AddOil(s *models.Oil) error {
	if s.StationId == 0 {
		return errors.New("添加油品需要指定加油站")
	}
	stations := make([]*db.Oil, 0)
	if err := db.Get().Where("id = ?", s.StationId).Find(&stations).Error; err != nil {
		log.Printf("[AddOil] get station failed, err: %+v", err)
		return errors.New("查找加油站失败")
	} else if len(stations) == 0 {
		return errors.New("指定的加油站不存在")
	}

	uVo := OilDto2Vo(s)
	return db.Get().Create(uVo).Error
}

func UpdateOil(s *models.Oil) error {
	uVo := OilDto2Vo(s)
	return db.Get().Save(uVo).Error
}

func DeleteOil(oilId int64) error {
	updateFields := map[string]interface{}{
		"is_delete": 1,
	}
	return db.Get().Model(&db.Oil{}).Where("id = ?", oilId).Updates(updateFields).Error
}
