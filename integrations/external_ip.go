package integrations

import (
	"fmt"
	"net"
)

const NoExtIPApi = "no external ip address API available"

// GetExternalIP
// function to get the external ip from any avaliable external ip integration
func GetExternalIP() (err error, IP net.IP) {
	var eh,ei error
	if IP, eh = httpBinApi.GetExtIP(); eh == nil {
		return
	}
	if IP, ei = ipFyApi.GetExtIP(); ei == nil {
		return
	}
	err=fmt.Errorf("%v httpBinErr: %v IpFyErr: %v",NoExtIPApi,eh,ei)
	return
}
