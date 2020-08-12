package _go

import "testing"

func TestLogc_Log(t *testing.T) {

	logc := NewLogC("http://127.0.0.1:6767",
		"client.com",
		"127.0.0.1",
		8080,
		"logd")

	respBytes, err := logc.Accs("123",
		"sa",
		"logServ",
		"CreateLog",
		"logd",
		"hello logd")

	if err!=nil {
		t.Error(err)
	}
	t.Log(string(respBytes))

}


func TestLogc_FindLogs(t *testing.T) {

	logc := NewLogC("http://127.0.0.1:6767",
		"client.com",
		"127.0.0.1",
		8080,
		"logd")

	respBytes, err := logc.FindLogs(LOG_TYPE_ACCESS,
		"123",
		"sa",
		"logServ",
		"CreateLog",
		"logd")

	if err!=nil {
		t.Error(err)
	}
	t.Log(string(respBytes))
}


func TestLogc_FindLog(t *testing.T) {

	logc := NewLogC("http://127.0.0.1:6767",
		"client.com",
		"127.0.0.1",
		8080,
		"logd")

	respBytes, err := logc.FindLog("a89f98b9-ea09-40fb-a973-b1f3bd85c6ba")

	if err!=nil {
		t.Error(err)
	}
	t.Log(string(respBytes))
}
