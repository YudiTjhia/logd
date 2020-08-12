package main

import (
	"applib/server"
	"logd/env"
	"logd/routes"
)

func main() {

	envObj := env.SetEnv("./conf/env.json", "dev")
	httpServer := server.HttpServer{}
	httpServer.Create( *envObj.ServerConf, nil)

	router := routes.Routes{}
	router.DefineRoutes(&httpServer)

	httpServer.Listen()

}
