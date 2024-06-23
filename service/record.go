package service

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/util"

	"gorm.io/gorm"
)

func RecordVo2Dto(r *db.Record, stationMap map[int64]*db.Station, staffMap map[int64]*db.User, oilMap map[int64]*db.Oil) *models.Record {
	if r == nil {
		return nil
	}
	record := &models.Record{
		Id:          util.Int642Str(r.Id),
		StationId:   util.Int642Str(r.StationId),
		OilId:       util.Int642Str(r.OilId),
		StaffId:     util.Int642Str(r.StaffId),
		Price:       r.Price,
		Liter:       r.Liter,
		Amount:      r.Amount,
		DriverId:    util.Int642Str(r.DriverId),
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
		record.StaffPhone = s.Phone
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
		Id:          util.Str2Int64(r.Id),
		StationId:   util.Str2Int64(r.StationId),
		OilId:       util.Str2Int64(r.OilId),
		StaffId:     util.Str2Int64(r.StaffId),
		Price:       r.Price,
		Liter:       r.Liter,
		Amount:      r.Amount,
		DriverId:    util.Str2Int64(r.DriverId),
		DriverName:  r.DriverName,
		DriverPhone: r.DriverPhone,
	}
}

func ListRecord(param *models.ListRecordParam) ([]*models.Record, int64, error) {
	query := db.Get()
	if param != nil {
		if len(param.Ids) > 0 {
			query = query.Where("id in (?)", util.StrSliceToInt64(param.Ids))
		}
		if len(param.OilIds) > 0 {
			query = query.Where("oil_id in (?)", util.StrSliceToInt64(param.OilIds))
		}
		if len(param.StaffIds) > 0 {
			query = query.Where("staff_id in (?)", util.StrSliceToInt64(param.StaffIds))
		}
		if len(param.StationIds) > 0 {
			query = query.Where("station_id in (?)", util.StrSliceToInt64(param.StationIds))
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
	var (
		total int64
	)
	err := query.Model(&db.Record{}).Count(&total).Error
	if err != nil {
		fmt.Printf("查询记录失败,err:%v", err)
		return nil, 0, errors.New("查询记录失败")
	}
	records := make([]*db.Record, 0)
	query = db.AttachPage(query, param.Page, param.PageSize)
	err = query.Find(&records).Error
	if err != nil {
		fmt.Printf("查询记录失败,err:%v", err)
		return nil, 0, errors.New("查询记录失败")
	}

	var (
		stationIds = make([]int64, 0)
		staffIds   = make([]int64, 0)
		oilIds     = make([]int64, 0)
		stationMap = make(map[int64]*db.Station)
		staffMap   = make(map[int64]*db.User)
		oilMap     = make(map[int64]*db.Oil)
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
		err = db.Get().Where("id in (?)", stationIds).Find(&stations).Error
		if err != nil {
			fmt.Printf("获取油站信息失败：err: %+v", err)
		}
		for _, s := range stations {
			stationMap[s.Id] = s
		}
	})
	wg.Add(1)
	util.GoWithDefaultRecovery(func() {
		defer wg.Done()
		users := make([]*db.User, 0)
		err = db.Get().Where("id in (?)", staffIds).Find(&users).Error
		if err != nil {
			fmt.Printf("获取加油员信息失败：err: %+v", err)
		}
		for _, u := range users {
			staffMap[u.Id] = u
		}
	})
	wg.Add(1)
	util.GoWithDefaultRecovery(func() {
		defer wg.Done()
		oils := make([]*db.Oil, 0)
		err = db.Get().Where("id in (?)", oilIds).Find(&oils).Error
		if err != nil {
			fmt.Printf("获取油品信息失败：err: %+v", err)
		}
		for _, o := range oils {
			oilMap[o.Id] = o
		}
	})
	wg.Wait()

	results := make([]*models.Record, 0, len(records))
	for _, r := range records {
		results = append(results, RecordVo2Dto(r, stationMap, staffMap, oilMap))
	}

	return results, total, nil
}

func AddRecord(r *models.Record) error {
	if util.Str2Int64(r.StationId) == 0 {
		return errors.New("未指定加油站")
	}
	if util.Str2Int64(r.StaffId) == 0 {
		return errors.New("未指定加油员")
	}
	if util.Str2Int64(r.OilId) == 0 {
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
		wg      sync.WaitGroup
	)
	wg.Add(1)
	util.GoWithDefaultRecovery(func() {
		defer wg.Done()
		found, err := db.FindOne(&station, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", r.StationId).Where("is_delete = 0")
		})
		if err != nil {
			fmt.Printf("获取加油站信息失败, err: %+v", err)
			station = nil
		} else if !found {
			station = nil
		}
	})
	wg.Add(1)
	util.GoWithDefaultRecovery(func() {
		defer wg.Done()
		found, err := db.FindOne(&staff, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", r.StaffId).Where("is_delete = 0")
		})
		if err != nil {
			fmt.Printf("获取加油员信息失败, err: %+v", err)
			staff = nil
		} else if !found {
			staff = nil
		}
	})
	wg.Add(1)
	util.GoWithDefaultRecovery(func() {
		defer wg.Done()
		found, err := db.FindOne(&oil, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", r.OilId).Where("is_delete = 0")
		})
		if err != nil {
			fmt.Printf("获取油品信息失败, err: %+v", err)
			oil = nil
		} else if !found {
			oil = nil
		}
	})
	wg.Wait()

	if station == nil {
		return errors.New("指定加油站不存在")
	}
	if staff == nil {
		return errors.New("指定加油员不存在")
	}
	if oil == nil {
		return errors.New("指定油品不存在")
	}

	rVo := RecordDto2Vo(r)
	return db.Get().Create(rVo).Error
}

