package service

import (
	"errors"
	"log"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/models"
)

func ListUser(param *models.ListUserParam) ([]*models.User, int64, error) {
	query := db.Get()
	if param != nil {
		if len(param.Ids) > 0 {
			query = query.Where("id in (?)", param.Ids)
		}
		if len(param.StationIds) > 0 {
			query = query.Where("station_id in (?)", param.StationIds)
		}
		if len(param.Types) > 0 {
			query = query.Where("type in (?)", param.Types)
		}
		if param.Name != "" {
			query = query.Where("name like %?%", param.Name)
		}
		if param.Phone != "" {
			query = query.Where("phone like %?%", param.Phone)
		}
		if param.PlateNumber != "" {
			query = query.Where("plate_number like %?%", param.PlateNumber)
		}
	}

	query = query.Where("is_delete = 0").Order("id desc")
	var (
		total int64
	)
	err := query.Model(&db.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	users := make([]*db.User, 0)
	query = db.AttachPage(query, param.Page, param.PageSize)
	err = query.Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	results := make([]*models.User, 0, len(users))
	for _, u := range users {
		results = append(results, UserVo2Dto(u, nil))
	}

	return results, total, nil
}

func UserVo2Dto(u *db.User, stationMap map[int64]*db.Station) *models.User {
	if u == nil {
		return nil
	}
	user := &models.User{
		Id:          u.Id,
		OpenId:      u.OpenId,
		Name:        u.Name,
		StationiId:  u.StationId,
		Type:        u.Type,
		Phone:       u.Phone,
		PlateNumber: u.PlateNumber,
		Avatar:      u.Avatar,
		Status:      u.Status,
		CreateTime:  u.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  u.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	if station, ok := stationMap[u.StationId]; ok {
		user.StationName = station.Name
	}
	user.PermissionList = common.UserPermissionMap[common.UserType(user.Type)]
	return user
}

func UserDto2Vo(u *models.User) *db.User {
	if u == nil {
		return nil
	}
	return &db.User{
		Id:          u.Id,
		Name:        u.Name,
		Type:        u.Type,
		Phone:       u.Phone,
		PlateNumber: u.PlateNumber,
	}
}

func AddUser(u *models.User) error {
	// 如果添加加油员，检查油站存不存在
	if u.Type == int(common.UserType_Staff) {
		if u.StationiId == 0 {
			return errors.New("添加加油员需要指定加油站")
		}
		stations := make([]*db.Station, 0)
		if err := db.Get().Where("id = ?", u.StationiId).Find(&stations).Error; err != nil {
			log.Printf("[AddUser] get station failed, err: %+v", err)
			return errors.New("查找加油站失败")
		} else if len(stations) == 0 {
			return errors.New("指定的加油站不存在")
		}
	}

	uVo := UserDto2Vo(u)
	return db.Get().Create(uVo).Error
}

func UpdateUser(u *models.User) error {
	uVo := UserDto2Vo(u)
	return db.Get().Save(uVo).Error
}

func DeleteUser(userId int64) error {
	updateFields := map[string]interface{}{
		"is_delete": 1,
	}
	return db.Get().Model(&db.User{}).Where("id = ?", userId).Updates(updateFields).Error
}
