package util

import (
	"gostock/backend/logger"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Get(url string) string {
	// sleeping for 1 seconds because API won't allow more than 1 request per second
	time.Sleep(1 * time.Second)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		logger.Log.Error("Error running get",
			zap.Error(err),
			zap.Any("Response", resp),
		)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Error("Error reading response",
			zap.Error(err),
			zap.Any("body", body),
		)
	}
	return string(body)
}
