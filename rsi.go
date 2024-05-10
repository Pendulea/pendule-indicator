package indicator

import (
	"math"
	"strconv"

	pcommon "github.com/pendulea/pendule-common"
)

const MIN_RSI_PERIOD = 9

type RSIValue float64

func (f RSIValue) Bytes() ([]byte, error) {
	v := int(f * 100)
	return []byte(strconv.Itoa(v)), nil
}

func CalculateRSI(lastTicks pcommon.TickTimeArray) RSIValue {
	period := len(lastTicks) - 1
	if period < MIN_RSI_PERIOD {
		return -1 // Not enough data
	}

	gains := 0.0
	losses := 0.0

	// First, calculate initial average gains and losses
	for i := 1; i <= period; i++ {
		change := lastTicks[i].Close - lastTicks[i-1].Close
		if change > 0 {
			gains += change
		} else {
			losses -= change // Losses are positive numbers
		}
	}

	averageGain := gains / float64(period)
	averageLoss := losses / float64(period)

	// Calculate RSI using smoothed moving averages
	for i := period + 1; i < period; i++ {
		change := lastTicks[i].Close - lastTicks[i-1].Close
		if change > 0 {
			gains = change
		} else {
			losses = -change // Losses are positive numbers
		}

		// Apply smoothing formula
		averageGain = (averageGain*(float64(period)-1) + gains) / float64(period)
		averageLoss = (averageLoss*(float64(period)-1) + losses) / float64(period)
	}

	rs := averageGain / averageLoss
	if math.IsNaN(rs) {
		return -1
	}
	rsi := 100 - (100 / (1 + rs))
	return RSIValue(rsi)
}
