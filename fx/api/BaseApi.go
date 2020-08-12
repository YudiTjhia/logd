package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type BaseApi struct {
	ApiModel
	Header *HeaderModel
}

func (api *BaseApi) ParseHeaders(r *http.Request) {
	api.Header = &HeaderModel{}
	api.Header.AccountID = r.Header.Get("Account-ID")
	api.Header.LoginID = r.Header.Get("Login-ID")
	//api.Header.AccountID = "123"
	//api.Header.LoginID = "sa"
}

func (api *BaseApi) ParseQuery(r *http.Request) {
	api.Q = api.DecodeQueryKey(r, "q")
	api.UsePaging = api.DecodeQueryBool(r, "usePaging")
	api.Page = api.DecodeQueryInt(r, "page")
	api.PageSize = api.DecodeQueryInt(r, "pageSize")
}

func (api BaseApi) OK(w http.ResponseWriter, data interface{}) {
	if data != nil {
		bytesEnt, err := json.Marshal(data)
		//fmt.Println(string(pretty.Pretty(bytesEnt)))
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(bytesEnt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (api BaseApi) Error(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}

func (api BaseApi) DecodeQueryKey(r *http.Request, key string) string {
	val := r.URL.Query()[key]
	if val != nil {
		return val[0]
	}
	return ""
}

func (api BaseApi) DecodeQueryInt(r *http.Request, field string) int{
	val := api.DecodeQueryKey(r, field)
	if val!="" {
		intVal, _ := strconv.Atoi(val)
		return intVal
	}
	return 0
}

func (api BaseApi) DecodeQueryBool(r *http.Request, field string) bool{
	val := api.DecodeQueryKey(r, field)
	if val!="" {
		boolVal, _ := strconv.ParseBool(val)
		return boolVal
	}
	return false
}

func (api BaseApi) DecodeQueryFloat32(r *http.Request, field string) float32 {
	val := api.DecodeQueryKey(r, field)
	if val!="" {
		intVal, _ := strconv.ParseFloat(val,32)
		return float32(intVal)
	}
	return float32(0)
}

func (api BaseApi) DecodeQueryFloat64(r *http.Request, field string) float64 {
	val := api.DecodeQueryKey(r, field)
	if val!="" {
		intVal, _ := strconv.ParseFloat(val,64)
		return float64(intVal)
	}
	return float64(0)
}