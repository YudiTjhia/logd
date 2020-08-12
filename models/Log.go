package models

import (
	"logd/db"
	"logd/fx/models"
	"time"
)

type Log struct {

	models.BaseModel

	LogID  string `json:"logID" gorm:"Column:log_id;Type:varchar(40);primary_key;not null"`
	LogTyp string `json:"logTyp" gorm:"Column:log_typ;Type:varchar(10);not null"`
	Domain string `json:"domain" gorm:"Column:domain;Type:varchar(100);not null"`
	IpAddr string `json:"ipAddr" gorm:"Column:ip_addr;Type:varchar(40);not null"`
	Port   int    `json:"port" gorm:"Column:port;Type:integer;not null"`
	Acc    string `json:"acc" gorm:"Column:acc;Type:varchar(100);null"`
	Login  string `json:"login" gorm:"Column:login;Type:varchar(100);not null"`
	App    string `json:"app" gorm:"Column:app;Type:varchar(100);not null"`
	Serv   string `json:"serv" gorm:"Column:serv;Type:varchar(255);null"`
	Methd  string `json:"methd" gorm:"Column:methd;Type:varchar(255);null"`
	Tag    string `json:"tag" gorm:"Column:tag;Type:varchar(100);null"`
	Mesg  string `json:"mesg" gorm:"Column:mesg;Type:text;null"`
	CreatedAt time.Time  `json:"createdAt" gorm:"Column:created_at;Type:timestamptz;not null"`

}

func (model Log) Migrate() {

	if model.Connection().HasTable(model) {
		model.Connection().DropTable(model)
	}

	model.Connection().CreateTable(model)

	model.Connection().AddIndex("idx_log_acc", "acc")
	model.Connection().AddIndex("idx_log_domain_port", "domain","port")
	model.Connection().AddIndex("idx_log_ip_addr_port", "ip_addr", "port")
	model.Connection().AddIndex("idx_log_login", "login")
	model.Connection().AddIndex("idx_log_app", "app")
	model.Connection().AddIndex("idx_log_serv", "serv")
	model.Connection().AddIndex("idx_log_methd", "methd")
	model.Connection().AddIndex("idx_log_tag", "tag")
	model.Connection().AddIndex("idx_log_created_at", "acc", "created_at")

}


func(model Log) Truncate() error  {
	db := model.Connection().Exec("truncate table \"" + model.Table() + "\"")
	return db.Error
}


func NewLog() *Log {
	log := &Log{}
	log.SetDb( db.GetDbConnection(db.Name()),"log")
	return log
}
