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
	F5config  Device           `json:"f5"`
	Webconfig WebService       `json:"webservice"`
	Groups    map[string]Group `json:"groups"`
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
	Blue  []string `json:"blue"`
	Green []string `json:"green"`
	State string   `json:"state"`
	Error string   `json:"error"`
}

func (t Pool) SetState(s string) Pool {
	t.State = s
	return t
}

func (t Pool) SetError(s string) Pool {
	t.Error = s
	return t
}

type Group struct {
	Pools map[string]Pool `json:"pools"`
	State string          `json:"state"`
}

func (t Group) SetState(s string) Group {
	t.State = s
	return t
}

type GroupPutData struct {
	Name  string `json:"name"`
	State string `json:"state"`
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
