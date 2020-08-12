package services

import (
	"logd/client/go"
	"logd/env"
	testing2 "logd/fx/testing"
	"logd/models"
	"testing"
)

func TrunctLog (t *testing.T) {
	log := models.NewLog()
	err := log.Truncate()
	if err!=nil {
		t.Error(err)
	}
}

func CreateLog (t *testing.T) *models.Log {

	logServ := LogServ{}
	log, err := logServ.CreateLog(_go.LOG_TYPE_ACCESS,
		"localhost",
		"127.0.0.1",
		8080,
		"123",
		"sa",
		"logd",
		"LogServ",
		"CreateLog",
		"logd",
		"hello to logd")
	if err!=nil {
		t.Error(err)
	}
	return log
}
func TestLogServ_CreateLog(t *testing.T) {

	env.SetEnv("../conf/env.json", "migration")
	TrunctLog(t)
	CreateLog(t)
}

func assertLog(t *testing.T, expc *models.Log, actl *models.Log) {
	testing2.AssertEquals(t, expc.LogTyp, actl.LogTyp)
	testing2.AssertEquals(t, expc.Domain, actl.Domain)
	testing2.AssertEquals(t, expc.IpAddr, actl.IpAddr)
	testing2.AssertEquals(t, expc.Port, actl.Port)
	testing2.AssertEquals(t, expc.Acc, actl.Acc)
	testing2.AssertEquals(t, expc.Login, actl.Login)
	testing2.AssertEquals(t, expc.App, actl.App)
	testing2.AssertEquals(t, expc.Serv, actl.Serv)
	testing2.AssertEquals(t, expc.Methd, actl.Methd)
	testing2.AssertEquals(t, expc.Tag, actl.Tag)
	testing2.AssertEquals(t, expc.Mesg, actl.Mesg)
}

func TestLogServ_FindLog(t *testing.T) {

	env.SetEnv("../conf/env.json", "migration")
	TrunctLog(t)
	creatdLog := CreateLog(t)

	logServ := LogServ{}
	findLog,err  := logServ.FindLog(creatdLog.LogID)
	if err!=nil {
		t.Error(err)
	}

	assertLog(t, creatdLog, findLog)

}
