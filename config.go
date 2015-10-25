package main

import (
	"encoding/json"
	"io/ioutil"
)

var (
	config  Config
	cfgfile string = "config.json"
)

type Config struct {
	F5config    F5 `json:"f5"`
	Webconfig   WebService `json:"webservice"`
	Groups      []Group  `json:"groups"`
}

type WebService struct {
	BindAddress string   `json:"address"`
	BindPort    int      `json:"port"`
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

func InitialiseConfig(cfg string) (err error) {

	// read in json file
	dat, err := ioutil.ReadFile(cfg)
	if err != nil {
		return err
	}

	// convert json to config struct
	err = json.Unmarshal(dat, &config)
	if err != nil {
		return err
	}

	return nil
}
