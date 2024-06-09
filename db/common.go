package db

import (
	"errors"

	"gorm.io/gorm"
)

type WhereFunc func(db *gorm.DB) *gorm.DB

func combineWhereFuncs(db *gorm.DB, whereFuncs ...WhereFunc) *gorm.DB {
	for _, whereFunc := range whereFuncs {
		db = whereFunc(db)
	}
	return db
}

func FindOneWithDb(db *gorm.DB, ret interface{}, whereFuncs ...WhereFunc) (bool, error) {
	if len(whereFuncs) == 0 {
		return false, errors.New("no where func")
	}
	dbRes := combineWhereFuncs(db, whereFuncs...).Model(ret).First(ret)
	if dbRes.Error != nil {
		if errors.Is(gorm.ErrRecordNotFound, dbRes.Error) {
			return false, nil
		} else {
			return false, dbRes.Error
		}
	}
	return true, nil
}

func FindDataWithDb(db *gorm.DB, ret interface{}, whereFuncs ...WhereFunc) (bool, error) {
	if len(whereFuncs) == 0 {
		return false, errors.New("no where func")
	}
	dbRes := combineWhereFuncs(db, whereFuncs...).Model(ret).Find(ret)
	if dbRes.Error != nil {
		return false, dbRes.Error
	}
	return true, nil
}

func FindOne(ret interface{}, whereFuncs ...WhereFunc) (bool, error) {
	db := Get()
	return FindOneWithDb(db, ret, whereFuncs...)
}

func FindOneWithWrite(ret interface{}, whereFuncs ...WhereFunc) (bool, error) {
	db := GetWrite()
	return FindOneWithDb(db, ret, whereFuncs...)
}

func FindData(ret interface{}, whereFuncs ...WhereFunc) (bool, error) {
	db := Get()
	return FindDataWithDb(db, ret, whereFuncs...)
}

func FindDataWithWrite(ret interface{}, whereFuncs ...WhereFunc) (bool, error) {
	db := GetWrite()
	return FindDataWithDb(db, ret, whereFuncs...)
}

func AttachPage(db *gorm.DB, page, pageSize int32) *gorm.DB {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return db.Offset(int(pageSize * (page - 1))).Limit(int(pageSize))
}
