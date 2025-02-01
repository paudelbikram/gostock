package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

//go:embed data/template/template.html
var templatesFS embed.FS

func joinIntArray(arr []int64, sep string) string {
	strArr := make([]string, len(arr))
	for i, num := range arr {
		strArr[i] = strconv.FormatInt(num, 10)
	}
	return "[" + strings.Join(strArr, sep) + "]"
}

func joinFloatArray(arr []float64, sep string) string {
	strArr := make([]string, len(arr))
	for i, num := range arr {
		strArr[i] = strconv.FormatFloat(num, 'f', 3, 64)
	}
	return "[" + strings.Join(strArr, sep) + "]"
}

func joinStringArray(arr []string, sep string) string {
	strArr := make([]string, len(arr))
	for i, num := range arr {
		strArr[i] = "'" + num + "'"
	}
	return "[" + strings.Join(strArr, sep) + "]"
}

func createView(ticker string) {
	// parse overview
	var overviewJson map[string]interface{}
	overviewErr := json.Unmarshal(getFileContent(fmt.Sprintf("./data/cache/%s/overview.json", ticker)), &overviewJson)
	if overviewErr != nil {
		log.Println(overviewErr)
	}
	// parse incomestatement
	var incomeJson map[string]interface{}
	incomeErr := json.Unmarshal(getFileContent(fmt.Sprintf("./data/cache/%s/incomestatement.json", ticker)), &incomeJson)
	if incomeErr != nil {
		log.Println(incomeErr)
	}
	// parse earning
	var earningJson map[string]interface{}
	earningErr := json.Unmarshal(getFileContent(fmt.Sprintf("./data/cache/%s/earning.json", ticker)), &earningJson)
	if earningErr != nil {
		log.Println(earningErr)
	}
	// parse cashflow
	var cashflowJson map[string]interface{}
	cashflowErr := json.Unmarshal(getFileContent(fmt.Sprintf("./data/cache/%s/cashflow.json", ticker)), &cashflowJson)
	if cashflowErr != nil {
		log.Println(cashflowErr)
	}
	// parse balancesheet
	var balancesheetJson map[string]interface{}
	balancesheetErr := json.Unmarshal(getFileContent(fmt.Sprintf("./data/cache/%s/balancesheet.json", ticker)), &balancesheetJson)
	if balancesheetErr != nil {
		log.Println(balancesheetErr)
	}
	funcMap := template.FuncMap{
		"joinStringArray": joinStringArray,
		"joinIntArray":    joinIntArray,
		"joinFloatArray":  joinFloatArray,
	}
	stockTemplate := template.New("template.html").Funcs(funcMap)
	stockTemplate, templateErr := stockTemplate.ParseFS(templatesFS, "data/template/template.html")
	if templateErr != nil {
		log.Println(templateErr)
	}
	revenueTrend := getRevenueTrend(incomeJson)
	cashflowTrend := getCashflowTrend(cashflowJson)
	profitMarginTrend := getProfitMarginTrend(incomeJson)
	operatingMarginTrend := getOperatingMarginTrend(incomeJson)
	debt2equityRatioTrend := getDebt2EquityTrend(balancesheetJson)
	data := map[string]interface{}{
		"ticker":                ticker,
		"overview":              overviewJson,
		"income":                incomeJson,
		"earning":               earningJson,
		"cashflow":              cashflowJson,
		"balancesheet":          balancesheetJson,
		"revenueTrend":          revenueTrend,
		"cashflowTrend":         cashflowTrend,
		"profitMarginTrend":     profitMarginTrend,
		"operatingMarginTrend":  operatingMarginTrend,
		"debt2equityRatioTrend": debt2equityRatioTrend,
	}
	os.MkdirAll("./data/output", os.ModePerm)
	outputfile, outputfileCreationErr := os.Create(fmt.Sprintf("./data/output/%s.html", ticker))
	if outputfileCreationErr != nil {
		log.Println(outputfileCreationErr)
	}
	defer outputfile.Close()

	outputErr := stockTemplate.Execute(outputfile, data)
	if outputErr != nil {
		log.Println(outputErr)
	}
}

func calculateMargin(numerator float64, denominator float64) float64 {
	return (numerator / denominator) * 100
}

func calculateRatio(numerator float64, denominator float64) float64 {
	return numerator / denominator
}

func getDebt2EquityTrend(balancesheetJson map[string]interface{}) interface{} {
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
		yearlyDebt2EquityRatio = append(yearlyDebt2EquityRatio, calculateRatio(float64(totalLiabilities), float64(totalEquity)))
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
		quarterlyDebt2EquityRatio = append(quarterlyDebt2EquityRatio, calculateRatio(float64(totalLiabilities), float64(totalEquity)))
	}
	reverseFloatArray(yearlyDebt2EquityRatio)
	reverseFloatArray(quarterlyDebt2EquityRatio)
	reverseStringArray(yearlyDates)
	reverseStringArray(quarterlyDates)
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

func getOperatingMarginTrend(incomeJson map[string]interface{}) interface{} {
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
		yearlyOperatingMargin = append(yearlyOperatingMargin, calculateMargin(float64(operatingIncome), float64(totalReveneue)))
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
		quarterlyOperatingMargin = append(quarterlyOperatingMargin, calculateMargin(float64(operatingIncome), float64(totalReveneue)))
	}
	reverseFloatArray(yearlyOperatingMargin)
	reverseFloatArray(quarterlyOperatingMargin)
	reverseStringArray(yearlyDates)
	reverseStringArray(quarterlyDates)
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

func getProfitMarginTrend(incomeJson map[string]interface{}) interface{} {
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
		yearlyProfitMargin = append(yearlyProfitMargin, calculateMargin(float64(netIncome), float64(totalReveneue)))
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
		quarterlyProfitMargin = append(quarterlyProfitMargin, calculateMargin(float64(netIncome), float64(totalReveneue)))
	}
	reverseFloatArray(yearlyProfitMargin)
	reverseFloatArray(quarterlyProfitMargin)
	reverseStringArray(yearlyDates)
	reverseStringArray(quarterlyDates)
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

/*

func getDebtToEquityTrend(incomeJson map[string]interface{}) interface{} {
}
*/

func getCashflowTrend(incomeJson map[string]interface{}) interface{} {
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
	reverseIntArray(yearlyCashflow)
	reverseIntArray(quarterlyCashflow)
	reverseStringArray(yearlyDates)
	reverseStringArray(quarterlyDates)
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

func getRevenueTrend(incomeJson map[string]interface{}) interface{} {
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
	reverseIntArray(yearlyRevenue)
	reverseIntArray(quarterlyRevenue)
	reverseStringArray(yearlyDates)
	reverseStringArray(quarterlyDates)
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

func reverseIntArray(arr []int64) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func reverseFloatArray(arr []float64) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func reverseStringArray(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
