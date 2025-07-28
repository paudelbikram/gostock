package util

import (
	"gostock/backend/logger"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		logger.Log.Error("Error",
			zap.Error(err),
		)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Error("Error",
			zap.Error(err),
		)
	}
	return string(body)
}