package util

func CalculateMargin(numerator float64, denominator float64) float64 {
	return (numerator / denominator) * 100
}

func CalculateRatio(numerator float64, denominator float64) float64 {
	return numerator / denominator
}