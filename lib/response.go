package cloudflare

type CfResponse struct {
	Request struct {
		Act string `json:"act"`
	} `json:"request"`
	Response struct {
		Zones   ZoneList   `json:"zones"`
		Records RecordList `json:"recs"`
	} `json:"response"`
	Result string      `json:"result"`
	Msg    interface{} `json:"msg"`
}
