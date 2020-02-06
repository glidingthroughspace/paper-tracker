package models

type ScanResult struct {
	TrackerID int `codec:"-"`
	SSID      string
	BSSID     string
	RSSID     int
}
