package integrations

import "github.com/cloudflare/cloudflare-go"

var api *cloudflare.API

func InitAPI(apiKey, apiEmail string)(err error){
	api,err=cloudflare.New(apiKey,apiEmail)
	if err!=nil{
		return err
	}
	return nil
}
