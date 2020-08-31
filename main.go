package main

import (
	"fmt"
	"friendly-potato/integrations"
	"log"
)

func main() {

	err := integrations.InitCloudFlareAPI("WbEIMNSSDoTAxANuuN9OgTphO7Fq7h7dc6YUzw8g")

	if err != nil {
		log.Fatalf("Fatal on auth %v", err)
	}

	zones, err := integrations.ListZones()
	if err != nil {
		log.Fatal(err)
	}

	for _, z := range zones {
		fmt.Printf("zone: %v", z.Resource.Name)
	}

	testZone:=integrations.Zone{Resource: integrations.DomainResource{Name: "1-qr.me",},}

	zone, err := integrations.CreateZone(testZone)
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Printf("%v\n",zone.Resource.Name)
	}

	zone, err = integrations.DeleteZone(testZone)
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Printf("%v\n",zone.Resource.Name)
	}

}
