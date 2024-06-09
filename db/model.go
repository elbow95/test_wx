package db

import (
	"time"
)

// 用户表
type User struct {
	Id          int64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:主键id" json:"id"`
	OpenId      string    `gorm:"column:open_id;type:varchar(255);comment:微信open_id;NOT NULL" json:"open_id"`
	Name        string    `gorm:"column:name;type:varchar(255);comment:姓名;NOT NULL" json:"name"`
	Type        int       `gorm:"column:type;type:tinyint(4);default:0;comment:类型 1-管理员 10-加油员 20-司机;NOT NULL" json:"type"`
	StationId   int64     `gorm:"column:station_id;type:bigint(20) ;comment:所属油站id" json:"station_id"`
	Phone       string    `gorm:"column:phone;type:varchar(255);comment:手机号;NOT NULL" json:"phone"`
	PlateNumber string    `gorm:"column:plate_number;type:varchar(255);comment:车牌号;NOT NULL" json:"plate_number"`
	Avatar      string    `gorm:"column:avatar;type:varchar(255);comment:用户头像;NOT NULL" json:"avatar"`
	Extra       string    `gorm:"column:extra;type:text;comment:冗余信息" json:"extra"`
	Status      int       `gorm:"column:status;type:tinyint(4);default:0;comment:状态 1-正常 2-待验证;NOT NULL" json:"status"`
	IsDelete    int       `gorm:"column:is_delete;type:tinyint(4);default:0;comment:是否删除;NOT NULL" json:"is_delete"`
	CreateTime  time.Time `gorm:"->" json:"create_time"`
	UpdateTime  time.Time `gorm:"->" json:"update_time"`
}

func (m *User) TableName() string {
	return "user"
}

type UserExtra struct {
	License string `json:"license"`
	Company string `json:"company"`
}

// 油站表
type Station struct {
	Id         int64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:主键id" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(255);comment:油站名称;NOT NULL" json:"name"`
	Address    string    `gorm:"column:address;type:text;comment:油站位置" json:"address"`
	Longitude  float64   `gorm:"column:longitude;type:double;comment:经度" json:"longitude"`
	Latitude   float64   `gorm:"column:latitude;type:double;comment:维度" json:"latitude"`
	IsDelete   int       `gorm:"column:is_delete;type:tinyint(4);default:0;comment:是否删除;NOT NULL" json:"is_delete"`
	CreateTime time.Time `gorm:"->" json:"create_time"`
	UpdateTime time.Time `gorm:"->" json:"update_time"`
}

func (m *Station) TableName() string {
	return "station"
}

// 油品表
type Oil struct {
	Id         int64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:主键id" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(255);comment:油品名称;NOT NULL" json:"name"`
	StationId  int64     `gorm:"column:station_id;type:bigint(20) unsigned;default:0;comment:关联油站;NOT NULL" json:"station_id"`
	Price      int64     `gorm:"column:price;type:bigint(20) unsigned;default:0;comment:单价/分;NOT NULL" json:"price"`
	IsDelete   int       `gorm:"column:is_delete;type:tinyint(4);default:0;comment:是否删除;NOT NULL" json:"is_delete"`
	CreateTime time.Time `gorm:"->" json:"create_time"`
	UpdateTime time.Time `gorm:"->" json:"update_time"`
}

func (m *Oil) TableName() string {
	return "oil"
}

// 加油记录表
type Record struct {
	Id          int64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:主键id" json:"id"`
	StationId   int64     `gorm:"column:station_id;type:bigint(20) unsigned;default:0;comment:油站id;NOT NULL" json:"station_id"`
	OilId       int64     `gorm:"column:oil_id;type:bigint(20) unsigned;default:0;comment:油品id;NOT NULL" json:"oil_id"`
	StaffId     int64     `gorm:"column:staff_id;type:bigint(20) unsigned;default:0;comment:加油员id;NOT NULL" json:"staff_id"`
	Price       int64     `gorm:"column:price;type:bigint(20) unsigned;default:0;comment:单价/分;NOT NULL" json:"price"`
	Liter       int64     `gorm:"column:liter;type:bigint(20) unsigned;default:0;comment:数量/百分之一升;NOT NULL" json:"liter"`
	Amount      int64     `gorm:"column:amount;type:bigint(20) unsigned;default:0;comment:总价/分;NOT NULL" json:"amount"`
	DriverId    int64     `gorm:"column:driver_id;type:bigint(20) unsigned;default:0;comment:司机id;NOT NULL" json:"driver_id"`
	DriverName  string    `gorm:"column:driver_name;type:varchar(255);comment:司机姓名;NOT NULL" json:"driver_name"`
	DriverPhone string    `gorm:"column:driver_phone;type:varchar(255);comment:司机手机号;NOT NULL" json:"driver_phone"`
	IsDelete    int       `gorm:"column:is_delete;type:tinyint(4);default:0;comment:是否删除;NOT NULL" json:"is_delete"`
	CreateTime  time.Time `gorm:"->" json:"create_time"`
	UpdateTime  time.Time `gorm:"->" json:"update_time"`
}

func (m *Record) TableName() string {
	return "record"
}
