package cloudflare

import (
	"errors"
)

const (
	AutomaticTtl = "1"
)

type RecordList struct {
	HasMore bool     `json:"has_more"`
	Count   int      `json:"count"`
	Records []Record `json:"objs"`
}

type Record struct {
	RecId          string      `json:"rec_id"`
	RecTag         string      `json:"rec_tag"`
	ZoneName       string      `json:"zone_name"`
	Name           string      `json:"name"`
	DisplayName    string      `json:"display_name"`
	Type           string      `json:"type"`
	Prio           interface{} `json:"prio"`
	Content        string      `json:"content"`
	DisplayContent string      `json:"display_content"`
	Ttl            string      `json:"ttl"`
	TtlCeil        int         `json:"ttl_ceil"`
	SslId          interface{} `json:"ssl_id"`
	SslStatus      interface{} `json:"ssl_status"`
	SslExpiresOn   interface{} `json:"ssl_expires_on"`
	AutoTtl        int         `json:"auto_ttl"`
	ServiceMode    string      `json:"service_mode"`
	Props          struct {
		Proxiable   int `json:"proxiable"`
		CloudOn     int `json:"cloud_on"`
		CfOpen      int `json:"cf_open"`
		Ssl         int `json:"ssl"`
		ExpiredSsl  int `json:"expired_ssl"`
		ExpiringSsl int `json:"expiring_ssl"`
		PendingSsl  int `json:"pending_ssl"`
	} `json:"props"`
}

func (list *RecordList) Find(content string) (record Record, err error) {
	for _, record := range list.Records {
		if record.Content == content {
			return record, nil
		}
	}
	return Record{}, errors.New("record not found!")
}

func (list *RecordList) FindAll(content string) (records *RecordList) {
	r := RecordList{}
	for _, record := range list.Records {
		if record.Content == content {
			r.Records = append(r.Records, record)
		}
	}
	return &r
}
