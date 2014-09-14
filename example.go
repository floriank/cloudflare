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
	records, err := client.GetRecordList(domains.Zones[0], 0)
	fmt.Println(records)
}
