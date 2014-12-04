package main

import (
	"fmt"
	"github.com/consulted/cloudflare"
	"github.com/voxelbrain/goptions"
	"os"
)

var email, token, name, content, zoneName, rType string

const (
	CFToken = "CF_TOKEN"
	CFEmail = "CF_EMAIL"
	CFZone  = "CF_ZONE"
)

func main() {

	options := struct {
		Zone  string        `goptions:"-z, --zone, description='The zone to use'"`
		Token string        `goptions:"--token, description='The token for the CF API'"`
		Email string        `goptions:"--email, description='The email for the CF API'"`
		Help  goptions.Help `goptions:"-h, --help, description='Show this help'"`

		goptions.Verbs
		Add struct {
			Content string `goptions:"-c, --content, obligatory, description='The content for the record'"`
			Name    string `goptions:"-n, --name, obligatory, description='The name of the record'"`
			Type    string `goptions:"-t, --type, description='The type of the record'"`
		} `goptions:"add"`

		Delete struct {
			Content          string `goptions:"-c, --content, obligatory, description='The content for the record'"`
			SkipConfirmation bool   `goptions:"-y, --yes, description='Skip confirmation'"`
		} `goptions:"delete"`
	}{
		Zone:  os.Getenv(CFZone),
		Email: os.Getenv(CFEmail),
		Token: os.Getenv(CFToken),
	}

	goptions.ParseAndFail(&options)

	opts := Options{
		Email: options.Email,
		Token: options.Token,
		Zone:  options.Zone,
	}

	switch {
	case options.Verbs == "add":
		opts.Content = options.Add.Content
		opts.Name = options.Add.Name
		opts.Type = options.Add.Type

		addRecord(&opts)
	case options.Verbs == "delete":
		fmt.Println("delete all records")
	}
}

// func deleteRecord(zone cloudflare.Zone) {
// 	recordList, err := client.GetRecordList(zone, 0)
// 	if err != nil {
// 		fmt.Println("Could not retrieve record list: " + err.Error())
// 		os.Exit(1)
// 	}

// 	record, err = recordList.Find(record.Name)
// 	if err != nil {
// 		fmt.Println("Could not find record: " + err.Error())
// 		os.Exit(1)
// 	}

// 	_, err = client.RemoveRecord(zone, record)
// 	if err != nil {
// 		fmt.Println("Could not remove record: " + err.Error())
// 		os.Exit(1)
// 	}
// }

func addRecord(o *Options) {
	client := makeClient(o)
	zones, err := client.GetZoneList()

	if err != nil {
		fmt.Println("Could not fetch zones: " + err.Error())
		exitWithError()
	}

	zone, err := zones.Find(o.Zone)

	if err != nil {
		fmt.Println("Could not fetch zone: " + err.Error())
		exitWithError()
	}

	record := cloudflare.Record{
		Name:    o.Name,
		Type:    o.Type,
		Content: o.Content,
		Ttl:     "1", // automatic
	}

	record, err = client.AddRecord(zone, record)

	if err != nil {
		fmt.Println("Could not add record: " + err.Error())
		exitWithError()
	}

	// activate cloudflare proxy
	record.ServiceMode = "1"
	record, err = client.UpdateRecord(zone, record)

	if err != nil {
		fmt.Println("Could not update record: " + err.Error())
		exitWithError()
	}
}

func makeClient(o *Options) *cloudflare.Client {
	return &cloudflare.Client{
		Email: o.Email,
		Token: o.Token,
	}
}

func exitWithError() {
	os.Exit(1)
}
