package models

// BSSIDTrackingData is consolidated tracking data for a single BSSID
// Note that all sizes (min, max, ...) are in their mathematical form, not the logical form:
// a RSSI of -90dBm is worse than -10dBm. The Minimum in this case is -90dBm, as one might expect.
type BSSIDTrackingData struct {
	ID        uint      `json:"id"`
	RoomID    RoomID    `json:"roomID"`
	BSSID     string    `json:"bssid" gorm:"column:bssid"`
	Minimum   int       `json:"minimum"`
	Maximum   int       `json:"maximum"`
	Median    float64   `json:"median"`
	Mean      float64   `json:"mean"`
	Quantiles Quantiles `json:"quantiles" gorm:"embedded"`
}

type Quantiles struct {
	FirstQuartile float64
	ThirdQuartile float64
}
