package util 

import (
	"strconv"
	"log"
)

func GetDebt2EquityTrend(balancesheetJson map[string]interface{}) interface{} {
	var yearlyDates []string
	var yearlyDebt2EquityRatio []float64
	for _, value := range balancesheetJson["annualReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		yearlyDates = append(yearlyDates, valueCasted["fiscalDateEnding"].(string))
		totalEquity, equityErr := strconv.ParseInt(valueCasted["totalShareholderEquity"].(string), 10, 64)
		totalLiabilities, liabilityErr := strconv.ParseInt(valueCasted["totalLiabilities"].(string), 10, 64)
		if equityErr != nil {
			log.Println(equityErr)
			totalEquity = -1
		}
		if liabilityErr != nil {
			log.Println(liabilityErr)
			totalLiabilities = -1
		}
		yearlyDebt2EquityRatio = append(yearlyDebt2EquityRatio, CalculateRatio(float64(totalLiabilities), float64(totalEquity)))
	}
	var quarterlyDates []string
	var quarterlyDebt2EquityRatio []float64
	for _, value := range balancesheetJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		quarterlyDates = append(quarterlyDates, valueCasted["fiscalDateEnding"].(string))
		totalEquity, equityErr := strconv.ParseInt(valueCasted["totalShareholderEquity"].(string), 10, 64)
		totalLiabilities, liabilityErr := strconv.ParseInt(valueCasted["totalLiabilities"].(string), 10, 64)
		if equityErr != nil {
			log.Println(equityErr)
			totalEquity = -1
		}
		if liabilityErr != nil {
			log.Println(liabilityErr)
			totalLiabilities = -1
		}
		quarterlyDebt2EquityRatio = append(quarterlyDebt2EquityRatio, CalculateRatio(float64(totalLiabilities), float64(totalEquity)))
	}
	ReverseFloatArray(yearlyDebt2EquityRatio)
	ReverseFloatArray(quarterlyDebt2EquityRatio)
	ReverseStringArray(yearlyDates)
	ReverseStringArray(quarterlyDates)
	return struct {
		YearlyDates               []string
		QuarterlyDates            []string
		YearlyDebt2EquityRatio    []float64
		QuarterlyDebt2EquityRatio []float64
	}{
		YearlyDates:               yearlyDates,
		QuarterlyDates:            quarterlyDates,
		YearlyDebt2EquityRatio:    yearlyDebt2EquityRatio,
		QuarterlyDebt2EquityRatio: quarterlyDebt2EquityRatio,
	}
}

func GetOperatingMarginTrend(incomeJson map[string]interface{}) interface{} {
	var yearlyDates []string
	var yearlyOperatingMargin []float64
	for _, value := range incomeJson["annualReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		yearlyDates = append(yearlyDates, valueCasted["fiscalDateEnding"].(string))
		totalReveneue, revErr := strconv.ParseInt(valueCasted["totalRevenue"].(string), 10, 64)
		operatingIncome, opIncomeErr := strconv.ParseInt(valueCasted["operatingIncome"].(string), 10, 64)
		if revErr != nil {
			log.Println(revErr)
			totalReveneue = -1
		}
		if opIncomeErr != nil {
			log.Println(opIncomeErr)
			operatingIncome = -1
		}
		yearlyOperatingMargin = append(yearlyOperatingMargin, CalculateMargin(float64(operatingIncome), float64(totalReveneue)))
	}
	var quarterlyDates []string
	var quarterlyOperatingMargin []float64
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		quarterlyDates = append(quarterlyDates, valueCasted["fiscalDateEnding"].(string))
		totalReveneue, revErr := strconv.ParseInt(valueCasted["totalRevenue"].(string), 10, 64)
		operatingIncome, opIncomeErr := strconv.ParseInt(valueCasted["operatingIncome"].(string), 10, 64)
		if revErr != nil {
			log.Println(revErr)
			totalReveneue = -1
		}
		if opIncomeErr != nil {
			log.Println(opIncomeErr)
			operatingIncome = -1
		}
		quarterlyOperatingMargin = append(quarterlyOperatingMargin, CalculateMargin(float64(operatingIncome), float64(totalReveneue)))
	}
	ReverseFloatArray(yearlyOperatingMargin)
	ReverseFloatArray(quarterlyOperatingMargin)
	ReverseStringArray(yearlyDates)
	ReverseStringArray(quarterlyDates)
	return struct {
		YearlyDates              []string
		QuarterlyDates           []string
		YearlyOperatingMargin    []float64
		QuarterlyOperatingMargin []float64
	}{
		YearlyDates:              yearlyDates,
		QuarterlyDates:           quarterlyDates,
		YearlyOperatingMargin:    yearlyOperatingMargin,
		QuarterlyOperatingMargin: quarterlyOperatingMargin,
	}
}

