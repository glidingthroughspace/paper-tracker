package models

type BSSIDTrackingData struct {
	BSSID   int     `json:"bssid"`
	Minimum int     `json:"minimum"`
	Maximum int     `json:"maximum"`
	Median  int     `json:"median"`
	Average float32 `json:"average"`
}
