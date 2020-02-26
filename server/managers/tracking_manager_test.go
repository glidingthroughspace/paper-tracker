package managers

import (
	"paper-tracker/models"

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

	Context("Test getMedian", func() {
		It("Calculates Median for an odd amount of values", func() {
			Expect(getMedian(10, 25, 1)).Should(BeNumerically("==", 10.0))
			Expect(getMedian(1, 10, 25)).Should(BeNumerically("==", 10.0))
			Expect(getMedian(-1, 10, 25)).Should(BeNumerically("==", 10.0))
			Expect(getMedian(-99, -10, -25)).Should(BeNumerically("==", -25.0))
		})

		It("Calculates Median for an even amount of values", func() {
			Expect(getMedian(10, 25, 1, 12)).Should(BeNumerically("==", 11.0))
			Expect(getMedian(1, 12, 11, 25)).Should(BeNumerically("==", 11.5))
			Expect(getMedian(-1, 10, -10, 25)).Should(BeNumerically("==", 4.5))
			Expect(getMedian(-99, -10, -25, -12)).Should(BeNumerically("==", -18.5))
		})
	})

	Context("Test getMean", func() {
		It("Calculates the correct Mean", func() {
			Expect(getMean(10, 25, 1)).Should(BeNumerically("==", 12.0))
			Expect(getMean(1, 10, 25)).Should(BeNumerically("==", 12.0))
			Expect(getMean(-1, 10, 25)).Should(BeNumerically("==", 34.0/3.0))
			Expect(getMean(-99, -10, -25)).Should(BeNumerically("==", -134.0/3.0))
		})
	})

	Context("Test ConsolidateScanResults", func() {
		It("Consolidates ScanResults for a single BSSID", func() {
			rssis := []int{-31, -12, -80, -99}
			scanResults := getScanResultsForBSSIDWithRSSIs("AA:BB:CC:DD:EE", rssis...)
			expected := []models.BSSIDTrackingData{
				{BSSID: "AA:BB:CC:DD:EE", Minimum: getMin(rssis...), Maximum: getMax(rssis...), Median: getMedian(rssis...), Mean: getMean(rssis...)},
			}

			Expect(manager.ConsolidateScanResults(scanResults)).To(Equal(expected))
		})

		It("Consolidates ScanResults for two BSSIDs", func() {
			rssis1 := []int{-31, -12, -80, -99}
			scanResults1 := getScanResultsForBSSIDWithRSSIs("AA:BB:CC:DD:EE", rssis1...)
			rssis2 := []int{-40, -33, -17, -22}
			scanResults2 := getScanResultsForBSSIDWithRSSIs("EE:DD:CC:BB:AA", rssis2...)
			expected := []models.BSSIDTrackingData{
				{BSSID: "AA:BB:CC:DD:EE", Minimum: getMin(rssis1...), Maximum: getMax(rssis1...), Median: getMedian(rssis1...), Mean: getMean(rssis1...)},
				{BSSID: "EE:DD:CC:BB:AA", Minimum: getMin(rssis2...), Maximum: getMax(rssis2...), Median: getMedian(rssis2...), Mean: getMean(rssis2...)},
			}

			Expect(manager.ConsolidateScanResults(append(scanResults1, scanResults2...))).To(Equal(expected))
		})
	})
})

func getScanResultsForBSSIDWithRSSIs(bssid string, rssis ...int) []models.ScanResult {
	var scanResults []models.ScanResult
	for _, v := range rssis {
		scanResults = append(scanResults, models.ScanResult{models.TrackerID(1), "TestNetwork", bssid, v})
	}
	return scanResults
}
