package util 

import (
	"strconv"
	"log"
	"golang.org/x/text/message"
)


func FormatNumber(numStr string) string {
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.Println(err)
		return ""
	}
	m := message.NewPrinter(message.MatchLanguage("en"))
	return m.Sprintf("%d", num)
}