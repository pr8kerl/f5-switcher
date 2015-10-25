package F5

import (
	"encoding/json"
	"log"
	"strings"
	//	"github.com/kr/pretty"
)

type LBVirtualPolicy struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	FullPath  string `json:"fullPath"`
}

type LBVirtualPoliciesRef struct {
	Items []LBVirtualPolicy `json":items"`
}

type LBVirtualProfile struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	FullPath  string `json:"fullPath"`
	Context   string `json:"context"`
}

type LBVirtualPersistProfile struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	TmDefault string `json:"tmDefault"`
}

type LBVirtualProfileRef struct {
	Items []LBVirtualProfile `json":items"`
}

type LBVirtual struct {
	Name             string                    `json:"name"`
	FullPath         string                    `json:"fullPath"`
	Partition        string                    `json:"partition"`
	Destination      string                    `json:"destination"`
	Pool             string                    `json:"pool"`
	AddressStatus    string                    `json:"addressStatus"`
	AutoLastHop      string                    `json:"autoLasthop"`
	CmpEnabled       string                    `json:"cmpEnabled"`
	ConnectionLimit  int                       `json:"connectionLimit"`
	Enabled          bool                      `json:"enabled"`
	IpProtocol       string                    `json:"ipProtocol"`
	Source           string                    `json:"source"`
	SourcePort       string                    `json:"sourcePort"`
	SynCookieStatus  string                    `json:"synCookieStatus"`
	TranslateAddress string                    `json:"translateAddress"`
	TranslatePort    string                    `json:"translatePort"`
	Profiles         LBVirtualProfileRef       `json:"profilesReference"`
	Policies         LBVirtualPoliciesRef      `json:"policiesReference"`
	Rules            []string                  `json:"rules"`
	Persist          []LBVirtualPersistProfile `json:"persist"`
}

type LBVirtuals struct {
	Items []LBVirtual
}

func (f *Device) ShowVirtual(vname string) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/virtual/" + vname + "?expandSubcollections=true"
	res := LBVirtual{}

	err, resp := f.SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	f.PrintResponse(&res)

}

func (f *Device) UpdateVirtual(vname string) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/virtual/" + vname
	res := LBVirtual{}
	body := json.RawMessage{}

	// read in json file
/*
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a virtual struct
	err := json.Unmarshal(dat, &body)
	if err != nil {
		log.Fatal(err)
	}
*/

	// put the request
	err, resp := f.SendRequest(u, PUT, &sessn, &body, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	f.PrintResponse(&res)
}
