package service

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/models"
)

func ListUser(param *models.ListUserParam) ([]*models.User, error) {
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
		if param.Phone != "" {
			query = query.Where("phone like %?%", param.Phone)
		}
		if param.PlateNumber != "" {
			query = query.Where("plate_number like %?%", param.PlateNumber)
		}
	}
	query = query.Where("is_delete = 0").Order("id desc")
	users := make([]*model.User, 0)
	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}

	results := make([]*models.User, 0, len(users))
	for _, u := range users {
		results = append(results, UserVo2Dto(u))
	}

	return results, nil
}

func UserVo2Dto(u *model.User) *models.User {
	if u == nil {
		return nil
	}
	return &models.User{
		Id:          u.Id,
		Name:        u.Name,
		StationiId:  u.StationId,
		Type:        u.Type,
		Phone:       u.Phone,
		PlateNumber: u.PlateNumber,
		CreateTime:  u.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  u.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}

func UserDto2Vo(u *models.User) *model.User {
	if u == nil {
		return nil
	}
	return &model.User{
		Id:          u.Id,
		Name:        u.Name,
		Type:        u.Type,
		Phone:       u.Phone,
		PlateNumber: u.PlateNumber,
	}
}

func AddUser(u *models.User) error {
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
	return db.Get().Model(&model.User{}).Where("id = ?", userId).Updates(updateFields).Error
}
