package api

type ApiProvider interface {
	GetOverviewUrl(ticker string) string
	GetIncomeStatementUrl(ticker string) string
	GetBalanceSheetUrl(ticker string) string
	GetEarningUrl(ticker string) string
	GetCashflowUrl(ticker string) string
	GetData(ticket string) map[string]interface{}
}
