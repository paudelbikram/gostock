package util

import (
	"errors"
	"fmt"
	"gostock/backend/logger"
	"os"
	"path/filepath"
	"time"

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
	fileInfo, statErr := os.Stat(filePath)
	if statErr != nil {
		logger.Log.Error("Error",
			zap.Error(statErr),
		)
		return "", statErr
	}
	// if file is older than a month, let's recreate a new one 
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	if fileInfo.ModTime().Before(oneMonthAgo) {
		os.Remove(filePath)
		return "", errors.New("old data found; recreating cache")
	}
	
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

func GetCacheStock() ([]string, error) {
	matches, err := filepath.Glob("./core/data/*/cache/*")
	if err != nil {
		logger.Log.Error("Error reading folder with regex",
			zap.Error(err),
		)
		return nil, err
	}
	var folderNames []string
	for _, path := range matches {
		info, err := os.Stat(path)
		if err != nil {
			logger.Log.Error("Error reading folder match",
				zap.Error(err),
			)
			continue // skip invalid paths
		}
		if info.IsDir() {
			folderNames = append(folderNames, filepath.Base(path))
		}
	}
	return folderNames, nil
}
