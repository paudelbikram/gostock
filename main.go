package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	settingApiKeyCommand := "APIKEY"
	runningStockAnalysisCommand := "TICKER"
	if len(os.Args) < 3 {
		log.Println("Invalid command provided.")
		commandHelp()
		return
	}
	commandProvided := strings.ToUpper(os.Args[1])
	if commandProvided == settingApiKeyCommand {
		apiKey := os.Args[2]
		setApiKey(apiKey)
		return
	} else if commandProvided == runningStockAnalysisCommand {
		if getApiKey() == "" {
			log.Println("Api Key not found.")
			apiKeyHelp()
			return
		}
		tickerSymbol := strings.ToUpper(os.Args[2])
		get(tickerSymbol)
		createView(tickerSymbol)
		return
	}
}

func commandHelp() {
	log.Println("Following command are supported.")
	log.Println("gostock apikey yourapikey")
	log.Println("gostock ticker tickersymbol")
}

func apiKeyHelp() {
	log.Println("Please visit https://www.alphavantage.co/support/#api-key to get free api key.")
	log.Println("Once you have api key, please run following command to set it.")
	log.Println("gostock apikey yourapikey")
}
