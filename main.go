package main

import (
	"CmcScraper/scrapers"
	u "CmcScraper/utils"
	"context"
	"github.com/chromedp/chromedp"
	"strconv"
	"strings"
)

func ScrapeData(numIters int, ctx context.Context) {
	for i := 0; i <= numIters; i++ {
		// Loop through each coin on the page
		for j := 0; j <= 100; j++ {
			var rawData u.RawCoinStruct
			var coinData u.ParsedCoin

			_ = chromedp.Run(ctx, scrapers.ScrapeCoin(j, &rawData))
		}
		scrapers.LoadNExtPage()
	}
}
func PreliminaryTasks(ctx context.Context) (int, error) {
	url := "https://coinmarketcap.com/"
	// Load Config from file
	params := u.ReadConfigFromJSON("config.json")

	//Navigate to URL
	err := chromedp.Run(ctx, chromedp.Navigate(url))
	u.CheckErr(err)

	// Filter Coins
	_ = chromedp.Run(ctx, scrapers.FilterCoins(params))

	//Customize list view
	_ = chromedp.Run(ctx, scrapers.CustomizeListView())

	//Find & return Num Coins returned from filtering
	var numCoins string
	_ = chromedp.Run(ctx, scrapers.FindNumIterations(&numCoins))

	// DO i split by spaces and pick last element in array?
	// Or split at 'of'
	chars := strings.Split(numCoins, "of ") // Intentional Traiing space

	return strconv.Atoi(chars[1])
}
