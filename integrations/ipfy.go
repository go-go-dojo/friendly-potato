package integrations

import (
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net"
	"net/http"
)

const IpFyExtIPEndpoint = "https://api.ipify.org?format=json"


type IPFy struct {
	Endpoint string
	Resolved net.IP
}

var ipFyApi *IPFy

//InitIPFy
//function to initialize IPFyApi
func InitIPFy() (err error) {
	ipFyApi,err = NewIPFy()
	return
}

func NewIPFy()(*IPFy,error){
	api:=&IPFy{
		Endpoint: IpFyExtIPEndpoint,
		Resolved: net.IP{},
	}
	return api,api.healthCheck()
}

func (h *IPFy) healthCheck() (err error) {
	ip, err := h.GetExtIP()
	if err != nil {
		return
	}
	h.Resolved = ip
	return
}

func (h IPFy) treatResponse(resp []byte) (ip net.IP, err error) {
	val, _, _, err := jsonparser.Get(resp, "ip")
	if err != nil {
		return
	}
	ip = net.ParseIP(string(val))
	return
}

//IPFy
//function to get the external ip address from IPFy
func (h IPFy) GetExtIP() (ip net.IP, err error) {
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
