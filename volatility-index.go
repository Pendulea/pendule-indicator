package indicator

import (
	"math"
	"strconv"

	pcommon "github.com/pendulea/pendule-common"
)

type VolatilityIndexValue float64

func (f VolatilityIndexValue) Bytes() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(f), 'f', -1, 64)), nil
}

func CalculateVolatilityIndex(prevMark *VolatilityIndexValue, currentTick pcommon.TickTime, percentChangeTarget float64) *VolatilityIndexValue {
	if prevMark == nil {
		v := VolatilityIndexValue(currentTick.Close)
		return &v
	}

	percentChange := math.Abs((currentTick.Low - float64(*prevMark)) / float64(*prevMark) * 100)
	if percentChange >= percentChangeTarget {
		v := VolatilityIndexValue(currentTick.Low)
		return &v
	}

	percentChange2 := math.Abs((currentTick.High - float64(*prevMark)) / float64(*prevMark) * 100)
	if percentChange2 >= percentChangeTarget {
		v := VolatilityIndexValue(currentTick.High)
		return &v
	}

	return nil
}
