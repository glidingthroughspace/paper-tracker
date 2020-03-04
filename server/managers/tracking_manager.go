package managers

import (
	"math"
	"paper-tracker/models"
	"paper-tracker/utils/collections"
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

func (*TrackingManager) ConsolidateScanResults(scanResults []*models.ScanResult) []models.BSSIDTrackingData {
	scanResultsPerBSSID := make(map[string][]*models.ScanResult)
	for _, v := range scanResults {
		scanResultsPerBSSID[v.BSSID] = append(scanResultsPerBSSID[v.BSSID], v)
	}
	var trackingData []models.BSSIDTrackingData
	for bssid, scanResults := range scanResultsPerBSSID {
		rssis := getRSSIs((scanResults))
		trackingData = append(trackingData, models.BSSIDTrackingData{
			BSSID:   bssid,
			Minimum: collections.MinOf(rssis...),
			Maximum: collections.MaxOf(rssis...),
			Median:  collections.MedianOf(rssis...),
			Mean:    collections.MeanOf(rssis...),
			Quantiles: models.Quantiles{
				FirstQuartile: collections.FirstQuartileOf(rssis...),
				ThirdQuartile: collections.ThirdQuartileOf(rssis...),
			},
		})
	}
	return trackingData
}

func (tm *TrackingManager) GetRoomMatchingBest(rooms []*models.Room, scanResults []*models.ScanResult) *models.Room {
	var bestMatch *models.ScoredRoom = nil
	scoredRooms := tm.ScoreRoomsForScanResults(rooms, scanResults)
	for _, room := range scoredRooms {
		if bestMatch == nil && room.Score > 0.1e-7 {
			bestMatch = room
		} else if bestMatch != nil && room.Score > bestMatch.Score {
			bestMatch = room
		}
	}
	// no room had a score > 0
	if bestMatch == nil {
		return nil
	}
	return &bestMatch.Room
}

func (tm *TrackingManager) ScoreRoomsForScanResults(rooms []*models.Room, scanResults []*models.ScanResult) []*models.ScoredRoom {
	scoredRooms := []*models.ScoredRoom{}
	for _, room := range rooms {
		scoredRooms = append(scoredRooms, models.ScoredRoomFromRoom(room, tm.ScoreRoomForScanResults(room, scanResults)))
	}
	return scoredRooms
}

// ScoreRoomForScanResults calculates a score of how likely the scan results are from the given
// room. The larger the returned score, the greater the likelyness of match.
func (*TrackingManager) ScoreRoomForScanResults(room *models.Room, scanResults []*models.ScanResult) float64 {
	score := 0.0
	for _, trackingData := range room.TrackingData {
		srs := getScanResultsForBSSID(trackingData.BSSID, scanResults)
		for _, sr := range srs {
			score += getScoreForScanResultAndTrackingData(trackingData, sr)
		}
	}
	return score
}

// getScoreForScanResultAndTrackingData calculates the score for a single scan result and the given
// tracking data.
func getScoreForScanResultAndTrackingData(td models.BSSIDTrackingData, sr *models.ScanResult) float64 {
	// TODO: Evaluate how good this scoring works
	//       We might also use distances (e.g. d(Mean, RSSI)) to get more fine-grained scores
	score := 0.0
	if sr.RSSI < td.Maximum && sr.RSSI > td.Minimum {
		score += 1
	}
	if float64(sr.RSSI) < td.Quantiles.ThirdQuartile && float64(sr.RSSI) > td.Quantiles.FirstQuartile {
		score += 5
	}
	if isInRange(float64(sr.RSSI), td.Mean, 10) {
		score += math.Abs(td.Mean - float64(sr.RSSI))
	}
	if isInRange(float64(sr.RSSI), td.Median, 10) {
		score += math.Abs(td.Median - float64(sr.RSSI))
	}
	return score
}

// isInRange returns whether the given actual number is in the range of the wanted number +/- a
// delta.
func isInRange(actual, wanted, delta float64) bool {
	return actual > wanted-delta && actual < wanted+delta
}

func getScanResultsForBSSID(bssid string, scanResults []*models.ScanResult) []*models.ScanResult {
	matches := []*models.ScanResult{}
	for _, v := range scanResults {
		if v.BSSID == bssid {
			matches = append(matches, v)
		}
	}
	return matches
}

func getRSSIs(scanResults []*models.ScanResult) []int {
	var rssis []int
	for _, v := range scanResults {
		rssis = append(rssis, v.RSSI)
	}
	return rssis
}
