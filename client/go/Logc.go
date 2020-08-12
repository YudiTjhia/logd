package _go

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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

	return &logc
}

func (logc Logc) Accs(acc string, login string,
	serv string, methd string,
	tag string, mesg interface{}) ([]byte, error) {

	return logc.Log(LOG_TYPE_ACCESS, acc, login, serv, methd,
		tag, mesg)
}

func (logc Logc) Err(acc string, login string,
	serv string, methd string,
	tag string, mesg interface{}) ([]byte, error) {

	return logc.Log(LOG_TYPE_ERR, acc, login, serv, methd,
		tag, mesg)
}

func (logc Logc) Data(acc string, login string,
	serv string, methd string,
	tag string, mesg interface{}) ([]byte, error) {

	return logc.Log(LOG_TYPE_DATA, acc, login, serv, methd,
		tag, mesg)

}

func (logc Logc) Log(logTyp string, acc string, login string,
	serv string, methd string,
	tag string, mesg interface{}) ([]byte, error) {

	respBytes := []byte{}

	var payload *bytes.Buffer =  new(bytes.Buffer)
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
		"mesg":    mesg,
	}


	err := json.NewEncoder(payload).Encode(mapData)
	if err != nil {
		return respBytes, err
	}

	request, err := http.NewRequest("POST", logc.server + "/log", payload)
	if err != nil {
		return respBytes, err
	}

	return logc.sendRequest(request)

}


func (logc Logc) FindLogs(logTyp string, acc string, login string,
	serv string, methd string,
	tag string) ([]byte, error) {

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

	err := json.NewEncoder(payload).Encode(mapData)
	if err != nil {
		return respBytes, err
	}

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
	if resp.StatusCode == http.StatusOK {
		return respBody, nil

	} else {
		strRespBody := string(respBody)
		return []byte{}, errors.New(strRespBody)
	}
}

func (logc Logc) FindLog(logID string) ([]byte, error) {

	respBytes := []byte{}
	var payload *bytes.Buffer = new(bytes.Buffer)
	mapData := map[string]interface{}{
		"logID": logID,
	}

	err := json.NewEncoder(payload).Encode(mapData)
	if err != nil {
		return respBytes, err
	}

	request, err := http.NewRequest("GET", logc.server + "/log", payload)
	if err != nil {
		return respBytes, err
	}

	return logc.sendRequest(request)

}
