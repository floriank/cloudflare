package main

import (
	"flag"
	"fmt"
	"github.com/consulted/cloudflare"
	"os"
)

// Email, Token, Name, IP

var email, token, name, content, zoneName, rType string
var delRecord bool
var client cloudflare.Client
var record cloudflare.Record

const (
	CFToken = "CF_TOKEN"
	CFEmail = "CF_EMAIL"
	CFZone  = "CF_ZONE"
)

func init() {
	flag.StringVar(&name, "n", "", "the Name for the new record (Type A)")
	flag.StringVar(&content, "c", "", "the content value for the new record")
	flag.StringVar(&rType, "t", "A", "the type for the record (A/CNAME/MX/TXT/SPF/AAAA/NS/SRV/LOC)")
	flag.StringVar(&zoneName, "z", os.Getenv(CFZone), "the name of the zone to use (settable via CF_ZONE)")
	flag.BoolVar(&delRecord, "d", false, "delete the given record")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	email = os.Getenv(CFEmail)
	token = os.Getenv(CFToken)
	checkParams()
}

func main() {

	client.Email = email
	client.Token = token

	record.Name = name
	record.Type = "A"
	record.Content = content
	record.Ttl = "1" // Automatic

	zones, err := client.GetZoneList()

	if err != nil {
		fmt.Println("could not retrieve the zone list")
		os.Exit(1)
	}

	zone, err := zones.Find(zoneName)

	if err != nil {
		fmt.Println("could not find zone specified: " + zoneName)
	}

	if delRecord == true {
		deleteRecord(zone)
	} else {
		addRecord(zone)
	}
}

func deleteRecord(zone cloudflare.Zone) {
	recordList, err := client.GetRecordList(zone, 0)
	if err != nil {
		fmt.Println("Could not retrieve record list: " + err.Error())
		os.Exit(1)
	}

	record, err = recordList.Find(record.Name)
	if err != nil {
		fmt.Println("Could not find record: " + err.Error())
		os.Exit(1)
	}

	_, err = client.RemoveRecord(zone, record)
	if err != nil {
		fmt.Println("Could not remove record: " + err.Error())
		os.Exit(1)
	}
}

func addRecord(zone cloudflare.Zone) {
	record, err := client.AddRecord(zone, record)

	if err != nil {
		fmt.Println("Could not add record: " + err.Error())
		os.Exit(1)
	}

	// activate cloudflare proxy
	record.ServiceMode = "1"
	record, err = client.UpdateRecord(zone, record)

	if err != nil {
		fmt.Println("Could not update record: " + err.Error())
		os.Exit(1)
	}
}

func checkParams() {
	message := " is not set!"
	err := false
	if email == "" {
		fmt.Println(CFEmail + message)
		err = true
	}

	if token == "" {
		fmt.Println(CFToken + message)
		err = true
	}

	if name == "" {
		fmt.Println("Record name" + message)
		err = true
	}

	if content == "" && delRecord == false {
		fmt.Println("Content" + message)
		err = true
	}

	if err == true {
		os.Exit(1)
	}

}
