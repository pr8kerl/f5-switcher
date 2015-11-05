package main

import (
	"encoding/json"
	"io/ioutil"
)

var (
	cfg     Config
	cfgfile string = "config.json"
)

type Config struct {
	F5config  Device     `json:"f5"`
	Webconfig WebService `json:"webservice"`
	Groups    []Group    `json:"groups"`
}

type Device struct {
	Hostname string `json:"hostname"`
	Username string `json:"username"`
	Password string `json:"passwd"`
}

type WebService struct {
	BindAddress string `json:"address"`
	BindPort    int    `json:"port"`
}

type PoolGroup struct {
	Name  string `json:"name"`
	Blue  Pool   `json:"blue"`
	Green Pool   `json:"green"`
}
type Pool struct {
	Members []string `json:"members"`
	Active  int      `json:"active"`
}

type Group struct {
	Name  string      `json:"name"`
	Pools []PoolGroup `json:"pools"`
	State string
}

func InitialiseConfig(c string) (err error) {

	// read in json file
	dat, err := ioutil.ReadFile(c)
	if err != nil {
		return err
	}

	// convert json to config struct
	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		return err
	}

	return nil

}
