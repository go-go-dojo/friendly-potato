package integrations

import (
	"time"

	cf "github.com/cloudflare/cloudflare-go"
)

var api *cf.API

type Record struct {
	Id        string
	Zone      string
	Ttl       uint
	DnsType   string
	DnsData   string
	Timestamp time.Time
}

type Zone struct {
	Id      string
	Name    string
	Records []Record
}

type Zones []Zone

// InitAPI -- Configure token to cloudflare
func InitAPI(apiToken string) (err error) {
	api, err = cf.NewWithAPIToken(apiToken)
	if err != nil {
		return err
	}
	return nil
}

func HealthCheck() bool {
	_, err := api.ListZones()
	if err != nil {
		return false
	}
	return true
}

// ListZones -- get all zones
func ListZones() ([]cf.Zone, error) {
	return api.ListZones()
}
