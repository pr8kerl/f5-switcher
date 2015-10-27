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

type Pool struct {
	Name  string   `json:"name"`
	Blue  []string `json:"blue"`
	Green []string `json:"green"`
}

type Group struct {
	Name  string `json:"name"`
	Pools []Pool `json:"pools"`
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
