package utils

import (
	"io/ioutil"
	"encoding/json"
	"os"
)

type GlobalConf struct {
	Name string
	IPVerion string
	IP string
	Port uint16
	Version string

	ConfFilePath string

	MaxPacketSize uint32
	WorkPoolSize uint16
	MaxQueueLength uint16
	MaxConn uint32
	MaxMsgChanLen uint32
}

var GlobalConfig *GlobalConf

func PathExist(path string) (exsit bool, err error) {
	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func (gc *GlobalConf) Reload() {

	if pathExist, _ := PathExist(gc.ConfFilePath); pathExist != true {
		return
	}

	data, err := ioutil.ReadFile(gc.ConfFilePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, gc)
	if err !=nil {
		panic(err)
	}

}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		pwd=""
	}

	GlobalConfig = &GlobalConf {
		Name: "Fortc Server",
		IPVerion: "IPV4",
		IP: "127.0.0.1",
		Port: 8888,

		Version: "V.01",
		ConfFilePath: pwd+"/conf/conf.json",

		MaxPacketSize:1024,
		WorkPoolSize:4,
		MaxQueueLength:1024,
		MaxConn:100000,
		MaxMsgChanLen: 1024,
	}

	GlobalConfig.Reload()
}