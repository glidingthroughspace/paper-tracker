package managers

import (
	"paper-tracker/models"
	"paper-tracker/utils/collections"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TrackingManager", func() {
	var (
		manager *TrackingManager
	)

	BeforeEach(func() {
		trackingManager = nil

		CreateTrackingManager()
	})

	AfterEach(func() {
	})

	Context("Test ConsolidateScanResults", func() {
		It("Consolidates ScanResults for a single BSSID", func() {
			rssis := []int{-31, -12, -80, -99}
			scanResults := getScanResultsForBSSIDWithRSSIs("AA:BB:CC:DD:EE", rssis...)
			expected := []models.BSSIDTrackingData{
				getBSSIDTrackingDataForBSSIDWithRSSIs("AA:BB:CC:DD:EE", rssis),
			}

			Expect(manager.ConsolidateScanResults(scanResults)).To(Equal(expected))
		})

		It("Consolidates ScanResults for two BSSIDs", func() {
			rssis1 := []int{-31, -12, -80, -99}
			scanResults1 := getScanResultsForBSSIDWithRSSIs("AA:BB:CC:DD:EE", rssis1...)
			rssis2 := []int{-40, -33, -17, -22}
			scanResults2 := getScanResultsForBSSIDWithRSSIs("EE:DD:CC:BB:AA", rssis2...)
			expected := []models.BSSIDTrackingData{
				getBSSIDTrackingDataForBSSIDWithRSSIs("AA:BB:CC:DD:EE", rssis1),
				getBSSIDTrackingDataForBSSIDWithRSSIs("EE:DD:CC:BB:AA", rssis2),
			}

			Expect(manager.ConsolidateScanResults(append(scanResults1, scanResults2...))).Should(ConsistOf(expected))
		})
	})
})

func getScanResultsForBSSIDWithRSSIs(bssid string, rssis ...int) []*models.ScanResult {
	var scanResults []*models.ScanResult
	for _, v := range rssis {
		scanResults = append(scanResults, &models.ScanResult{TrackerID: models.TrackerID(1), SSID: "TestNetwork", BSSID: bssid, RSSI: v})
	}
	return scanResults
}

func getBSSIDTrackingDataForBSSIDWithRSSIs(bssid string, rssis []int) models.BSSIDTrackingData {
	return models.BSSIDTrackingData{
		BSSID:   bssid,
		Minimum: collections.MinOf(rssis...),
		Maximum: collections.MaxOf(rssis...),
		Median:  collections.MedianOf(rssis...),
		Mean:    collections.MeanOf(rssis...),
		Quantiles: models.Quantiles{
			FirstQuartile: collections.FirstQuartileOf(rssis...),
			ThirdQuartile: collections.ThirdQuartileOf(rssis...),
		},
	}

}
