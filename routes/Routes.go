package routes

import (
	"applib/server"
	"logd/apis"
)

type Routes struct{}

func (routes *Routes) DefineRoutes(server *server.HttpServer) {

	logApi := apis.LogApi{}
	logApi.BindMethods(server.MuxRouter)

}
