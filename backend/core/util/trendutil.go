package util

import (
	"log"
	"strconv"
)

func GetDebt2EquityTrend(balancesheetJson map[string]interface{}) interface{} {
	var yearly []interface{}
	for _, value := range balancesheetJson["annualReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
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
		yearly = append(yearly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": CalculateRatio(float64(totalLiabilities), float64(totalEquity))})
	}
	var quarterly []interface{}
	for _, value := range balancesheetJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
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
		quarterly = append(quarterly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": CalculateRatio(float64(totalLiabilities), float64(totalEquity))})
	}
	return struct {
		Yearly    []interface{}
		Quarterly []interface{}
	}{
		Yearly:    yearly,
		Quarterly: quarterly,
	}
}

func GetOperatingMarginTrend(incomeJson map[string]interface{}) interface{} {
	var yearly []interface{}
	for _, value := range incomeJson["annualReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
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
		yearly = append(yearly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": CalculateMargin(float64(operatingIncome), float64(totalReveneue))})
	}
	var quarterly []interface{}
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
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
		quarterly = append(quarterly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": CalculateMargin(float64(operatingIncome), float64(totalReveneue))})
	}
	return struct {
		Yearly    []interface{}
		Quarterly []interface{}
	}{
		Yearly:    yearly,
		Quarterly: quarterly,
	}
}

func GetProfitMarginTrend(incomeJson map[string]interface{}) interface{} {
	var yearly []interface{}
	for _, value := range incomeJson["annualReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
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
		yearly = append(yearly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": CalculateMargin(float64(netIncome), float64(totalReveneue))})
	}
	var quarterly []interface{}
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
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
		quarterly = append(quarterly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": CalculateMargin(float64(netIncome), float64(totalReveneue))})
	}
	return struct {
		Yearly    []interface{}
		Quarterly []interface{}
	}{
		Yearly:    yearly,
		Quarterly: quarterly,
	}
}

func GetCashflowTrend(incomeJson map[string]interface{}) interface{} {
	var yearly []interface{}
	for _, value := range incomeJson["annualReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
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
		yearly = append(yearly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": operatingCashflow - capitalExpenditures})
	}
	var quarterly []interface{}
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
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
		quarterly = append(quarterly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": operatingCashflow - capitalExpenditures})
	}
	return struct {
		Yearly    []interface{}
		Quarterly []interface{}
	}{
		Yearly:    yearly,
		Quarterly: quarterly,
	}
}

func GetRevenueTrend(incomeJson map[string]interface{}) interface{} {
	var yearly []interface{}
	for _, value := range incomeJson["annualReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		revenue, err := strconv.ParseInt(valueCasted["totalRevenue"].(string), 10, 64)
		if err != nil {
			log.Println(err)
			revenue = -1
		}
		yearly = append(yearly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": revenue})
	}
	var quarterly []interface{}
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		revenue, err := strconv.ParseInt(valueCasted["totalRevenue"].(string), 10, 64)
		if err != nil {
			log.Println(err)
			revenue = -1
		}
		quarterly = append(quarterly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": revenue})
	}
	return struct {
		Yearly    []interface{}
		Quarterly []interface{}
	}{
		Yearly:    yearly,
		Quarterly: quarterly,
	}
}
