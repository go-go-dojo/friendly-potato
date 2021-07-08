package integrations

import (
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net"
	"net/http"
)

const HttpBinExtIPEndpoint = "http://httpbin.org/ip"

type HttpBin struct {
	Endpoint string
	Resolved net.IP
}

var httpBinApi *HttpBin

//InitHttpBin
//function to initialize httpBinApi
func InitHttpBin() (err error) {
	httpBinApi,err = NewHttpBin()
	return
}

func NewHttpBin()(*HttpBin,error){
	api:=&HttpBin{
		Endpoint: HttpBinExtIPEndpoint,
		Resolved: net.IP{},
	}
	return api,api.healthCheck()
}

func (h *HttpBin) healthCheck() (err error) {
	ip, err := h.GetExtIP()
	if err != nil {
		return
	}
	h.Resolved = ip
	return
}

func (h HttpBin) treatResponse(resp []byte) (ip net.IP, err error) {
	val, _, _, err := jsonparser.Get(resp, "origin")
	if err != nil {
		return
	}
	ip = net.ParseIP(string(val))
	return
}

//HttpBin
//function to get the external ip address from HttpBin
func (h HttpBin) GetExtIP() (ip net.IP, err error) {
	resp, err := http.Get(h.Endpoint)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return h.treatResponse(body)
}
