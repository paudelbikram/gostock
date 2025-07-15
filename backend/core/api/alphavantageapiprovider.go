package api

import (
	"fmt"
	"gostock/backend/core/util"
	"encoding/json"
	"log"
)

type AlphaVantageApiProvider struct {
	apiKey       string
	providerName string
}

func NewAlphaVantageApiProvider() *AlphaVantageApiProvider {
	provider := &AlphaVantageApiProvider{}
	provider.providerName = "alpha-vantage"
	provider.apiKey = util.GetApiKey(provider.providerName)
	return provider
}

func (a *AlphaVantageApiProvider) GetOverviewUrl(ticker string) string {
	return fmt.Sprintf("https://www.alphavantage.co/query?function=OVERVIEW&symbol=%s&apikey=%s",
		ticker,
		a.apiKey,
	)
}

func (a *AlphaVantageApiProvider) GetIncomeStatementUrl(ticker string) string {
	return fmt.Sprintf("https://www.alphavantage.co/query?function=INCOME_STATEMENT&symbol=%s&apikey=%s",
		ticker,
		a.apiKey,
	)
}

func (a *AlphaVantageApiProvider) GetBalanceSheetUrl(ticker string) string {
	return fmt.Sprintf("https://www.alphavantage.co/query?function=BALANCE_SHEET&symbol=%s&apikey=%s",
		ticker,
		a.apiKey,
	)
}

func (a *AlphaVantageApiProvider) GetEarningUrl(ticker string) string {
	return fmt.Sprintf("https://www.alphavantage.co/query?function=EARNINGS&symbol=%s&apikey=%s",
		ticker,
		a.apiKey,
	)
}

func (a *AlphaVantageApiProvider) GetCashflowUrl(ticker string) string {
	return fmt.Sprintf("https://www.alphavantage.co/query?function=CASH_FLOW&symbol=%s&apikey=%s",
		ticker,
		a.apiKey,
	)
}

func (a *AlphaVantageApiProvider) GetData(ticker string) map[string]interface{} {
	//getting overview data
	overviewData := util.GetCacheData(a.providerName, ticker, "overview")
	if overviewData == "" {
		overviewData = util.SetCacheData(a.providerName, ticker, "overview",
			util.Get(a.GetOverviewUrl(ticker)))
	}
	//getting incomestatement data
	incomeData := util.GetCacheData(a.providerName, ticker, "incomestatement")
	if incomeData == "" {
		incomeData = util.SetCacheData(a.providerName, ticker, "incomestatement",
			util.Get(a.GetIncomeStatementUrl(ticker)))
	}
	//getting balancesheet data
	balancesheetData := util.GetCacheData(a.providerName, ticker, "balancesheet")
	if balancesheetData == "" {
		balancesheetData = util.SetCacheData(a.providerName, ticker, "balancesheet",
			util.Get(a.GetBalanceSheetUrl(ticker)))
	}
	//getting earning data
	earningData := util.GetCacheData(a.providerName, ticker, "earning")
	if earningData == "" {
		earningData = util.SetCacheData(a.providerName, ticker, "earning",
			util.Get(a.GetEarningUrl(ticker)))
	}
	//getting cashflow data
	cashflowData := util.GetCacheData(a.providerName, ticker, "cashflow")
	if cashflowData == "" {
		cashflowData = util.SetCacheData(a.providerName, ticker, "cashflow",
			util.Get(a.GetCashflowUrl(ticker)))
	}

	var overviewJson map[string]interface{}
	overeviewErr := json.Unmarshal([]byte(overviewData), &overviewJson)
	if overeviewErr != nil {
		log.Println(overeviewErr)
	}
	var cashflowJson map[string]interface{}
	cashflowErr := json.Unmarshal([]byte(cashflowData), &cashflowJson)
	if cashflowErr != nil {
		log.Println(cashflowErr)
	}
	var balancesheetJson map[string]interface{}
	balancesheetErr := json.Unmarshal([]byte(balancesheetData), &balancesheetJson)
	if balancesheetErr != nil {
		log.Println(balancesheetErr)
	}
	var incomeJson map[string]interface{}
	incomeErr := json.Unmarshal([]byte(incomeData), &incomeJson)
	if incomeErr != nil {
		log.Println(incomeErr)
	}
	var earningJson map[string]interface{}
	earningErr := json.Unmarshal([]byte(earningData), &earningJson)
	if earningErr != nil {
		log.Println(earningErr)
	}

	revenueTrend := util.GetRevenueTrend(incomeJson)
	cashflowTrend := util.GetCashflowTrend(cashflowJson)
	profitMarginTrend := util.GetProfitMarginTrend(incomeJson)
	operatingMarginTrend := util.GetOperatingMarginTrend(incomeJson)
	debt2equityRatioTrend := util.GetDebt2EquityTrend(balancesheetJson)
	return map[string]interface{}{
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
}
