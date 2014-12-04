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
		opts.Content = options.Delete.Content
		opts.SkipConfirm = options.Delete.SkipConfirmation

		deleteRecords(&opts)
	}
}

func deleteRecords(o *Options) {
	client := makeClient(o)
	zone, err := findZone(o.Zone, client)

	if err != nil {
		fmt.Println("Zone not found: " + err.Error())
		exitWithError()
	}

	list, err := client.GetRecordList(zone, 0)

	if err != nil {
		fmt.Println("Records could not be retrieved: " + err.Error())
	}
	records := list.FindAll(o.Content)

	for _, record := range records.Records {
		if o.SkipConfirm || askForConfirmation("Do you really want to remove \""+record.Name+"\"?") {
			_, err := client.RemoveRecord(zone, record)
			if err != nil {
				fmt.Println("Could not remove \"" + record.Name + "\"")
			}
		}
	}
}

func addRecord(o *Options) {
	client := makeClient(o)

	zone, err := findZone(o.Zone, client)

	if err != nil {
		fmt.Println("Zone not found: " + err.Error())
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

func findZone(name string, c *cloudflare.Client) (z cloudflare.Zone, err error) {
	zones, err := c.GetZoneList()

	if err != nil {
		return cloudflare.Zone{}, err
	}

	zone, err := zones.Find(name)

	if err != nil {
		fmt.Println("Could not fetch zone: " + err.Error())
		return cloudflare.Zone{}, err
	}

	return zone, nil
}

func exitWithError() {
	os.Exit(1)
}

func askForConfirmation(message string) bool {
	fmt.Println(message)
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		panic(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		return askForConfirmation(message)
	}
}

func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}
