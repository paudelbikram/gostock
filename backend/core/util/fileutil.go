package util

import (
	"fmt"
	"gostock/backend/logger"
	"os"

	"go.uber.org/zap"
)

func GetFileContent(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		logger.Log.Error("Error",
			zap.Error(err),
		)
		return nil
	}
	return data
}

func GetCacheData(providerName string, ticker string, dataType string) (string, error) {
	filePath := fmt.Sprintf("./core/data/%s/cache/%s/%s.json", providerName, ticker, dataType)
	data, err := os.ReadFile(filePath)
	if err != nil {
		logger.Log.Error("Error",
			zap.Error(err),
		)
		return "", err
	}
	return string(data), nil
}

func SetCacheData(providerName string, ticker string, dataType string, data string) (string, error) {
	filePath := fmt.Sprintf("./core/data/%s/cache/%s/%s.json", providerName, ticker, dataType)
	err := os.MkdirAll(fmt.Sprintf("./core/data/%s/cache/%s", providerName, ticker), os.ModePerm)
	if err != nil {
		logger.Log.Error("Error",
			zap.Error(err),
		)
		return "", err
	}
	err = os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		logger.Log.Error("Error",
			zap.Error(err),
		)
		return "", err
	}
	return data, nil
}

func GetApiKey(providerName string) string {
	keyPath := fmt.Sprintf("./core/data/%s/api.key", providerName)
	return string(GetFileContent(keyPath))
}
