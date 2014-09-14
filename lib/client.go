package cloudflare

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	endPoint = "https://www.cloudflare.com/api_json.html"
)

type Client struct {
	Email, Token string
}

func (client *Client) GetZoneList() (zones ZoneList, err error) {
	var params map[string]string
	resp, err := client.post("zone_load_multi", params)
	if err == nil {
		return makeZoneList(resp)
	} else {
		return ZoneList{}, err
	}
}

func (client *Client) GetRecordList(zone Zone, offset int) (records RecordList, err error) {
	params := make(map[string]string, 2)

	params["z"] = zone.ZoneName

	if offset > 0 {
		params["o"] = string(offset)
	}

	resp, err := client.post("rec_load_all", params)

	if err != nil {
		return RecordList{}, err
	} else {
		return makeRecordList(resp)
	}
}

func (client *Client) AddRecord(zone Zone, record Record) (r Record, err error) {
	params := make(map[string]string, 5)

	params["z"] = zone.ZoneName
	params["type"] = record.Type
	params["name"] = record.Name
	params["content"] = record.Content
	params["ttl"] = record.Ttl

	_, err = client.post("rec_new", params)

	return record, err
}

func (client *Client) UpdateRecord(zone Zone, record Record) (r Record, err error) {
	params := make(map[string]string, 6)

	params["z"] = zone.ZoneName
	params["type"] = record.Type
	params["id"] = record.RecId
	params["name"] = record.Name
	params["content"] = record.Content
	params["ttl"] = record.Ttl

	_, err = client.post("rec_edit", params)
	return record, err
}

func (client *Client) RemoveRecord(zone Zone, record Record) (r Record, err error) {
	params := make(map[string]string, 2)

	params["z"] = zone.ZoneName
	params["id"] = record.RecId

	_, err = client.post("rec_delete", params)

	return record, err
}

func makeZoneList(resp *http.Response) (zones ZoneList, err error) {
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ZoneList{}, err
	}
	var response CfResponse
	err = json.Unmarshal(contents, &response)
	return response.Response.Zones, nil
}

func makeRecordList(resp *http.Response) (records RecordList, err error) {
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RecordList{}, err
	}
	var response CfResponse
	err = json.Unmarshal(contents, &response)
	return response.Response.Records, nil
}

func (client *Client) post(act string, params map[string]string) (resp *http.Response, err error) {
	clientParams := url.Values{}
	clientParams.Set("a", act)
	clientParams.Set("tkn", client.Token)
	clientParams.Set("email", client.Email)

	if len(params) > 0 {
		for k, v := range params {
			clientParams.Set(k, v)
		}
	}

	return http.PostForm(endPoint, clientParams)
}
