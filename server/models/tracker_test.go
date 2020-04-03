package models

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tracker", func() {
	Context("Test GetSecondsToNextPoll", func() {
		It("GetSecondsToNextPoll returns positive values if next poll is ahead", func() {
			tracker := &Tracker{LastPoll: time.Now().Add(time.Second * -5), LastSleepTimeSec: 10}
			Expect(tracker.GetSecondsToNextPoll()).To(And(BeNumerically(">", 0), BeNumerically("<=", 5)))
		})

		It("GetSecondsToNextPoll returns zero if next poll was in past", func() {
			tracker := &Tracker{LastPoll: time.Now().Add(time.Second * -15), LastSleepTimeSec: 10}
			Expect(tracker.GetSecondsToNextPoll()).To(Equal(0))
		})
	})
})
