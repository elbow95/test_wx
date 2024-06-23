package service

import (
	"errors"
	"log"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/util"

	"gorm.io/gorm"
)

func ListOil(param *models.ListOilParam) ([]*models.Oil, error) {
	if len(param.StationIds) > 0 {
		stations := make([]*db.Station, 0)
		found, err := db.FindData(&stations, func(db *gorm.DB) *gorm.DB {
			return db.Where("id in (?)", param.StationIds).Where("is_delete = 0")
		})
		if err != nil {
			return nil, errors.New("查询油站信息失败")
		}
		if !found {
			return nil, nil
		}
		newStationIds := make([]string, 0)
		for _, station := range stations {
			newStationIds = append(newStationIds, util.Int642Str(station.Id))
		}
		param.StationIds = newStationIds
	}

	query := db.Get()
	if param != nil {
		if len(param.Ids) > 0 {
			query = query.Where("id in (?)", util.StrSliceToInt64(param.Ids))
		}
		if len(param.StationIds) > 0 {
			query = query.Where("station_id in (?)", util.StrSliceToInt64(param.StationIds))
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
		Id:         util.Int642Str(o.Id),
		Name:       o.Name,
		StationId:  util.Int642Str(o.StationId),
		Price:      o.Price,
		CreateTime: o.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: o.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	if s, ok := sMap[o.StationId]; ok {
		oil.StationName = s.Name
	}
	return oil
}

func OilDto2Vo(o *models.Oil) *db.Oil {
	if o == nil {
		return nil
	}
	return &db.Oil{
		Id:        util.Str2Int64(o.Id),
		Name:      o.Name,
		StationId: util.Str2Int64(o.StationId),
		Price:     o.Price,
	}
}

func AddOil(s *models.Oil) error {
	stationId := util.Str2Int64(s.StationId)
	if stationId == 0 {
		return errors.New("添加油品需要指定加油站")
	}

	station := &db.Station{}
	found, err := db.FindOne(&station, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", stationId).Where("is_delete = 0")
	})
	if err != nil {
		log.Printf("[AddOil] get station failed, err: %+v", err)
		return errors.New("查找加油站失败")
	} else if !found {
		return errors.New("指定的加油站不存在")
	}
	oil := &db.Oil{}
	found, err = db.FindOne(&oil, func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", s.Name).Where("station_id = ?", stationId).Where("is_delete = 0")
	})
	if err != nil {
		log.Printf("[AddOil] get oil failed, err: %+v", err)
		return errors.New("查找油品失败")
	} else if found {
		return errors.New("该加油站已有同名油品")
	}

	uVo := OilDto2Vo(s)
	return db.Get().Create(uVo).Error
}

func UpdateOil(param *models.UpdateOilParam) error {
	oilId := util.Str2Int64(param.OilId)
	if oilId == 0 {
		return errors.New("未指定油品")
	}
	if param.Price == nil {
		return nil
	}
	oil := &db.Oil{}
	found, err := db.FindOne(&oil, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", oilId).Where("is_delete = 0")
	})
	if err != nil {
		log.Printf("[UpdateOil] get oil failed, err: %+v", err)
		return errors.New("查找油品失败")
	} else if !found {
		return errors.New("指定的油品不存在")
	}
	oil.Price = *param.Price

	return db.Get().Save(oil).Error
}

func DeleteOil(oilIdStr string) error {
	oilId := util.Str2Int64(oilIdStr)
	if oilId == 0 {
		return errors.New("未指定油品")
	}
	updateFields := map[string]interface{}{
		"is_delete": 1,
	}
	return db.Get().Model(&db.Oil{}).Where("id = ?", oilId).Updates(updateFields).Error
}

func BatchUpdateOilPrice(param *models.BatchUpdateOilPriceParam) error {
	if len(param.Oils) == 0 {
		return nil
	}
	oilIds := make([]int64, 0)
	oilPriceMap := make(map[int64]float64)
	for _, oil := range param.Oils {
		if oil.Price <= 0.0 {
			continue
		}
		oilPriceMap[util.Str2Int64(oil.Id)] = oil.Price
		oilIds = append(oilIds, util.Str2Int64(oil.Id))
	}
	err := db.GetWrite().Transaction(func(tx *gorm.DB) error {
		oils := make([]*db.Oil, 0)
		_, err := db.FindData(&oils, func(db *gorm.DB) *gorm.DB {
			return db.Where("id in (?)", oilIds)
		})
		if err != nil {
			log.Fatalf("查询待更新油品信息失败，err: %v", err)
			return errors.New("查询待更新油品失败，请重试")
		}
		for _, oil := range oils {
			if err = tx.Model(&db.Oil{}).Where("id = ?", oil.Id).Updates(map[string]interface{}{
				"price": oilPriceMap[oil.Id],
			}).Error; err != nil {
				log.Fatalf("更新待更新油品信息失败，err: %v", err)
				return errors.New("更新油品信息失败，请重试")
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}