func UpdateRecord(param *models.UpdateRecrodParam) error {
	recordId := util.Str2Int64(param.RecordId)
	if recordId == 0 {
		return errors.New("指定加油记录不存在")
	}
	existRecord := &db.Record{}
	found, err := db.FindOne(&existRecord, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", recordId).Where("is_delete = 0")
	})
	if err != nil {
		fmt.Printf("获取加油记录失败, err: %+v", err)
		return errors.New("获取加油记录失败")
	}
	if !found {
		return errors.New("加油记录未找到")
	}
	if param.StationId != nil && util.Str2Int64(*param.StationId) > 0 {
		station := &db.Station{}
		found, err = db.FindOne(&station, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", util.Str2Int64(*param.StationId)).Where("is_delete = 0")
		})
		if err != nil {
			fmt.Printf("获取油站信息失败,err: %+v", err)
			return errors.New("获取油站信息失败")
		}
		if !found {
			return errors.New("油站信息未找到")
		}
		existRecord.StationId = util.Str2Int64(*param.StationId)
	}
	if param.OilId != nil && util.Str2Int64(*param.OilId) > 0 {
		oil := &db.Oil{}
		found, err = db.FindOne(&oil, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", util.Str2Int64(*param.OilId)).Where("is_delete = 0")
		})
		if err != nil {
			fmt.Printf("获取油品信息失败, err: %+v", err)
			return errors.New("获取油品信息失败")
		}
		if !found {
			return errors.New("油品信息未找到")
		}
		existRecord.OilId = util.Str2Int64(*param.OilId)
	}
	if param.Price != nil && *param.Price > 0 {
		existRecord.Price = *param.Price
	}
	if param.Liter != nil && *param.Liter > 0 {
		existRecord.Liter = *param.Liter
	}
	if param.Amount != nil && *param.Amount > 0 {
		existRecord.Amount = *param.Amount
	}
	if param.DriverName != nil && *param.DriverName != "" {
		existRecord.DriverName = *param.DriverName
	}
	if param.DriverPhone != nil && *param.DriverPhone != "" {
		existRecord.DriverPhone = *param.DriverPhone
	}
	return db.Get().Save(existRecord).Error
}

func DeleteRecord(recordIdStr string) error {
	recordId := util.Str2Int64(recordIdStr)
	if recordId == 0 {
		return errors.New("未指定加油记录")
	}
	updateFields := map[string]interface{}{
		"is_delete": 1,
	}
	return db.Get().Model(&db.Record{}).Where("id = ?", recordId).Updates(updateFields).Error
}
