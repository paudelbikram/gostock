package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type KeyUrl struct {
	data map[string]string
}

func getKeyUrlConfig() KeyUrl {
	return KeyUrl{
		data: map[string]string{
			"overview":        "https://www.alphavantage.co/query?function=OVERVIEW&symbol=%s&apikey=%s",
			"incomestatement": "https://www.alphavantage.co/query?function=INCOME_STATEMENT&symbol=%s&apikey=%s",
			"balancesheet":    "https://www.alphavantage.co/query?function=BALANCE_SHEET&symbol=%s&apikey=%s",
			"earning":         "https://www.alphavantage.co/query?function=EARNINGS&symbol=%s&apikey=%s",
			"cashflow":        "https://www.alphavantage.co/query?function=CASH_FLOW&symbol=%s&apikey=%s",
		},
	}
}

func httpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return string(body)
}

func getFileContent(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return nil
	}
	return data
}

func getApiKey() string {
	return string(getFileContent("./data/api.key"))
}

func setApiKey(key string) {
	os.MkdirAll("./data", os.ModePerm)
	err := os.WriteFile("./data/api.key", []byte(key), 0644)
	if err != nil {
		log.Println(err)
	}
}

func getCacheData(ticker string, dataType string) string {
	filePath := fmt.Sprintf("./data/cache/%s/%s.json", ticker, dataType)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(data)
}

func setCacheData(ticker string, dataType string, data string) string {
	filePath := fmt.Sprintf("./data/cache/%s/%s.json", ticker, dataType)
	os.MkdirAll(fmt.Sprintf("./data/cache/%s/", ticker), os.ModePerm)
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		log.Println(err)
		return ""
	}
	return data
}

func getStockData(ticker string, apiKey string, dataType string) string {
	cacheData := getCacheData(ticker, dataType)

	if cacheData != "" {
		return cacheData
	}
	return setCacheData(ticker, dataType,
		httpGet(fmt.Sprintf(getKeyUrlConfig().data[dataType], ticker, apiKey)))
}

func get(ticker string) {
	apiKey := getApiKey()
	keyUrlConfig := getKeyUrlConfig()
	for key := range keyUrlConfig.data {
		getStockData(ticker, apiKey, key)
	}
}
