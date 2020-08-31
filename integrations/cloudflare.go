package integrations

import (
	"fmt"
	cf "github.com/cloudflare/cloudflare-go"
	"time"
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
	if b,err:=healthCheck();!b{
		return fmt.Errorf("gotErr: %v", err)
	}
	return
}

func healthCheck() (bool,error) {
	_, err := api.ListZones()
	if err != nil {
		return false,err
	}
	return true,nil
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

// CreateZone -- creates a zone
func CreateZone(zone Zone) (createdZone Zone, err error) {
	z, err := api.CreateZone(zone.Resource.Name, true, cf.Account{ID: api.AccountID}, "full")
	if err != nil {
		return
	}
	createdZone = Zone{}
	createdZone.translateFromCloudflare(z)
	return
}

// DeleteZone -- deletes a zone
func DeleteZone(zone Zone)(deletedZone Zone,err error){
	cfz:=zone.Resource.translateToCloudflare()
	if cfz.ID == ""{
		cfz.ID,err=api.ZoneIDByName(zone.Resource.Name)
		return
	}
	dz, err := api.DeleteZone(cfz.ID)
	if err!=nil{
		return
	}
	cfz.ID=dz.ID
	cfz.Status="deleted"
	cfz.ModifiedOn=time.Now()
	deletedZone.translateFromCloudflare(cfz)
	return
}