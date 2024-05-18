package service

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
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
	oils := make([]*model.Oil, 0)
	err := query.Find(&oils).Error
	if err != nil {
		return nil, err
	}

	results := make([]*models.Oil, 0, len(oils))
	for _, u := range oils {
		results = append(results, OilVo2Dto(u))
	}

	return results, nil
}

func OilVo2Dto(s *model.Oil) *models.Oil {
	if s == nil {
		return nil
	}
	return &models.Oil{
		Id:         s.Id,
		Name:       s.Name,
		StationId:  s.StationId,
		Price:      s.Price,
		CreateTime: s.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: s.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}

func OilDto2Vo(s *models.Oil) *model.Oil {
	if s == nil {
		return nil
	}
	return &model.Oil{
		Id:        s.Id,
		Name:      s.Name,
		StationId: s.StationId,
		Price:     s.Price,
	}
}

func AddOil(s *models.Oil) error {
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
	return db.Get().Model(&model.Oil{}).Where("id = ?", oilId).Updates(updateFields).Error
}
