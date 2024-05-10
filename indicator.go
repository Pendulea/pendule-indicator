package indicator

type Indicator string

const RSI14 Indicator = "RSI-14"
const RSI9 Indicator = "RSI-9"
const VI1 Indicator = "VI-1"

var INDICATORS = []Indicator{
	RSI14,
	RSI9,
	VI1,
}

func (id Indicator) IsRSI() bool {
	return id == RSI14 || id == RSI9
}

func (id Indicator) IsVolatilityIndex() bool {
	return id == VI1
}
