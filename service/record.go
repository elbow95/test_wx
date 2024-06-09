package service

import (
	"errors"
	"log"
	"sync"
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/util"
)

func RecordVo2Dto(r *db.Record, stationMap map[int64]*db.Station, staffMap map[int64]*db.User, oilMap map[int64]*db.Oil) *models.Record {
	if r == nil {
		return nil
	}
	record := &models.Record{
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
	if s, ok := stationMap[r.StationId]; ok {
		record.StationName = s.Name
	}
	if s, ok := staffMap[r.StaffId]; ok {
		record.StaffName = s.Name
	}
	if o, ok := oilMap[r.OilId]; ok {
		record.OilName = o.Name
	}
	return record
}

func RecordDto2Vo(r *models.Record) *db.Record {
	if r == nil {
		return nil
	}
	return &db.Record{
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
	records := make([]*db.Record, 0)
	err := query.Find(&records, 0).Error
	if err != nil {
		return nil, err
	}

	var (
		stationIds []int64
		staffIds   []int64
		oilIds     []int64
		stationMap map[int64]*db.Station
		staffMap   map[int64]*db.User
		oilMap     map[int64]*db.Oil
		wg         sync.WaitGroup
	)
	for _, r := range records {
		stationIds = append(stationIds, r.StationId)
		staffIds = append(staffIds, r.StaffId)
		oilIds = append(oilIds, r.OilId)
	}
	wg.Add(1)
	util.GoWithDefaultRecovery(func() {
		defer wg.Done()
		stations := make([]*db.Station, 0)
		_ = db.Get().Where("id in (?)", stationIds).Find(&stations).Error
		for _, s := range stations {
			stationMap[s.Id] = s
		}
	})
	wg.Add(1)
	util.GoWithDefaultRecovery(func() {
		defer wg.Done()
		users := make([]*db.User, 0)
		_ = db.Get().Where("id in (?)", staffIds).Find(&users).Error
		for _, u := range users {
			staffMap[u.Id] = u
		}
	})
	wg.Add(1)
	util.GoWithDefaultRecovery(func() {
		defer wg.Done()
		oils := make([]*db.Oil, 0)
		_ = db.Get().Where("id in (?)", oilIds).Find(&oils).Error
		for _, o := range oils {
			oilMap[o.Id] = o
		}
	})
	wg.Wait()

	results := make([]*models.Record, 0, len(records))
	for _, r := range records {
		results = append(results, RecordVo2Dto(r, stationMap, staffMap, oilMap))
	}

	return results, nil
}

func AddRecord(r *models.Record) error {
	if r.StationId == 0 {
		return errors.New("未指定加油站")
	}
	if r.StaffId == 0 {
		return errors.New("未指定加油员")
	}
	if r.OilId == 0 {
		return errors.New("未指定油品")
	}
	if r.Price <= 0 {
		return errors.New("油品单价非法")
	}
	if r.Liter <= 0 {
		return errors.New("加油升数非法")
	}
	if r.Amount <= 0 {
		return errors.New("总价非法")
	}
	if r.DriverName == "" {
		return errors.New("未填写驾驶员姓名")
	}
	if r.DriverPhone == "" {
		return errors.New("未填写驾驶员手机号")
	}
	var (
		station = &db.Station{}
		staff   = &db.User{}
		oil     = &db.Oil{}
	)
	if err := db.Get().Where("id = ?", r.StationId).Where("is_delete = 0").First(&station); err != nil {
		log.Printf("[AddRecord] get station failed, err: %+v", err)
		return errors.New("获取油站信息失败")
	} else if station == nil || station.Id == 0 {
		log.Printf("[AddRecord] station not found, id: %+d", r.StationId)
		return errors.New("指定油站不存在")
	}
	if err := db.Get().Where("id = ?", r.StaffId).Where("is_delete = 0").First(&staff); err != nil {
		log.Printf("[AddRecord] get staff failed, err: %+v", err)
		return errors.New("获取加油员信息失败")
	} else if staff == nil || staff.Id == 0 {
		log.Printf("[AddRecord] staff not found, id: %+d", r.StaffId)
		return errors.New("指定加油员不存在")
	}
	if err := db.Get().Where("id = ?", r.OilId).Where("is_delete = 0").First(&oil); err != nil {
		log.Printf("[AddRecord] get oil failed, err: %+v", err)
		return errors.New("获取油品信息失败")
	} else if oil == nil || oil.Id == 0 {
		log.Printf("[AddRecord] oil not found, id: %+d", r.OilId)
		return errors.New("指定油品不存在")
	}

	rVo := RecordDto2Vo(r)
	return db.Get().Create(rVo).Error
}

func DeleteRecord(recordId int64) error {
	updateFields := map[string]interface{}{
		"is_delete": 1,
	}
	return db.Get().Model(&db.Record{}).Where("id = ?", recordId).Updates(updateFields).Error
}
