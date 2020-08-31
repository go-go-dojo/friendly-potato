package main

import (
	"fmt"
	"friendly-potato/integrations"
	"log"
)

func main() {
	fmt.Println("teste")
	err := integrations.InitCloudFlareAPI("WbEIMNSSDoTAxANuuN9OgTphO7Fq7h7dc6YUzw8g")

	if err != nil {
		log.Fatalf("Fatal on auth %v", err)
	}

	zones, err := integrations.ListZones()
	fmt.Printf("%v", zones)
	if err != nil {
		log.Fatal(err)
	}

	zones, err = integrations.ListZones()
	if err != nil {
		log.Fatal(err)
	}

	for _, z := range zones {
		fmt.Printf("zone: %v id: %v\n", z.Resource, z.Resource.Name)
	}

	testZone:=integrations.Zone{Resource: integrations.DomainResource{Name: "xptoteste9000.com",},}

	zone, err := integrations.CreateZone(testZone)
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Printf("%v",zone)
	}

	zone, err = integrations.DeleteZone(testZone)
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Printf("%v",zone)
	}

}
