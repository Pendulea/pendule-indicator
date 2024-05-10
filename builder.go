package indicator

import (
	"errors"
	"strconv"
	"strings"

	pcommon "github.com/pendulea/pendule-common"
)

type TickDataMap map[int64][]byte

type TickData struct {
	Data []byte
	Time int64
}

type IndicatorDataBuilderDB struct {
	indicator            Indicator
	prevTicks            pcommon.TickTimeArray
	MinRequiredPrevTicks int
	results              []TickData
	MinRequiredResults   int
}

func NewIndicatorDataBuilderDB(indicator Indicator) *IndicatorDataBuilderDB {
	switch indicator {
	case RSI14:
		return &IndicatorDataBuilderDB{
			MinRequiredPrevTicks: 14,
			MinRequiredResults:   0,
			indicator:            indicator,
		}
	case RSI9:
		return &IndicatorDataBuilderDB{
			MinRequiredPrevTicks: 9,
			MinRequiredResults:   0,
			indicator:            indicator,
		}
	case VI1:
		return &IndicatorDataBuilderDB{
			MinRequiredPrevTicks: 0,
			MinRequiredResults:   1,
			indicator:            indicator,
		}
	default:
		return nil
	}
}

func (b *IndicatorDataBuilderDB) addPrevTick(tick pcommon.TickTime) {
	if b.MinRequiredPrevTicks > 0 {
		b.prevTicks = append(b.prevTicks, tick)
		if len(b.prevTicks) > b.MinRequiredPrevTicks {
			b.prevTicks = b.prevTicks[1:]
		}
	}
}

func (b *IndicatorDataBuilderDB) addResult(result TickData) {
	if b.MinRequiredResults > 0 {
		b.results = append(b.results, result)
		if len(b.results) > b.MinRequiredResults {
			b.results = b.results[1:]
		}
	}
}

func (b *IndicatorDataBuilderDB) BuildData(tick pcommon.TickTime) ([]byte, error) {
	// Defers execution of adding the tick until the rest of the function is completed.
	result := TickData{Time: tick.Time}
	defer func() {
		b.addPrevTick(tick)
	}()

	// Initializes result with current time.
	var err error

	// Handle indicators for RSI calculations.
	if b.indicator.IsRSI() {
		if len(b.prevTicks) >= b.MinRequiredPrevTicks {
			rsi := CalculateRSI(append(b.prevTicks, tick))
			if rsi >= 0 {
				result.Data, err = rsi.Bytes()
			}
		}
		return result.Data, err
	}

	if b.indicator.IsVolatilityIndex() {
		parameterInt, err := strconv.Atoi(strings.Split(string(b.indicator), "-")[1])
		if err != nil {
			return nil, err
		}
		var prevResult *VolatilityIndexValue = nil
		if len(b.results) > 0 {
			result := b.results[len(b.results)-1]
			val, err := b.indicator.Parse(result.Data).VolatilityIndex()
			if err != nil {
				return nil, err
			}
			prevResult = &val
		}
		viValue := CalculateVolatilityIndex(prevResult, tick, float64(parameterInt))
		if viValue != nil {
			result.Data, err = viValue.Bytes()
			b.addResult(result)
		}
		return result.Data, err
	}
	return nil, errors.New("indicator not handled")
}
