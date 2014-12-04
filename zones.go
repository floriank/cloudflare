package cloudflare

import (
	"errors"
)

type Zone struct {
	ZoneId          string        `json:"zone_id"`
	UserId          string        `json:"user_id"`
	ZoneName        string        `json:"zone_name"`
	DisplayName     string        `json:"display_name"`
	ZoneStatus      string        `json:"zone_status"`
	ZoneMode        string        `json:"zone_mode"`
	HostId          interface{}   `json:"host_id"`
	ZoneType        string        `json:"zone_type"`
	HostPubname     interface{}   `json:"host_pubname"`
	HostWebsite     interface{}   `json:"host_website"`
	Vtxt            interface{}   `json:"vtxt"`
	Fqdns           []string      `json:"fqdns"`
	Step            string        `json:"step"`
	ZoneStatusClass string        `json:"zone_status_class"`
	ZoneStatusDesc  string        `json:"zone_status_desc"`
	NsVanityMap     []interface{} `json:"ns_vanity_map"`
	OrigRegistrar   string        `json:"orig_registrar"`
	OrigDnshost     interface{}   `json:"orig_dnshost"`
	OrigNsNames     string        `json:"orig_ns_names"`
	Props           struct {
		DnsCname       int           `json:"dns_cname"`
		DnsPartner     int           `json:"dns_partner"`
		DnsAnonPartner int           `json:"dns_anon_partner"`
		Plan           string        `json:"plan"`
		Pro            int           `json:"pro"`
		ExpiredPro     int           `json:"expired_pro"`
		ProSub         int           `json:"pro_sub"`
		PlanSub        int           `json:"plan_sub"`
		Ssl            int           `json:"ssl"`
		ExpiredSsl     int           `json:"expired_ssl"`
		ExpiredRsPro   int           `json:"expired_rs_pro"`
		ResellerPro    int           `json:"reseller_pro"`
		ResellerPlans  []interface{} `json:"reseller_plans"`
		ForceInteral   int           `json:"force_interal"`
		SslNeeded      int           `json:"ssl_needed"`
		AlexaRank      int           `json:"alexa_rank"`
		HasVanity      int           `json:"has_vanity"`
	} `json:"props"`
	ConfirmCode struct {
		ZoneDelete     string `json:"zone_delete"`
		ZoneDeactivate string `json:"zone_deactivate"`
		ZoneDevMode1   string `json:"zone_dev_mode1"`
	} `json:"confirm_code"`
	Allow []string `json:"allow"`
}

type ZoneList struct {
	HasMore bool   `json:"has_more"`
	Count   int    `json:"count"`
	Zones   []Zone `json:"objs"`
}

func (z *ZoneList) Find(name string) (zone Zone, err error) {
	for _, zone := range z.Zones {
		if zone.ZoneName == name {
			return zone, nil
		}
	}
	return Zone{}, errors.New("zone not found!")
}
