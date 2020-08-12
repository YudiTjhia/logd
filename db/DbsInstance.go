package db

import (
	"github.com/jinzhu/gorm"
	"logd/env"
	fxdb "logd/fx/db"
)

const (
	DB_NAME = "logd"
)

func Name() string {
	return DB_NAME
}

var _gormInstance *fxdb.GormDb

func GetDbConnection(connectionID string) *gorm.DB {
	if _gormInstance == nil {
		_gormInstance = &fxdb.GormDb{}
		_gormInstance.Init( env.GetEnv().DbConfs )
	}
	return _gormInstance.GetConnection(connectionID)
}



