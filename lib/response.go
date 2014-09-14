package lib

type CfResponse struct {
	Request struct {
		Act string `json:"act"`
	} `json:"request"`
	Response struct {
		Zones ZoneList `json:"zones"`
	} `json:"response"`
	Result string      `json:"result"`
	Msg    interface{} `json:"msg"`
}
