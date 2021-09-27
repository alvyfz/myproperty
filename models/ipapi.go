package model

type IpApi struct {
	IP     string `json:"ip"`
	City   string `json:"city"`
	Region string `json:"region"`
	Postal string `json:"postal"`
}
