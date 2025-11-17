package util

import (
	"gostock/backend/logger"
	"strconv"
	"go.uber.org/zap"
)

func GetFloatValue(numInString string) (float64, error) {
	if numInString == "None" {
		return 0.0, nil
	}
	return strconv.ParseFloat(numInString, 64)
}

func GetIntValue(numInString string) (int64, error) {
	if numInString == "None" {
		return 0, nil
	}
	return strconv.ParseInt(numInString, 10, 64)
}

func GetDebt2EquityTrend(balancesheetJson map[string]interface{}) interface{} {
	var yearly []interface{}
	for _, value := range balancesheetJson["annualReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		totalEquity, equityErr := GetIntValue(valueCasted["totalShareholderEquity"].(string))
		totalLiabilities, liabilityErr := GetIntValue(valueCasted["totalLiabilities"].(string))
		if equityErr != nil {
			logger.Log.Error("Error",
				zap.Error(equityErr),
			)
			totalEquity = 0
		}
		if liabilityErr != nil {
			logger.Log.Error("Error",
				zap.Error(liabilityErr),
			)
			totalLiabilities = 0
		}
		yearly = append(yearly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": CalculateRatio(float64(totalLiabilities), float64(totalEquity))})
	}
	var quarterly []interface{}
	for _, value := range balancesheetJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		totalEquity, equityErr := GetIntValue(valueCasted["totalShareholderEquity"].(string))
		totalLiabilities, liabilityErr := GetIntValue(valueCasted["totalLiabilities"].(string))
		if equityErr != nil {
			logger.Log.Error("Error",
				zap.Error(equityErr),
			)
			totalEquity = 0
		}
		if liabilityErr != nil {
			logger.Log.Error("Error",
				zap.Error(liabilityErr),
			)
			totalLiabilities = 0
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
		totalReveneue, revErr := GetIntValue(valueCasted["totalRevenue"].(string))
		operatingIncome, opIncomeErr := GetIntValue(valueCasted["operatingIncome"].(string))
		if revErr != nil {
			logger.Log.Error("Error",
				zap.Error(revErr),
			)
			totalReveneue = 0
		}
		if opIncomeErr != nil {
			logger.Log.Error("Error",
				zap.Error(opIncomeErr),
			)
			operatingIncome = 0
		}
		yearly = append(yearly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": CalculateMargin(float64(operatingIncome), float64(totalReveneue))})
	}
	var quarterly []interface{}
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		totalReveneue, revErr := GetIntValue(valueCasted["totalRevenue"].(string))
		operatingIncome, opIncomeErr := GetIntValue(valueCasted["operatingIncome"].(string))
		if revErr != nil {
			logger.Log.Error("Error",
				zap.Error(revErr),
			)
			totalReveneue = 0
		}
		if opIncomeErr != nil {
			logger.Log.Error("Error",
				zap.Error(opIncomeErr),
			)
			operatingIncome = 0
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
		totalReveneue, revErr := GetIntValue(valueCasted["totalRevenue"].(string))
		netIncome, incomeErr := GetIntValue(valueCasted["netIncome"].(string))
		if revErr != nil {
			logger.Log.Error("Error",
				zap.Error(revErr),
			)
			totalReveneue = 0
		}
		if incomeErr != nil {
			logger.Log.Error("Error",
				zap.Error(incomeErr),
			)
			netIncome = 0
		}
		yearly = append(yearly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": CalculateMargin(float64(netIncome), float64(totalReveneue))})
	}
	var quarterly []interface{}
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		totalReveneue, revErr := GetIntValue(valueCasted["totalRevenue"].(string))
		netIncome, incomeErr := GetIntValue(valueCasted["netIncome"].(string))
		if revErr != nil {
			logger.Log.Error("Error",
				zap.Error(revErr),
			)
			totalReveneue = 0
		}
		if incomeErr != nil {
			logger.Log.Error("Error",
				zap.Error(incomeErr),
			)
			netIncome = 0
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
		operatingCashflow, opErr := GetFloatValue(valueCasted["operatingCashflow"].(string))
		capitalExpenditures, capErr := GetFloatValue(valueCasted["capitalExpenditures"].(string))
		if opErr != nil {
			logger.Log.Error("Error",
				zap.Error(opErr),
			)
			operatingCashflow = 0
		}
		if capErr != nil {
			logger.Log.Error("Error",
				zap.Error(capErr),
			)
			capitalExpenditures = 0
		}
		yearly = append(yearly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": operatingCashflow - capitalExpenditures})
	}
	var quarterly []interface{}
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		operatingCashflow, opErr := GetIntValue(valueCasted["operatingCashflow"].(string))
		capitalExpenditures, capErr := GetIntValue(valueCasted["capitalExpenditures"].(string))
		if opErr != nil {
			logger.Log.Error("Error",
				zap.Error(opErr),
			)
			operatingCashflow = 0
		}
		if capErr != nil {
			logger.Log.Error("Error",
				zap.Error(capErr),
			)
			capitalExpenditures = 0
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
		revenue, err := GetIntValue(valueCasted["totalRevenue"].(string))
		if err != nil {
			logger.Log.Error("Error",
				zap.Error(err),
			)
			revenue = 0
		}
		yearly = append(yearly, map[string]interface{}{"key": valueCasted["fiscalDateEnding"].(string), "value": revenue})
	}
	var quarterly []interface{}
	for _, value := range incomeJson["quarterlyReports"].([]interface{}) {
		valueCasted := value.(map[string]interface{})
		revenue, err := GetIntValue(valueCasted["totalRevenue"].(string))
		if err != nil {
			logger.Log.Error("Error",
				zap.Error(err),
			)
			revenue = 0
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
