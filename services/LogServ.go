package services

import (
	"github.com/google/uuid"
	"logd/fx/svc"
	"logd/models"
)

type LogServ struct {
	 svc.BaseService
}

func (serv LogServ) CreateLog(logType string,
	domain string,
	ipAddr string,
	port int,
	acc string,
	login string,
	app string,
	servc string,
	methd string,
	tag string,
	mesg string) (*models.Log, error) {

	log:=models.NewLog()
	log.LogID = uuid.New().String()
	log.LogTyp = logType
	log.Domain = domain
	log.IpAddr = ipAddr
	log.Port = port
	log.Acc = acc
	log.Login = login
	log.App = app
	log.Serv = servc
	log.Methd = methd
	log.Tag= tag
	log.Mesg = mesg
	db := log.Connection().Create(&log)
	if db.Error!=nil {
		return nil, db.Error
	}

	return log, nil
}

func (serv LogServ) FindLogs(logType string,
	domain string,
	ipAddr string,
	port int,
	acc string,
	login string,
	app string,
	servc string,
	methd string,
	tag string) ([]models.Log, error) {

	log:=models.NewLog()
	log.LogTyp = logType
	log.Domain = domain
	log.IpAddr = ipAddr
	log.Port = port
	log.Acc = acc
	log.Login = login
	log.App = app
	log.Serv = servc
	log.Methd = methd
	log.Tag= tag

	var logs []models.Log
	db := log.Connection().Find(&logs, log)
	if db.Error!=nil {
		return nil, db.Error
	}
	return logs, nil
}

func (serv LogServ) FindLog(logID string) (*models.Log, error) {

	log:=models.NewLog()
	log.LogID = logID

	db := log.Connection().Take(&log)
	if db.Error!=nil {
		return nil, db.Error
	}

	return log, nil
}
