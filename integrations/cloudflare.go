package integrations

import (
	"github.com/cloudflare/cloudflare-go"
	cf "github.com/cloudflare/cloudflare-go"
)

var api *cloudflare.API

// InitAPI -- Configure token to cloudflare
func InitAPI(apiToken string) (err error) {
	api, err = cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		return err
	}
	return nil
}

func HealthCheck() {
	//api.Raw("", endpoint string, data interface{}) (json.RawMessage, error));
}

// ListZones -- get all zones
func ListZones() ([]cf.Zone, error) {
	return api.ListZones()
}
