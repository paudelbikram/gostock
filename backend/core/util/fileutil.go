package util

import (
	"fmt"
	"log"
	"os"
)

func GetFileContent(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return nil
	}
	return data
}

func GetCacheData(providerName string, ticker string, dataType string) string {
	filePath := fmt.Sprintf("./core/data/%s/cache/%s/%s.json", providerName, ticker, dataType)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(data)
}

func SetCacheData(providerName string, ticker string, dataType string, data string) string {
	filePath := fmt.Sprintf("./core/data/%s/cache/%s/%s.json", providerName, ticker, dataType)
	os.MkdirAll(fmt.Sprintf("./core/data/%s/cache/%s", providerName, ticker), os.ModePerm)
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		log.Println(err)
		return ""
	}
	return data
}

func GetApiKey(providerName string) string {
	keyPath := fmt.Sprintf("./core/data/%s/api.key", providerName)
	return string(GetFileContent(keyPath))
}
