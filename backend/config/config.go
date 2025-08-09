package config

import (
	"encoding/json"
	"go.uber.org/zap"
	"gostock/backend/logger"
	"os"
)

type Config struct {
	Port       int    `json:"port"`
	CORSOrigin string `json:"cors_origin"`
}

func NewConfig() (*Config, error) {
	file, err := os.Open("./config.json")
	if err != nil {
		logger.Log.Error("Failed to open config file",
			zap.Error(err),
		)
		return nil, err
	}
	defer file.Close()
	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		logger.Log.Error("Failed to desearilize file",
			zap.Error(err),
		)
		return nil, err
	}
	return &config, nil
}
