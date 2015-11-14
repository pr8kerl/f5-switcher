package F5

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/jmcvetta/napping"
	"log"
	"net/http"
	"net/url"
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

type Response struct {
	Status  int
	Message string
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

func (f *Device) SendRequest(u string, method int, sess *napping.Session, pload interface{}, res interface{}) (error, *Response) {

	//
	// Send request to server
	//
	e := httperr{}
	var (
		err   error
		nresp *napping.Response
	)
	sess.Log = debug

	switch method {
	case GET:
		nresp, err = sess.Get(u, nil, &res, &e)
	case POST:
		nresp, err = sess.Post(u, &pload, &res, &e)
	case PUT:
		nresp, err = sess.Put(u, &pload, &res, &e)
	case PATCH:
		nresp, err = sess.Patch(u, &pload, &res, &e)
	case DELETE:
		nresp, err = sess.Delete(u, nil, &res, &e)
	}

	var resp = Response{Status: nresp.Status(), Message: e.Message}
	if err != nil {
		return err, &resp
	}
	if nresp.Status() == 401 {
		return errors.New("unauthorised - check your username and passwd"), &resp
	}
	if nresp.Status() >= 300 {
		return errors.New(e.Message), &resp
	} else {
		// all is good in the world
		return nil, &resp
	}
}

func (f *Device) PrintResponse(input interface{}) {

	jsonresp, err := json.MarshalIndent(&input, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(jsonresp))

}
