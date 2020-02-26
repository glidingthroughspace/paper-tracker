package managers

import (
	"paper-tracker/models"
	"sort"
)

var trackingManager *TrackingManager

type TrackingManager struct {
}

func CreateTrackingManager() *TrackingManager {
	if trackingManager != nil {
		return trackingManager
	}

	return trackingManager
}

func GetTrackingManager() *TrackingManager {
	return trackingManager
}

func (*TrackingManager) ConsolidateScanResults(scanResults []models.ScanResult) []models.BSSIDTrackingData {
	scanResultsPerBSSID := make(map[string][]models.ScanResult)
	for _, v := range scanResults {
		scanResultsPerBSSID[v.BSSID] = append(scanResultsPerBSSID[v.BSSID], v)
	}
	var trackingData []models.BSSIDTrackingData
	for bssid, scanResults := range scanResultsPerBSSID {
		trackingData = append(trackingData, models.BSSIDTrackingData{
			BSSID:   bssid,
			Minimum: getMin(getRSSIs(scanResults)...),
			Maximum: getMax(getRSSIs(scanResults)...),
			Median:  getMedian(getRSSIs(scanResults)...),
			Mean:    getMean(getRSSIs(scanResults)...),
		})
	}
	return trackingData
}

func getRSSIs(scanResults []models.ScanResult) []int {
	var rssis []int
	for _, v := range scanResults {
		rssis = append(rssis, v.RSSI)
	}
	return rssis
}

func getMean(values ...int) float64 {
	var sum float64 = 0.0
	for _, v := range values {
		sum += float64(v)
	}
	return sum / float64(len(values))
}

func getMedian(values ...int) float64 {
	isOddAmountOfValues := (len(values)%2 == 1)
	middleIndex := len(values) / 2

	sort.Ints(values)

	if isOddAmountOfValues {
		return float64(values[middleIndex])
	}

	return (float64(values[middleIndex-1]) + float64(values[middleIndex])) / 2.0
}

func getMin(values ...int) int {
	sort.Ints(values)
	return values[0]
}

func getMax(values ...int) int {
	sort.Ints(values)
	return values[len(values)-1]
}
