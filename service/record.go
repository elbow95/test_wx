package service

import (
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/models"
)

func RecordVo2Dto(r *model.Record) *models.Record {
	if r == nil {
		return nil
	}
	return &models.Record{
		Id:          r.Id,
		StationId:   r.StationId,
		OilId:       r.OilId,
		StaffId:     r.StaffId,
		Price:       r.Price,
		Liter:       r.Liter,
		Amount:      r.Amount,
		DriverId:    r.DriverId,
		DriverName:  r.DriverName,
		DriverPhone: r.DriverPhone,
		CreateTime:  r.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  r.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}

func RecordDto2Vo(r *models.Record) *model.Record {
	if r == nil {
		return nil
	}
	return &model.Record{
		Id:          r.Id,
		StationId:   r.StationId,
		OilId:       r.OilId,
		StaffId:     r.StaffId,
		Price:       r.Price,
		Liter:       r.Liter,
		Amount:      r.Amount,
		DriverId:    r.DriverId,
		DriverName:  r.DriverName,
		DriverPhone: r.DriverPhone,
	}
}

func ListRecord(param *models.ListRecordParam) ([]*models.Record, error) {
	query := db.Get()
	if param != nil {
		if len(param.Ids) > 0 {
			query = query.Where("id in (?)", param.Ids)
		}
		if len(param.OilIds) > 0 {
			query = query.Where("oil_id in (?)", param.OilIds)
		}
		if len(param.StaffIds) > 0 {
			query = query.Where("staff_id in (?)", param.StaffIds)
		}
		if len(param.StationIds) > 0 {
			query = query.Where("station_id in (?)", param.StationIds)
		}
		if param.CreateTime != nil {
			if param.CreateTime.Left > 0 {
				query = query.Where("create_time >= ?", time.Unix(param.CreateTime.Left, 0))
			}
			if param.CreateTime.Right > 0 {
				query = query.Where("create_time <= ?", time.Unix(param.CreateTime.Right, 0))
			}
		}
	}
	query = query.Where("is_delete = 0").Order("id desc")
	records := make([]*model.Record, 0)
	err := query.Find(&records, 0).Error
	if err != nil {
		return nil, err
	}

	results := make([]*models.Record, 0, len(records))
	for _, r := range records {
		results = append(results, RecordVo2Dto(r))
	}

	return results, nil
}

func AddRecord(r *models.Record) error {
	rVo := RecordDto2Vo(r)
	return db.Get().Create(rVo).Error
}

func DeleteRecord(recordId int64) error {
	updateFields := map[string]interface{}{
		"is_delete": 1,
	}
	return db.Get().Model(&model.Record{}).Where("id = ?", recordId).Updates(updateFields).Error
}
