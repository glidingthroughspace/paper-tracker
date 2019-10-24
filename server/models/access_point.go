package models

type AccessPoint struct {
	ID        int    `json:"id,omitempty"`
	NetworkID int    `json:"network_id,omitempty"`
	BSSID     string `json:"bssid,omitempty"`
}

type AccessPointReadout struct {
	AccessPoint
	SSID string `json:"ssid,omitempty"`
	RSSI int    `json:"rssi,omitempty"`
}
