package main

import (
	"flag"
	"fmt"
	cf "github.com/consulted/cloudflare/lib"
)

var token, email string

func init() {
	flag.StringVar(&token, "t", "", "the CF Token to use for the API")
	flag.StringVar(&email, "e", "", "the Email to use")
}

func main() {
	flag.Parse()
	client := cf.Client{
		Email: email,
		Token: token,
	}
	domains, err := client.GetZoneList()
	if err != nil {
		panic(err)
	}
	zone := domains.Zones[0]
	records, err := client.GetRecordList(zone, 0)

	record := records.Find("florian-test")
	result := client.RemoveRecord(zone, record)
	fmt.Println(result)
}
