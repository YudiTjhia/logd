package env

import (
	applibconf "applib/conf"
	"logd/fx/conf"
)


type Env struct {
	AppConf *conf.AppConf
	EnvConf *conf.EnvConf
	DbConfs map[string]applibconf.DbConf
	ServerConf *applibconf.ServerConf

}

var _envMap map[string]*Env
var _currentEnv *Env
func GetEnv () *Env {
	if _currentEnv == nil {
		panic("Must Call SetEnvFile and SetEnv first")
	}
	return _currentEnv
}
func SetEnv(envFile string, envID string) *Env {
	if envFile == "" {
		panic("envID is empty")
	}
	if envID == "" {
		panic("envID is empty")
	}
	if _envMap == nil {
		_envMap = map[string]*Env{}
	}

	env, e:= _envMap[envID]
	if !e {
		env := &Env{}
		env.EnvConf = conf.GetEnvConf(envFile, envID)
		env.AppConf = conf.ReadAppConfFile(env.EnvConf.AppConfigFile)
		env.DbConfs = applibconf.GetDbConfigs( env.EnvConf.DbConfigFile )

		serverConf := applibconf.ServerConf{}
		serverConf.Load( env.EnvConf.ServerConfigFile)

		env.ServerConf = &serverConf
		_envMap[envID] = env
		_currentEnv = env

		return env
	}

	_currentEnv = env
	return env
}

