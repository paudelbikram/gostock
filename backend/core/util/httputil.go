package util

import (
	"io"
	"log"
	"net/http"
)

func Get(url string) string {
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