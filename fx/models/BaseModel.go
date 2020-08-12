package models

import "github.com/jinzhu/gorm"

const (
	RECORD_NOT_FOUND = "record not found"
)

type BaseModel struct {
	db *gorm.DB `json:"-" gorm:"-"`
	connection *gorm.DB `json:"-" gorm:"-"`
	table string `json:"-" gorm:"-"`
}

func(model *BaseModel) SetDb(db *gorm.DB, table string) {
	model.db = db
	model.table = table
	model.connection = db.Table(table)
}

func(model *BaseModel) Db() *gorm.DB {
	return model.db
}

func (model BaseModel) Table() string {
	return model.table
}

func(model BaseModel) Connection() *gorm.DB {
	if model.connection == nil {
		panic("Model is not initialized with New{$ModelName} function")
	}
	return model.connection
}

func(model BaseModel) Create(pModel interface{}) *gorm.DB {
	return model.Connection().Create(&pModel)
}
func(model BaseModel) Update(pModel interface{}) *gorm.DB  {
	return model.Connection().Save(&pModel)
}

func(model BaseModel) Delete(pModel interface{}, condition interface{}) *gorm.DB  {
	return model.Connection().Delete(pModel, condition)
}

func(model BaseModel) First(pModel interface{}, condition interface{}) *gorm.DB  {
	return model.Connection().Take(&pModel, condition)
}
