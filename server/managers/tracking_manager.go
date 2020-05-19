package managers

import (
	"math"
	"paper-tracker/config"
	"paper-tracker/models"
	"paper-tracker/utils/collections"

	log "github.com/sirupsen/logrus"
)

var trackingManager TrackingManager

type TrackingManager interface {
	GetRoomMatchingBest(scoredRooms []map[models.RoomID]float64) models.RoomID
	ConsolidateScanResults(scanResults []*models.ScanResult) []models.BSSIDTrackingData
	ScoreRoomsForScanResults(rooms []*models.Room, scanResults []*models.ScanResult) map[models.RoomID]float64
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

func (tm *TrackingManagerImpl) GetRoomMatchingBest(scoredRooms []map[models.RoomID]float64) models.RoomID {
	scores := make(map[models.RoomID]float64)
	for _, sr := range scoredRooms {
		for room, score := range sr {
			scores[room] += score
		}
	}

	for room, score := range scores {
		scores[room] = score / float64(len(scoredRooms))
	}

	var bestScore float64 = 0.0
	var bestMatch models.RoomID = -1

	threshold := config.GetFloat64(config.KeyTrackingScoreThreshold)

	log.Debug(threshold)

	for room, score := range scores {
		if score < threshold {
			log.Debugf("Room %d not considerd, score is lower than threshold (%f < %f)", room, score, threshold)
			continue
		}
		log.Debugf("Room %d scored for matching: %f", room, score)
		if bestMatch == -1 || score > bestScore {
			bestScore = score
			bestMatch = room
		}
	}
	if bestMatch != -1 {
		log.Debugf("Selected the best match to be: %d", bestMatch)
	} else {
		log.Debug("Did not find a matching room")
	}
	return bestMatch
}

func (tm *TrackingManagerImpl) ScoreRoomsForScanResults(rooms []*models.Room, scanResults []*models.ScanResult) map[models.RoomID]float64 {
	scoredRooms := make(map[models.RoomID]float64)
	for _, room := range rooms {
		scoredRooms[room.ID] = tm.ScoreRoomForScanResults(room, scanResults)
	}
	return scoredRooms
}

// ScoreRoomForScanResults calculates a score of how likely the scan results are from the given
// room. The larger the returned score, the greater the likelyness of match.
func (*TrackingManagerImpl) ScoreRoomForScanResults(room *models.Room, scanResults []*models.ScanResult) float64 {
	score := 0.0
	for _, trackingData := range room.TrackingData {
		srs := getScanResultsForBSSID(trackingData.BSSID, scanResults)
		if len(srs) == 0 {
			score -= math.Abs(trackingData.Mean / 100.0)
		}
		for _, sr := range srs {
			score += getScoreForScanResultAndTrackingData(trackingData, sr)
		}
	}
	return score
}

// getScoreForScanResultAndTrackingData calculates the score for a single scan result and the given
// tracking data.
func getScoreForScanResultAndTrackingData(td models.BSSIDTrackingData, sr *models.ScanResult) float64 {
	score := 0.0
	rssi := float64(sr.RSSI)
	rssiFactor := math.Abs(rssi / 100.0)
	rangeForMean := config.GetFloat64(config.KeyTrackingRangeForMean)
	rangeForMedian := config.GetFloat64(config.KeyTrackingRangeForMedian)
	if sr.RSSI <= td.Maximum && sr.RSSI >= td.Minimum {
		score += config.GetFloat64(config.KeyTrackingScoreInMinMaxRangeFactor) * rssiFactor
	}
	if rssi < td.Quantiles.ThirdQuartile && rssi > td.Quantiles.FirstQuartile {
		score += config.GetFloat64(config.KeyTrackingScoreInQuartilesFactor) * rssiFactor
	}
	if isInRange(rssi, td.Mean, rangeForMean) {
		d := math.Abs(td.Mean - rssi)
		score += math.Abs(d-rangeForMean) * config.GetFloat64(config.KeyTrackingScoreMeanFactor) * rssiFactor
	}
	if isInRange(rssi, td.Median, rangeForMean) {
		d := math.Abs(td.Median - rssi)
		score += math.Abs(d-rangeForMedian) * config.GetFloat64(config.KeyTrackingScoreMedianFactor) * rssiFactor
	}
	return score
}

// isInRange returns whether the given actual number is in the range of the wanted number +/- delta
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
