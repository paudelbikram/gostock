package util

func CalculateMargin(numerator float64, denominator float64) float64 {
	if denominator == 0.0 {
		return 0.0
	}
	return (numerator / denominator) * 100
}

func CalculateRatio(numerator float64, denominator float64) float64 {
	if denominator == 0.0 {
		return 0.0
	}
	return numerator / denominator
}
