package scrapers

import (
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

// Mini Functions that return tasks to perform typically one thing
// All Functions here can be merged into one single task, But for error management we separate tasks into different functions
/*func ScrapeCoinName(name *string) chromedp.Tasks {
	log.Println("Scraping for Coin Name")
	return chromedp.Tasks{
		chromedp.Text("#__next > div > div > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div.sc-16r8icm-0.jKrmxw.container > div > div.sc-16r8icm-0.sc-19zk94m-1.gRSJaB > div.sc-16r8icm-0.iutcov > div.sc-16r8icm-0.sc-19zk94m-5.gshXSv > div > div > div.sc-16r8icm-0.sc-1etv19d-1.fyZQQz > div.sc-16r8icm-0.kjciSH > p.q7nmo0-0.smLap.converter-item-name", name, chromedp.ByQuery),
	}
}
func ScrapeCoinPrice(price *string) chromedp.Tasks {
	log.Println("Scraping for Coin Price")
	return chromedp.Tasks{
		chromedp.Text("#__next > div > div > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div.sc-16r8icm-0.jKrmxw.container > div > div.sc-16r8icm-0.sc-19zk94m-1.gRSJaB > div.sc-16r8icm-0.iutcov > div.sc-16r8icm-0.hgKnTV > div > div:nth-child(2) > table > tbody > tr:nth-child(1) > td", price, chromedp.ByQuery),
	}
}

func ScrapeVolume(volume *string) chromedp.Tasks {
	log.Println("Scraping for Volume")
	return chromedp.Tasks{
		chromedp.Text("#__next > div > div > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div.sc-16r8icm-0.jKrmxw.container > div > div.sc-16r8icm-0.sc-19zk94m-1.gRSJaB > div.sc-16r8icm-0.iutcov > div.sc-16r8icm-0.hgKnTV > div > div:nth-child(2) > table > tbody > tr:nth-child(4) > td > span", volume, chromedp.ByQuery),
	}
}

func ScrapeMarketCap(marketCap *string) chromedp.Tasks {
	log.Println("Scraping for MarketCap")
	return chromedp.Tasks{
		chromedp.Text("#__next > div > div > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div.sc-16r8icm-0.jKrmxw.container > div > div.sc-16r8icm-0.sc-19zk94m-1.gRSJaB > div.sc-16r8icm-0.iutcov > div.sc-16r8icm-0.hgKnTV > div > div:nth-child(3) > table > tbody > tr:nth-child(1) > td > span", marketCap, chromedp.ByQuery),
	}
}

func ScrapeCoinATH(ATH *string) chromedp.Tasks {
	log.Println("Scraping for Coin All Time High")
	//TODO :: NOTE We Clicking the Show more button here assuming we scrape sequentially
	return chromedp.Tasks{
		chromedp.Click("#__next > div > div > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div.sc-16r8icm-0.jKrmxw.container > div > div.sc-16r8icm-0.sc-19zk94m-1.gRSJaB > div.sc-16r8icm-0.iutcov > div.sc-16r8icm-0.hgKnTV > div > button", chromedp.ByQuery),
		chromedp.Text("#__next > div > div > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div.sc-16r8icm-0.jKrmxw.container > div > div.sc-16r8icm-0.sc-19zk94m-1.gRSJaB > div.sc-16r8icm-0.iutcov > div.sc-16r8icm-0.hgKnTV > div > div.sc-16r8icm-0.kjciSH.show > div:nth-child(2) > table > tbody > tr:nth-child(5) > td > span", ATH, chromedp.ByQuery),
	}
}

func ScrapeATL(ATL *string) chromedp.Tasks {
	log.Println("Scraping for Coin All Time Low")
	return chromedp.Tasks{
		chromedp.Text("#__next > div > div > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div.sc-16r8icm-0.jKrmxw.container > div > div.sc-16r8icm-0.sc-19zk94m-1.gRSJaB > div.sc-16r8icm-0.iutcov > div.sc-16r8icm-0.hgKnTV > div > div.sc-16r8icm-0.kjciSH.show > div:nth-child(2) > table > tbody > tr:nth-child(6) > td > span", ATL, chromedp.ByQuery),
	}
}
*/
func FilterCoins(minMarketCap, maxMarketCap, minSupply, maxSupply int) chromedp.Tasks {
	log.Println("Entering Filtering ")
	return chromedp.Tasks{
		chromedp.Click(FilterButtons.FilterButton, chromedp.ByQuery),
		chromedp.Click(FilterButtons.AddFilter, chromedp.ByQuery),

		chromedp.Click(FilterButtons.MarketCap, chromedp.ByQuery),
		// chromedp.Click(FilterButtons.MinMarketCap, chromedp.ByQuery),
		chromedp.SendKeys(FilterButtons.MinMarketCap, fmt.Sprintf("%v", minMarketCap), chromedp.ByQuery),
		// chromedp.Click(FilterButtons.MaxMarketCap, chromedp.ByQuery),
		chromedp.SendKeys(FilterButtons.MaxMarketCap, fmt.Sprintf("%v", maxMarketCap), chromedp.ByQuery),

		chromedp.Click(FilterButtons.ApplyFilter, chromedp.ByQuery),

		chromedp.Click(FilterButtons.CirculatingSupply, chromedp.ByQuery),
		// chromedp.Click(FilterButtons.MinCirculatingSupply, chromedp.ByQuery),
		chromedp.SendKeys(FilterButtons.MinCirculatingSupply, fmt.Sprintf("%v", minSupply), chromedp.ByQuery),
		// chromedp.Click(FilterButtons.MaxCirculatingSupply, chromedp.ByQuery),
		chromedp.SendKeys(FilterButtons.MaxCirculatingSupply, fmt.Sprintf("%v", maxSupply), chromedp.ByQuery),

		chromedp.Click(FilterButtons.ApplyFilter, chromedp.ByQuery),
		chromedp.Click(FilterButtons.ShowResults, chromedp.ByQuery),
	}
}

func CustomizeListView() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Click(CustomListView.CustomizeButton, chromedp.ByQuery),

		chromedp.Sleep(3 * time.Second),

		chromedp.Click(CustomListView.D7Btn, chromedp.BySearch),
		chromedp.Click(CustomListView.H24Btn, chromedp.BySearch),
		chromedp.Click(CustomListView.Last7days, chromedp.BySearch),

		chromedp.Sleep(time.Second),

		chromedp.Click(CustomListView.ATHBtn, chromedp.BySearch),
		chromedp.Click(CustomListView.ATLBtn, chromedp.BySearch),

		chromedp.Sleep(time.Second),

		chromedp.Click(CustomListView.ApplyButton, chromedp.ByQuery),
		chromedp.Click(CustomListView.CloseBtn, chromedp.ByQuery),
	}
}