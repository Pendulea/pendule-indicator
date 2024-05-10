package indicator

import (
	"errors"
	"strconv"
)

type IndicatorParser struct {
	Indicator Indicator
	Data      []byte
}

func (idn Indicator) Parse(data []byte) *IndicatorParser {
	return &IndicatorParser{
		Indicator: idn,
		Data:      data,
	}
}

func (idn IndicatorParser) CheckEmptyDataError() error {
	if len(idn.Data) == 0 {
		return errors.New("empty data")
	}
	return nil
}

func (p *IndicatorParser) RSI() (RSIValue, error) {
	if err := p.CheckEmptyDataError(); err != nil {
		return 0, err
	}
	if p.Indicator.IsRSI() {
		v, err := strconv.Atoi(string(p.Data))
		if err != nil {
			return 0, err
		}
		return RSIValue(v) / 100, nil
	}
	return 0, errors.New("not rsi indicator")
}

func (p *IndicatorParser) VolatilityIndex() (VolatilityIndexValue, error) {
	if err := p.CheckEmptyDataError(); err != nil {
		return 0, err
	}
	if p.Indicator.IsVolatilityIndex() {
		v, err := strconv.ParseFloat(string(p.Data), 64)
		if err != nil {
			return 0, err
		}
		return VolatilityIndexValue(v), nil
	}
	return 0, errors.New("not volatility index indicator")
}
