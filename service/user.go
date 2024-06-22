package service

import (
	"errors"
	"fmt"
	"log"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/models"
	"wxcloudrun-golang/util"

	"github.com/tidwall/sjson"
	"gorm.io/gorm"
)

func ListUser(param *models.ListUserParam) ([]*models.User, int64, error) {
	query := db.Get()
	if param != nil {
		if len(param.UserIds) > 0 {
			userIds := make([]int64, 0)
			for _, i := range param.UserIds {
				userIds = append(userIds, util.Str2Int64(i))
			}
			query = query.Where("id in (?)", userIds)
		}
		if len(param.StationIds) > 0 {
			stationIds := make([]int64, 0)
			for _, i := range param.StationIds {
				stationIds = append(stationIds, util.Str2Int64(i))
			}
			query = query.Where("station_id in (?)", stationIds)
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
		fmt.Printf("查询用户失败,err:%v", err)
		return nil, 0, errors.New("查询用户失败")
	}
	users := make([]*db.User, 0)
	query = db.AttachPage(query, param.Page, param.PageSize)
	err = query.Find(&users).Error
	if err != nil {
		fmt.Printf("查询用户失败,err:%v", err)
		return nil, 0, errors.New("查询用户失败")
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
		Id:          util.Int642Str(u.Id),
		OpenId:      u.OpenId,
		Name:        u.Name,
		StationId:   util.Int642Str(u.StationId),
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
		Id:          util.Str2Int64(u.Id),
		Name:        u.Name,
		Type:        u.Type,
		StationId:   util.Str2Int64(u.StationId),
		Phone:       u.Phone,
		PlateNumber: u.PlateNumber,
		Extra:       util.MarshalJsonIgnoreError(&db.UserExtra{Company: u.Company, License: u.License}),
		Status:      u.Status,
	}
}

func AddUser(u *models.User) error {
	if u.Phone == "" {
		return errors.New("添加人员需要填写手机号")
	}
	user := &db.User{}
	found, err := db.FindOne(&user, func(db *gorm.DB) *gorm.DB {
		return db.Where("phone = ?", u.Phone).Where("type = ?", u.Type).Where("is_delete = 0")
	})
	if err != nil {
		log.Printf("[AddUser] get user failed, err: %+v", err)
		return errors.New("查找人员失败")
	}
	if found {
		return errors.New("人员已存在")
	}
	// 如果添加加油员，检查油站存不存在
	if u.Type == int(common.UserType_Staff) {
		if u.StationId == "" {
			return errors.New("添加加油员需要指定加油站")
		}
		station := &db.Station{}
		found, err = db.FindOne(&station, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", util.Str2Int64(u.StationId)).Where("is_delete = 0")
		})
		if err != nil {
			log.Printf("[AddUser] get station failed, err: %+v", err)
			return errors.New("查找加油站失败")
		} else if !found {
			return errors.New("指定的加油站不存在")
		}
	}

	uVo := UserDto2Vo(u)
	return db.Get().Create(uVo).Error
}

func UpdateUser(param *models.UpdateUserParam) error {
	userId := util.Str2Int64(param.UserId)
	if userId == 0 {
		return errors.New("未指定更新用户")
	}
	existUser := &db.User{}
	found, err := db.FindOneWithWrite(&existUser, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", userId).Where("is_delete = 0")
	})
	if err != nil {
		fmt.Printf("err: %v", err)
		return errors.New("用户信息查询失败")
	}
	if !found {
		return errors.New("用户信息未找到")
	}
	if param.Name != nil && *param.Name != "" {
		existUser.Name = *param.Name
	}
	if param.StationId != nil && util.Str2Int64(*param.StationId) != 0 {
		existUser.StationId = util.Str2Int64(*param.StationId)
	}
	if param.Phone != nil && *param.Phone != "" {
		existUser.Phone = *param.Phone
	}
	if param.PlateNumber != nil && *param.PlateNumber != "" {
		existUser.PlateNumber = *param.PlateNumber
	}
	if param.Company != nil && *param.Company != "" {
		existUser.Extra, _ = sjson.Set(existUser.Extra, "company", *param.Company)
	}
	if param.License != nil && *param.License != "" {
		existUser.Extra, _ = sjson.Set(existUser.Extra, "license", *param.License)
	}

	return db.Get().Save(existUser).Error
}

func DeleteUser(userIdStr string) error {
	userId := util.Str2Int64(userIdStr)
	if userId == 0 {
		return errors.New("未指定用户")
	}
	updateFields := map[string]interface{}{
		"is_delete": 1,
	}
	return db.Get().Model(&db.User{}).Where("id = ?", userId).Updates(updateFields).Error
}
