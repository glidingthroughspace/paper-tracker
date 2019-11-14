package models

type BSSIDTrackingData struct {
	BSSID   int
	Minimum int
	Maximum int
	Median  int
	Average float32
}
