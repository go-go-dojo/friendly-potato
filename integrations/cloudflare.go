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

func (z *Zone) AppendRecords(r ...Record) {
	if z.Records == nil {
		z.Records = []Record{}
	}
	for _, rr := range r {
		z.Records = append(z.Records, rr)
	}
	return
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
func ListZones() (zones Zones, err error) {
	cfz, err := api.ListZones()
	if err != nil {
		return
	}

	for _, z := range cfz {
		zz := Zone{
			Id:      z.ID,
			Name:    z.Name,
			Records: []Record{},
		}
		zones = append(zones, zz)
	}
	return
}

func CreateZone(zone Zone) (createdZone Zone, err error) {
	z, err := api.CreateZone(zone.Name, true, cf.Account{ID: api.AccountID}, "full")
	if err != nil {
		return
	}
	createdZone = Zone{
		Id:      z.ID,
		Name:    z.Name,
		Records: []Record{},
	}
	return
}
