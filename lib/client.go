package lib

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
	resp, err := client.post("zone_load_multi")
	if err == nil {
		return makeZoneList(resp)
	} else {
		return ZoneList{}, err
	}
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

func (client *Client) post(act string) (resp *http.Response, err error) {
	return http.PostForm(endPoint, url.Values{"a": {act}, "tkn": {client.Token}, "email": {client.Email}})
}
