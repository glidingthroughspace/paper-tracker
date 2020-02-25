package models

type ScanResult struct {
	TrackerID TrackerID `codec:"-" json:"tracker_id"`
	SSID      string    `json:"ssid"`
	BSSID     string    `json:"bssid"`
	RSSI      int       `json:"rssi"`
}
