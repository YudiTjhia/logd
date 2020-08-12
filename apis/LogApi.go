package apis

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"logd/fx/api"
	"logd/models"
	"logd/services"
	"net/http"
)

type LogApi struct {
	api.BaseApi
	models.Log
}

func (objApi *LogApi) BindMethods(router *mux.Router) {

	router.HandleFunc("/log", objApi.Insert).Methods("POST")
	router.HandleFunc("/logs", objApi.List).Methods("GET")
	router.HandleFunc("/log", objApi.Single).Methods("GET")

}

func (api LogApi) Insert(w http.ResponseWriter, r *http.Request) {

	api.ParseHeaders(r)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&api)

	if err == nil {
		logServ := services.LogServ{}
		log, err := logServ.CreateLog(api.LogTyp,
			api.Domain,
			api.IpAddr,
			api.Port,
			api.Acc,
			api.Login,
			api.App,
			api.Serv,
			api.Methd,
			api.Tag,
			api.Mesg)
		if err == nil {
			api.OK(w, log)
		} else {
			api.Error(w, err)
		}
	} else {
		api.Error(w, err)
	}

}

func (api *LogApi) Single(w http.ResponseWriter, r *http.Request) {

	api.ParseHeaders(r)
	api.ParseQuery(r)
	api.LogID = api.DecodeQueryKey(r, "logID")
	logServ := services.LogServ{}
	log, err := logServ.FindLog(api.LogID)

	if err == nil {
		api.OK(w, log)
	} else {
		api.Error(w, err)
	}
}

func (api LogApi) List(w http.ResponseWriter, r *http.Request) {

	api.ParseHeaders(r)
	api.ParseQuery(r)

	logServ := services.LogServ{}
	logs, err := logServ.FindLogs(api.LogTyp,
		api.Domain,
		api.IpAddr,
		api.Port,
		api.Acc,
		api.Login,
		api.App,
		api.Serv,
		api.Methd,
		api.Tag)

	if err == nil {
		api.OK(w, logs)
	} else {
		api.Error(w, err)
	}

}