package models

type ScanResult struct {
	TrackerID int `codec:"-"`
	SSID      string
	BSSID     int
	RSSID     int
}
