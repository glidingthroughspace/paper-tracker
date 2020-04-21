package managers

import (
	"math"
	"paper-tracker/models"
	"paper-tracker/utils/collections"

	log "github.com/sirupsen/logrus"
)

var trackingManager TrackingManager

type TrackingManager interface {
	GetRoomMatchingBest(rooms []*models.Room, scanResults []*models.ScanResult) *models.Room
	ConsolidateScanResults(scanResults []*models.ScanResult) []models.BSSIDTrackingData
	ScoreRoomsForScanResults(rooms []*models.Room, scanResults []*models.ScanResult) map[*models.Room]float64
}

type TrackingManagerImpl struct {
}

func CreateTrackingManager() TrackingManager {
	if trackingManager != nil {
		return trackingManager
	}

	trackingManager = &TrackingManagerImpl{}

	return trackingManager
}

func GetTrackingManager() TrackingManager {
	return trackingManager
}

func (*TrackingManagerImpl) ConsolidateScanResults(scanResults []*models.ScanResult) []models.BSSIDTrackingData {
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

func (tm *TrackingManagerImpl) GetRoomMatchingBest(rooms []*models.Room, scanResults []*models.ScanResult) *models.Room {
	var bestMatch *models.Room = nil
	var bestScore float64 = -1.0
	scoredRooms := tm.ScoreRoomsForScanResults(rooms, scanResults)
	for room, score := range scoredRooms {
		log.Debugf("Scored room for matching: %s (%d): %f", room.Label, room.ID, score)
		if (bestMatch == nil && score > 0.1e-7) || (bestMatch != nil && score > bestScore) {
			bestScore = score
			bestMatch = room
		}
	}
	if bestMatch != nil {
		log.Debugf("Selected the best match to be: %s", bestMatch.Label)
	} else {
		log.Debug("Did not find a matching room")
	}
	return bestMatch
}

func (tm *TrackingManagerImpl) ScoreRoomsForScanResults(rooms []*models.Room, scanResults []*models.ScanResult) map[*models.Room]float64 {
	scoredRooms := make(map[*models.Room]float64)
	for _, room := range rooms {
		scoredRooms[room] = tm.ScoreRoomForScanResults(room, scanResults)
	}
	return scoredRooms
}

// ScoreRoomForScanResults calculates a score of how likely the scan results are from the given
// room. The larger the returned score, the greater the likelyness of match.
func (*TrackingManagerImpl) ScoreRoomForScanResults(room *models.Room, scanResults []*models.ScanResult) float64 {
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
	if sr.RSSI <= td.Maximum && sr.RSSI >= td.Minimum {
		score += 1
	}
	if float64(sr.RSSI) < td.Quantiles.ThirdQuartile && float64(sr.RSSI) > td.Quantiles.FirstQuartile {
		score += 5
	}
	// The +0.5e-7 additions below are to prevent division by zero
	if isInRange(float64(sr.RSSI), td.Mean, 10) {
		score += (1.0 / (math.Abs(td.Mean-float64(sr.RSSI)) + 0.1)) * 5
	}
	if isInRange(float64(sr.RSSI), td.Median, 10) {
		score += (1.0 / (math.Abs(td.Median-float64(sr.RSSI)) + 0.1)) * 5
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
