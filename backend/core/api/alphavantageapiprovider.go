package api

import (
	"encoding/json"
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

func (a *AlphaVantageApiProvider) GetData(ticker string) (map[string]interface{}, error) {
	//getting overview data
	overviewData, err := util.GetCacheData(a.providerName, ticker, "overview")
	if err != nil {
		logger.Log.Info("Cache not found",
			zap.Error(err),
		)
	}
	if overviewData == "" {
		overviewData, err = util.SetCacheData(a.providerName, ticker, "overview",
			util.Get(a.GetOverviewUrl(ticker)))
		if err != nil {
			logger.Log.Info("Data retrieval failed",
				zap.Error(err),
			)
			return nil, err
		}
	}
	//getting incomestatement data
	incomeData, err := util.GetCacheData(a.providerName, ticker, "incomestatement")
	if err != nil {
		logger.Log.Info("Cache not found",
			zap.Error(err),
		)
	}
	if incomeData == "" {
		incomeData, err = util.SetCacheData(a.providerName, ticker, "incomestatement",
			util.Get(a.GetIncomeStatementUrl(ticker)))
		if err != nil {
			logger.Log.Info("Data retrieval failed",
				zap.Error(err),
			)
			return nil, err
		}
	}
	//getting balancesheet data
	balancesheetData, err := util.GetCacheData(a.providerName, ticker, "balancesheet")
	if err != nil {
		logger.Log.Info("Cache not found",
			zap.Error(err),
		)
	}
	if balancesheetData == "" {
		balancesheetData, err = util.SetCacheData(a.providerName, ticker, "balancesheet",
			util.Get(a.GetBalanceSheetUrl(ticker)))
		if err != nil {
			logger.Log.Info("Data retrieval failed",
				zap.Error(err),
			)
			return nil, err
		}
	}
	//getting earning data
	earningData, err := util.GetCacheData(a.providerName, ticker, "earning")
	if err != nil {
		logger.Log.Info("Cache not found",
			zap.Error(err),
		)
	}
	if earningData == "" {
		earningData, err = util.SetCacheData(a.providerName, ticker, "earning",
			util.Get(a.GetEarningUrl(ticker)))
		if err != nil {
			logger.Log.Info("Data retrieval failed",
				zap.Error(err),
			)
			return nil, err
		}
	}
	//getting cashflow data
	cashflowData, err := util.GetCacheData(a.providerName, ticker, "cashflow")
	if err != nil {
		logger.Log.Info("Cache not found",
			zap.Error(err),
		)
	}
	if cashflowData == "" {
		cashflowData, err = util.SetCacheData(a.providerName, ticker, "cashflow",
			util.Get(a.GetCashflowUrl(ticker)))
		if err != nil {
			logger.Log.Info("Data retrieval failed",
				zap.Error(err),
			)
			return nil, err
		}
	}

	var overviewJson map[string]interface{}
	overeviewErr := json.Unmarshal([]byte(overviewData), &overviewJson)
	if overeviewErr != nil {
		logger.Log.Info("Json unmarshal failed",
			zap.Error(err),
		)
		return nil, overeviewErr
	}
	var cashflowJson map[string]interface{}
	cashflowErr := json.Unmarshal([]byte(cashflowData), &cashflowJson)
	if cashflowErr != nil {
		logger.Log.Info("Json unmarshal failed",
			zap.Error(err),
		)
		return nil, cashflowErr
	}
	var balancesheetJson map[string]interface{}
	balancesheetErr := json.Unmarshal([]byte(balancesheetData), &balancesheetJson)
	if balancesheetErr != nil {
		logger.Log.Info("Json unmarshal failed",
			zap.Error(err),
		)
		return nil, balancesheetErr
	}
	var incomeJson map[string]interface{}
	incomeErr := json.Unmarshal([]byte(incomeData), &incomeJson)
	if incomeErr != nil {
		logger.Log.Info("Json unmarshal failed",
			zap.Error(err),
		)
		return nil, incomeErr
	}
	var earningJson map[string]interface{}
	earningErr := json.Unmarshal([]byte(earningData), &earningJson)
	if earningErr != nil {
		logger.Log.Info("Json unmarshal failed",
			zap.Error(err),
		)
		return nil, earningErr
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
