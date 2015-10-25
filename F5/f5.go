package F5

import (
	"crypto/tls"
	"encoding/json"
	"github.com/jmcvetta/napping"
	"log"
	"net/http"
	"net/url"
	"strings"
	"errors"
)

var (
	sessn   napping.Session
	tsport  http.Transport
	clnt    http.Client
	headers http.Header
	debug   bool
)

const (
	GET = iota
	POST
	PUT
	PATCH
	DELETE
)

type httperr struct {
	Message string
	Errors  []struct {
		Resource string
		Field    string
		Code     string
	}
}

type Device struct {
	Hostname string
	Username string
	Password string
}

func New(host string, username string, pwd string) *Device {
	f := Device{Hostname: host, Username: username, Password: pwd}
	f.InitSession()
	return &f
}

func (f *Device) InitSession() {

	// REST connection setup
	tsport = http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	clnt = http.Client{Transport: &tsport}
	headers = make(http.Header)

	//
	// Setup HTTP Basic auth for this session (ONLY use this with SSL).  Auth
	// can also be configured on a per-request basis when using Send().
	//
	sessn = napping.Session{
		Client:   &clnt,
		Log:      debug,
		Userinfo: url.UserPassword(f.Username, f.Password),
		Header:   &headers,
	}

}

func (f *Device) GetVirtual(vname string) (error, *LBVirtual) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/virtual/" + vname + "?expandSubcollections=true"
	res := LBVirtual{}

	err, resp := f.SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	return nil, &res

}

func (f *Device) SendRequest(u string, method int, sess *napping.Session, pload interface{}, res interface{}) (error, *napping.Response) {

	//
	// Send request to server
	//
	e := httperr{}
	var (
		err  error
		resp *napping.Response
	)
	sess.Log = debug

	switch method {
	case GET:
		resp, err = sess.Get(u, nil, &res, &e)
	case POST:
		resp, err = sess.Post(u, &pload, &res, &e)
	case PUT:
		resp, err = sess.Put(u, &pload, &res, &e)
	case PATCH:
		resp, err = sess.Patch(u, &pload, &res, &e)
	case DELETE:
		resp, err = sess.Delete(u, nil, &res, &e)
	}

	if err != nil {
		return err, resp
	}
	if resp.Status() == 401 {
		return errors.New("unauthorised - check your username and passwd"), resp
	}
	if resp.Status() >= 300 {
		return errors.New(e.Message), resp
	} else {
		// all is good in the world
		return nil, resp
	}
}

func (f *Device) PrintResponse(input interface{}) {

	jsonresp, err := json.MarshalIndent(&input, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(jsonresp))

}
