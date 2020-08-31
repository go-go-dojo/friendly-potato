package main

import (
	"fmt"
	"friendly-potato/integrations"
	"log"
)

func main() {
	fmt.Println("teste")
	err := integrations.InitAPI("OpmWl7p_ECwb1U2YMVlSXhqFW2017_we9lMCQ_4V")

	if err != nil {
		log.Fatalf("Fatal on auth %v", err)
	}

	zones, err := integrations.ListZones()
	fmt.Printf("%v", zones)
	if err != nil {
		log.Fatal(err)
	}

}
