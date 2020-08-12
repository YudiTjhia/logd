package apis

import (
	"applib/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"logd/client/go"
	"logd/env"
	"logd/models"
	"logd/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

var headers = map[string]string {
	"Content-Type": "application/json",
}

func TrunctLog (t *testing.T) {
	log := models.NewLog()
	err := log.Truncate()
	if err!=nil {
		t.Error(err)
	}
}

func LogApi_Setup() *mux.Router {

	env.SetEnv("../conf/env.json", "testing")
	r := mux.NewRouter()
	LogApi := LogApi{}
	LogApi.BindMethods(r)
	return r
}


func InsertLog_ApiTest(t *testing.T, r *mux.Router, expc *models.Log) *models.Log {

	TrunctLog(t)

	expectJson := util.ToJson(expc)
	LogReturn := &models.Log{}
	
	apitest.New().
		Debug().
		Handler(r).
		Post("/log").
		Headers(headers).
		JSON(expectJson).
		Expect(t).
		Status(http.StatusOK).
		Assert(func(response *http.Response, request *http.Request) error {

			actual, err := ResponseToLog(response.Body)
			if err != nil {
				return err
			}
			AssertLog(t, expc, actual)
			LogReturn = actual

			return nil

		}).
		End()

	return LogReturn
}


func Test_Insert_LogApi(t *testing.T) {

	r := LogApi_Setup()
	ts := httptest.NewServer(r)
	defer ts.Close()

	expect := CreateLog()
	InsertLog_ApiTest(t, r, expect)

}

func Test_Single_LogApi(t *testing.T) {

	r := LogApi_Setup()
	ts := httptest.NewServer(r)
	defer ts.Close()

	expect := CreateLog()
	actual := InsertLog_ApiTest(t, r, expect)

	apitest.New().
		Debug().
		Handler(r).
		Get("/log").
		Headers(headers).
		Query("LogID", actual.LogID).
		Expect(t).
		Assert(func(response *http.Response, request *http.Request) error {

			actual2, err := ResponseToLog(response.Body)
			if err != nil {
				return err
			}

			AssertLog(t, actual, actual2)

			return nil

		}).
		Status(http.StatusOK).
		End()		

}

func Test_List_LogApi(t *testing.T) {

	r := LogApi_Setup()
	ts := httptest.NewServer(r)
	defer ts.Close()

	expect := CreateLog()
	InsertLog_ApiTest(t, r, expect)

	apitest.New().
		Debug().
		Handler(r).
		Get("/log/list").
		Headers(headers).
		Expect(t).
		Assert(func(response *http.Response, request *http.Request) error {

			Log, err := ResponseToLogs(response.Body)
			if err != nil {
				return err
			}
			assert.Equal(t, 1,
				len(Log))
			return nil

		}).
		Status(http.StatusOK).
		End()

}

func CreateLog() *models.Log {

	log := models.Log{}
	log.LogTyp = _go.LOG_TYPE_ACCESS
	log.Domain = "localhost"
	log.IpAddr ="127.0.0.1"
	log.Port = 8080
	log.Acc = "123"
	log.Login = "sa"
	log.App = "logd"
	log.Serv = "LogServ"
	log.Methd = "CreateLog"
	log.Tag = "logd"
	log.Mesg = "hello logd"

	return &log
}

func Createlog(logServ services.LogServ, expc *models.Log) (*models.Log, error) {

	return logServ.CreateLog(
		expc.LogTyp,
		expc.Domain,
		expc.IpAddr,
		expc.Port,
		expc.Acc,
		expc.Login,
		expc.App,
		expc.Serv,
		expc.Methd,
		expc.Tag,
		expc.Mesg)
}

func CreateLogFunc(model *models.Log) *models.Log {
	logServ :=  services.LogServ{}
	log, err := Createlog(logServ, model)
	if err != nil {
		panic(err)
	}
	return log
}

func AssertLog(t *testing.T, expect *models.Log, actual *models.Log) {

	assert.NotEmpty(t, actual.LogID, "logID")
	assert.Equal(t, expect.LogTyp, actual.LogTyp)
	assert.Equal(t, expect.Domain, actual.Domain)
	assert.Equal(t, expect.IpAddr, actual.IpAddr)
	assert.Equal(t, expect.Port, actual.Port)
	assert.Equal(t, expect.Acc, actual.Acc)
	assert.Equal(t, expect.Login, actual.Login)
	assert.Equal(t, expect.App, actual.App)
	assert.Equal(t, expect.Serv, actual.Serv)
	assert.Equal(t, expect.Methd, actual.Methd)
	assert.Equal(t, expect.Tag, actual.Tag)
	assert.Equal(t, expect.Mesg, actual.Mesg)
}

func ResponseToLog(responseBody io.ReadCloser) (*models.Log, error) {

	body, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return nil, err
	}

	log := &models.Log{}
	err = json.Unmarshal(body, &log)
	if err != nil {
		return nil, err
	}

	return log, nil

}

func ResponseToLogs(responseBody io.ReadCloser) ([]models.Log, error) {

	body, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return []models.Log{}, err
	}
	log := []models.Log{}
	err = json.Unmarshal(body, &log)
	if err != nil {
		return []models.Log{}, err
	}
	return log, nil

}

