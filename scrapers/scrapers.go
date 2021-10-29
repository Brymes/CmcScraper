package scrapers

import (
	"github.com/chromedp/chromedp"
	"log"
)

// Mini Functions that return tasks to perform typically one thing
// All Functions here can be merged into one single task, But for error management we separate tasks into different functions
func ScrapeCoinName(name *string) chromedp.Tasks {
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
