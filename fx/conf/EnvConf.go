package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type EnvConf struct {
	EnvID string `json:"envID"`
	AppConfigFile string `json:"appConfigFile"`
	DbConfigFile string `json:"dbConfigFile"`
	ServerConfigFile string `json:"serverConfigFile"`
}

var _envConfs map[string]EnvConf
var _envConfMutex = sync.RWMutex{}

func GetEnvConf(envConfFile string, envID string) *EnvConf {
	if _envConfs == nil {
		readEnvConfFile(envConfFile)
	}
	envConf, e := _envConfs[envID]
	if !e {
		panic("cannot_find_envID=" + envID)
	}
	return &envConf
}

func readEnvConfFile(envConfFile string) {
	fopen, err := os.Open(envConfFile)
	if err != nil {
		panic(err)
	}
	bytesVal, err := ioutil.ReadAll(fopen)
	if err != nil {
		panic(err)
	}

	envConfs := []EnvConf{}
	err = json.Unmarshal(bytesVal, &envConfs)
	if err != nil {
		panic(err)
	}

	if len(envConfs) == 0 {
		panic("No Environments defined")
	}

	_envConfMutex.Lock()
	_envConfs = map[string]EnvConf{}
	for _, envConf := range envConfs {
		_, e := _envConfs[envConf.EnvID]
		if !e {
			_envConfs[envConf.EnvID] = envConf
		} else {
			panic("Duplicate envConf.EnvID =" + envConf.EnvID)
		}
	}
	_envConfMutex.Unlock()

}

