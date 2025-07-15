package core

import(
	"gostock/backend/core/api"
)

type DataCollector struct {
	apiProvider api.ApiProvider
}

func NewDataCollector(apiProvider api.ApiProvider) *DataCollector {
	return &DataCollector{
		apiProvider: apiProvider,
	}
}

func (d *DataCollector) RequestData(ticker string) map[string]interface{} {
	return d.apiProvider.GetData(ticker)
}
