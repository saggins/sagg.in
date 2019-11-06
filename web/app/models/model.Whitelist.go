package models

type Whitelist struct {
	Mcuuid string `json:"uuid"`
	Mcuser string `json:"name"`
	Name string `json:"rname"`
	Ip string `json:"ip"`
}