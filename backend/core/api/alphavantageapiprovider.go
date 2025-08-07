package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"gostock/backend/core/util"
	"gostock/backend/logger"

	"go.uber.org/zap"
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

func (a *AlphaVantageApiProvider) isDataValid(data map[string]interface{}) bool {
	//no data
	if len(data) == 0 {
		return false
	}
	//API warning
	if _, ok := data["Information"]; ok {
		return false
	}
	return true
}

func (a *AlphaVantageApiProvider) GetData(ticker string) (map[string]interface{}, error) {
	//init
	var overviewJson map[string]interface{}
	var cashflowJson map[string]interface{}
	var balancesheetJson map[string]interface{}
	var incomeJson map[string]interface{}
	var earningJson map[string]interface{}

	//getting overview data
	overviewData, err := util.GetCacheData(a.providerName, ticker, "overview")
	if err != nil {
		logger.Log.Info("Cache not found",
			zap.Error(err),
		)
	}
	if overviewData == "" {
		overviewData = util.Get(a.GetOverviewUrl(ticker))
	}
	overeviewErr := json.Unmarshal([]byte(overviewData), &overviewJson)
	if overeviewErr != nil {
		logger.Log.Error("Failed retrieving overview data",
			zap.Error(err),
		)
		return nil, overeviewErr
	}
	if !a.isDataValid(overviewJson) {
		logger.Log.Error("Invalid overview data",
			zap.Any("data", overviewJson),
		)
		return nil, errors.New("invalid overview data")
	}
	_, err = util.SetCacheData(a.providerName, ticker, "overview",
		overviewData)
	if err != nil {
		logger.Log.Info("Failed setting cache for overview data",
			zap.Error(err),
		)
	}

	//getting incomestatement data
	incomeData, err := util.GetCacheData(a.providerName, ticker, "incomestatement")
	if err != nil {
		logger.Log.Info("Cache not found",
			zap.Error(err),
		)
	}
	if incomeData == "" {
		incomeData = util.Get(a.GetIncomeStatementUrl(ticker))
	}
	incomeErr := json.Unmarshal([]byte(incomeData), &incomeJson)
	if incomeErr != nil {
		logger.Log.Error("Failed retrieving income data",
			zap.Error(err),
		)
		return nil, incomeErr
	}
	if !a.isDataValid(incomeJson) {
		logger.Log.Error("Invalid income data",
			zap.Any("data", incomeJson),
		)
		return nil, errors.New("invalid income data")
	}
	_, err = util.SetCacheData(a.providerName, ticker, "incomestatement",
		incomeData)
	if err != nil {
		logger.Log.Info("Failed setting cache for income data",
			zap.Error(err),
		)
	}

	//getting balancesheet data
	balancesheetData, err := util.GetCacheData(a.providerName, ticker, "balancesheet")
	if err != nil {
		logger.Log.Info("Cache not found",
			zap.Error(err),
		)
	}
	if balancesheetData == "" {
		balancesheetData = util.Get(a.GetBalanceSheetUrl(ticker))
	}
	balancesheetErr := json.Unmarshal([]byte(balancesheetData), &balancesheetJson)
	if balancesheetErr != nil {
		logger.Log.Error("Failed retrieving balancesheet data",
			zap.Error(err),
		)
		return nil, balancesheetErr
	}
	if !a.isDataValid(balancesheetJson) {
		logger.Log.Error("Invalid balancesheet data",
			zap.Any("data", balancesheetJson),
		)
		return nil, errors.New("invalid balancesheet data")
	}

	_, err = util.SetCacheData(a.providerName, ticker, "balancesheet",
		balancesheetData)
	if err != nil {
		logger.Log.Info("Failed setting cache for balancesheet data",
			zap.Error(err),
		)
	}

	//getting earning data
	earningData, err := util.GetCacheData(a.providerName, ticker, "earning")
	if err != nil {
		logger.Log.Info("Cache not found",
			zap.Error(err),
		)
	}
	if earningData == "" {
		earningData = util.Get(a.GetEarningUrl(ticker))
	}
	earningErr := json.Unmarshal([]byte(earningData), &earningJson)
	if earningErr != nil {
		logger.Log.Error("Failed retrieving earning data",
			zap.Error(err),
		)
		return nil, earningErr
	}
	if !a.isDataValid(earningJson) {
		logger.Log.Error("Invalid earning data",
			zap.Any("data", earningJson),
		)
		return nil, errors.New("invalid earning data")
	}
	_, err = util.SetCacheData(a.providerName, ticker, "earning",
		earningData)
	if err != nil {
		logger.Log.Info("Failed setting cache for earning data",
			zap.Error(err),
		)
	}

	//getting cashflow data
	cashflowData, err := util.GetCacheData(a.providerName, ticker, "cashflow")
	if err != nil {
		logger.Log.Info("Cache not found",
			zap.Error(err),
		)
	}
	if cashflowData == "" {
		cashflowData = util.Get(a.GetCashflowUrl(ticker))
	}
	cashflowErr := json.Unmarshal([]byte(cashflowData), &cashflowJson)
	if cashflowErr != nil {
		logger.Log.Error("Failed retrieving cashflow data",
			zap.Error(err),
		)
		return nil, cashflowErr
	}

	if !a.isDataValid(cashflowJson) {
		logger.Log.Error("Invalid cashflow data",
			zap.Any("data", earningJson),
		)
		return nil, errors.New("invalid cashflow data")
	}

	_, err = util.SetCacheData(a.providerName, ticker, "cashflow",
		cashflowData)
	if err != nil {
		logger.Log.Info("Failed setting cache for cashflow data",
			zap.Error(err),
		)
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
	}, nil
}
