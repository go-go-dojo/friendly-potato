package integrations

import (
	"fmt"
	cf "github.com/cloudflare/cloudflare-go"
)

const healthCheckFailed = "health check for cloudflare failed"

var api *cf.API

//InitApi
//initialize api token
func InitCloudFlareAPI(apiToken string) (err error) {
	api, err = cf.NewWithAPIToken(apiToken)
	if err != nil {
		return
	}
	if !healthCheck(){
		err = fmt.Errorf(healthCheckFailed)
		return
	}
	return
}

func healthCheck() bool {
	_, err := api.UserDetails()
	if err != nil {
		return false
	}
	return true
}

// TODO: later expand health check to:
// - check connection
// - see if user has permissions.

// ListZones -- get all zones
func ListZones() (zones Zones, err error) {
	cfz, err := api.ListZones()
	if err != nil {
		return
	}
	for _, z := range cfz {
		zz := Zone{}
		zz.translateFromCloudflare(z)
		zones.appendZone(zz)
	}
	return
}

func CreateZone(zone Zone) (createdZone Zone, err error) {
	z, err := api.CreateZone(zone.Name, true, cf.Account{ID: api.AccountID}, "full")
	if err != nil {
		return
	}
	createdZone = Zone{}
	createdZone.translateFromCloudflare(z)
	return
}