func GetProfitMarginTrend(incomeJson map[string]interface{}) interface{} {
	var yearlyDates []string
	var yearlyProfitMargin []float64
	for _, value := range incomeJson["annualReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		yearlyDates = append(yearlyDates, valueCasted["fiscalDateEnding"].(string))
		totalReveneue, revErr := strconv.ParseInt(valueCasted["totalRevenue"].(string), 10, 64)
		netIncome, incomeErr := strconv.ParseInt(valueCasted["netIncome"].(string), 10, 64)
		if revErr != nil {
			log.Println(revErr)
			totalReveneue = -1
		}
		if incomeErr != nil {
			log.Println(incomeErr)
			netIncome = -1
		}
		yearlyProfitMargin = append(yearlyProfitMargin, CalculateMargin(float64(netIncome), float64(totalReveneue)))
	}
	var quarterlyDates []string
	var quarterlyProfitMargin []float64
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		quarterlyDates = append(quarterlyDates, valueCasted["fiscalDateEnding"].(string))
		totalReveneue, revErr := strconv.ParseInt(valueCasted["totalRevenue"].(string), 10, 64)
		netIncome, incomeErr := strconv.ParseInt(valueCasted["netIncome"].(string), 10, 64)
		if revErr != nil {
			log.Println(revErr)
			totalReveneue = -1
		}
		if incomeErr != nil {
			log.Println(incomeErr)
			netIncome = -1
		}
		quarterlyProfitMargin = append(quarterlyProfitMargin, CalculateMargin(float64(netIncome), float64(totalReveneue)))
	}
	ReverseFloatArray(yearlyProfitMargin)
	ReverseFloatArray(quarterlyProfitMargin)
	ReverseStringArray(yearlyDates)
	ReverseStringArray(quarterlyDates)
	return struct {
		YearlyDates           []string
		QuarterlyDates        []string
		YearlyProfitMargin    []float64
		QuarterlyProfitMargin []float64
	}{
		YearlyDates:           yearlyDates,
		QuarterlyDates:        quarterlyDates,
		YearlyProfitMargin:    yearlyProfitMargin,
		QuarterlyProfitMargin: quarterlyProfitMargin,
	}
}


func GetCashflowTrend(incomeJson map[string]interface{}) interface{} {
	var yearlyDates []string
	var yearlyCashflow []int64
	for _, value := range incomeJson["annualReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		yearlyDates = append(yearlyDates, valueCasted["fiscalDateEnding"].(string))
		operatingCashflow, opErr := strconv.ParseInt(valueCasted["operatingCashflow"].(string), 10, 64)
		capitalExpenditures, capErr := strconv.ParseInt(valueCasted["capitalExpenditures"].(string), 10, 64)
		if opErr != nil {
			log.Println(opErr)
			operatingCashflow = -1
		}
		if capErr != nil {
			log.Println(capErr)
			capitalExpenditures = -1
		}
		yearlyCashflow = append(yearlyCashflow, operatingCashflow-capitalExpenditures)
	}
	var quarterlyDates []string
	var quarterlyCashflow []int64
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		quarterlyDates = append(quarterlyDates, valueCasted["fiscalDateEnding"].(string))
		operatingCashflow, opErr := strconv.ParseInt(valueCasted["operatingCashflow"].(string), 10, 64)
		capitalExpenditures, capErr := strconv.ParseInt(valueCasted["capitalExpenditures"].(string), 10, 64)
		if opErr != nil {
			log.Println(opErr)
			operatingCashflow = -1
		}
		if capErr != nil {
			log.Println(capErr)
			capitalExpenditures = -1
		}
		quarterlyCashflow = append(quarterlyCashflow, operatingCashflow-capitalExpenditures)
	}
	ReverseIntArray(yearlyCashflow)
	ReverseIntArray(quarterlyCashflow)
	ReverseStringArray(yearlyDates)
	ReverseStringArray(quarterlyDates)
	return struct {
		YearlyDates       []string
		QuarterlyDates    []string
		YearlyCashflow    []int64
		QuarterlyCashflow []int64
	}{
		YearlyDates:       yearlyDates,
		QuarterlyDates:    quarterlyDates,
		YearlyCashflow:    yearlyCashflow,
		QuarterlyCashflow: quarterlyCashflow,
	}
}

func GetRevenueTrend(incomeJson map[string]interface{}) interface{} {
	var yearlyDates []string
	var yearlyRevenue []int64
	for _, value := range incomeJson["annualReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		yearlyDates = append(yearlyDates, valueCasted["fiscalDateEnding"].(string))
		revenue, err := strconv.ParseInt(valueCasted["totalRevenue"].(string), 10, 64)
		if err != nil {
			log.Println(err)
			revenue = -1
		}
		yearlyRevenue = append(yearlyRevenue, revenue)
	}
	var quarterlyDates []string
	var quarterlyRevenue []int64
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		quarterlyDates = append(quarterlyDates, valueCasted["fiscalDateEnding"].(string))
		revenue, err := strconv.ParseInt(valueCasted["totalRevenue"].(string), 10, 64)
		if err != nil {
			log.Println(err)
			revenue = -1
		}
		quarterlyRevenue = append(quarterlyRevenue, revenue)
	}
	ReverseIntArray(yearlyRevenue)
	ReverseIntArray(quarterlyRevenue)
	ReverseStringArray(yearlyDates)
	ReverseStringArray(quarterlyDates)
	return struct {
		YearlyDates      []string
		QuarterlyDates   []string
		YearlyRevenue    []int64
		QuarterlyRevenue []int64
	}{
		YearlyDates:      yearlyDates,
		QuarterlyDates:   quarterlyDates,
		YearlyRevenue:    yearlyRevenue,
		QuarterlyRevenue: quarterlyRevenue,
	}
}