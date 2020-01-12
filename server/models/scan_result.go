package models

type ScanResult struct {
	TrackerID int    `codec:"-" json:"tracker_id"`
	SSID      string `json:"ssid"`
	BSSID     string `json:"bssid"`
	RSSID     int    `json:"rssid"`
}
