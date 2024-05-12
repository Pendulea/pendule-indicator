package indicator

import "strings"

type Indicator string

const RSI14 Indicator = "RSI-14"
const RSI9 Indicator = "RSI-9"
const VI1 Indicator = "VI-1"

var ALL_INDICATORS = []Indicator{
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

func ParseIndicators(list string, supportedList []Indicator) []Indicator {
	if list == "" {
		return nil
	}
	if list == "*" {
		return supportedList
	}

	allowedIndicators := strings.Split(list, ",")
	finalList := make([]Indicator, 0) // Use a separate slice for the results to avoid duplication

	for _, indicator := range allowedIndicators {
		if strings.Contains(indicator, "*") {
			// Handle wildcard entries
			sp := strings.Split(indicator, "*")
			if len(sp) != 2 {
				continue // Skip invalid formats
			}
			prefix, suffix := sp[0], sp[1]
			for _, ind := range supportedList {
				if strings.HasPrefix(string(ind), strings.ToUpper(prefix)) && strings.HasSuffix(string(ind), strings.ToUpper(suffix)) {
					finalList = appendIfNotExists(finalList, ind)
				}
			}
		} else {
			// Handle exact matches
			if contains(supportedList, Indicator(indicator)) {
				finalList = appendIfNotExists(finalList, Indicator(indicator))
			}
		}
	}

	return finalList
}

// Helper function to check if a slice contains a specific string
func contains(slice []Indicator, str Indicator) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// Append to the slice if the element does not already exist
func appendIfNotExists(slice []Indicator, str Indicator) []Indicator {
	if !contains(slice, str) {
		slice = append(slice, Indicator(strings.ToUpper(string(str))))
	}
	return slice
}
