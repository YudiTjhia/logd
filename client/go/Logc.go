package _go

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	LOG_TYPE_ACCESS = "access"
	LOG_TYPE_DATA   = "data"
	LOG_TYPE_ERR    = "error"
)

type Logc struct {
	server    string
	cliDomain string
	cliIpAddr string
	cliPort   int
	cliApp    string
	enable bool
}

func NewLogC(server string, cliDomain string,
	cliIpAddr string,
	cliPort int,
	cliApp string) *Logc {

	logc := Logc{}
	logc.server = server
	logc.cliDomain = cliDomain
	logc.cliIpAddr = cliIpAddr
	logc.cliPort = cliPort
	logc.cliApp = cliApp
	logc.enable = true

	return &logc
}

func (logc *Logc) Enable() {
	logc.enable = true
}

func (logc *Logc) Disable() {
	logc.enable = false
}

func (logc Logc) Accs(acc string, login string,
	serv string, methd string,
	tag string, mesg ...interface{}) ([]byte, error) {

	log.Println("logc.Accs>")

	return logc.Log(LOG_TYPE_ACCESS, acc, login, serv, methd,
		tag, mesg)
}

func (logc Logc) Err(acc string, login string,
	serv string, methd string,
	tag string, mesg ...interface{}) ([]byte, error) {

	log.Println("logc.Err>")

	return logc.Log(LOG_TYPE_ERR, acc, login, serv, methd,
		tag, mesg)
}

func (logc Logc) Data(acc string, login string,
	serv string, methd string,
	tag string, mesg ...interface{}) ([]byte, error) {

	log.Println("logc.Data>")

	return logc.Log(LOG_TYPE_DATA, acc, login, serv, methd,
		tag, mesg)

}

func (logc Logc) Log(logTyp string, acc string, login string,
	serv string, methd string,
	tag string, mesg ...interface{}) ([]byte, error) {

	log.Println("logc.Log>")

	if logc.enable == false {
		log.Println("logc.enable=", logc.enable)
		return []byte{}, nil
	}


	respBytes := []byte{}
	var payload *bytes.Buffer =  new(bytes.Buffer)


	byteMesg,err := json.Marshal(mesg)
	if err!=nil {
		return []byte{}, err
	}

	strMesg := string(byteMesg)

	mapData := map[string]interface{}{
		"logTyp": logTyp,
		"domain":  logc.cliDomain,
		"ipAddr":  logc.cliIpAddr,
		"port":    logc.cliPort,
		"acc":     acc,
		"login":   login,
		"app":     logc.cliApp,
		"serv":    serv,
		"methd":   methd,
		"tag":     tag,
		"mesg":    strMesg,
	}

	log.Println("mapData=", mapData)
	err = json.NewEncoder(payload).Encode(mapData)
	if err != nil {
		return respBytes, err
	}

	request, err := http.NewRequest("POST", logc.server + "/log", payload)
	log.Println("method=POST, url=", logc.server + "/log, payload=", payload)

	if err != nil {
		return respBytes, err
	}

	return logc.sendRequest(request)
}


func (logc Logc) FindLogs(logTyp string, acc string, login string,
	serv string, methd string,
	tag string) ([]byte, error) {


	log.Println("logc.FindLogs>")
	log.Println("logc.FindLogs>logTyp=", logTyp)
	log.Println("logc.FindLogs>acc=", acc)
	log.Println("logc.FindLogs>login=", login)
	log.Println("logc.FindLogs>serv=", serv)
	log.Println("logc.FindLogs>methd=", methd)
	log.Println("logc.FindLogs>tag=", tag)

	respBytes := []byte{}
	var payload *bytes.Buffer = new(bytes.Buffer)
	mapData := map[string]interface{}{
		"logTyp": logTyp,
		"domain":  logc.cliDomain,
		"ipAddr":  logc.cliIpAddr,
		"port":    logc.cliPort,
		"acc":     acc,
		"login":   login,
		"app":     logc.cliApp,
		"serv":    serv,
		"methd":   methd,
		"tag":     tag,
	}

	log.Println("mapData=", mapData)
	err := json.NewEncoder(payload).Encode(mapData)
	if err != nil {
		return respBytes, err
	}


	log.Println("method=GET, url=", logc.server + "/logs, payload=", payload)
	request, err := http.NewRequest("GET", logc.server + "/logs", payload)
	if err != nil {
		return respBytes, err
	}

	return logc.sendRequest(request)
}


func(logc Logc) sendRequest(request *http.Request) ([]byte, error) {

	httpClient := new(http.Client)
	resp, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	strRespBody := string(respBody)

	log.Println("resp.StatusCode=", resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		log.Println("respBody=", strRespBody[0:50])
		return respBody, nil

	} else {
		log.Println("respErr=", strRespBody[0:50])
		return []byte{}, errors.New(strRespBody)
	}
}

func (logc Logc) FindLog(logID string) ([]byte, error) {

	log.Println("logc.FindLogs>")
	log.Println("logc.FindLogs>logID=", logID)

	respBytes := []byte{}
	var payload *bytes.Buffer = new(bytes.Buffer)
	mapData := map[string]interface{}{
		"logID": logID,
	}

	log.Println("mapData=", mapData)

	err := json.NewEncoder(payload).Encode(mapData)
	if err != nil {
		return respBytes, err
	}

	log.Println("method=GET, url=", logc.server + "/log, payload=", payload)

	request, err := http.NewRequest("GET", logc.server + "/log", payload)
	if err != nil {
		return respBytes, err
	}

	return logc.sendRequest(request)

}
