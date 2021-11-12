// Structs and Variables holding Node paths to elements on the webPage

package scrapers

type ftlbtn struct {
	FilterButton, AddFilter, ApplyFilter, ShowResults, MarketCap, CirculatingSupply, MinCirculatingSupply, MaxCirculatingSupply, MinMarketCap, MaxMarketCap string
}

var FilterButtons = &ftlbtn{
	FilterButton: "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div:nth-child(1) > div.sc-1hxnufv-7.ckiJwf > div.sc-16r8icm-0.sc-36mytl-1.nhbuoh-0.cQmsOq.scroll-initial > div.sc-16r8icm-0.kjciSH.table-control-area > div.sc-1hxnufv-5.fmuRvw > button.x0o17e-0.ewuQaX.sc-36mytl-0.bBafzO.table-control-filter",

	AddFilter:   "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div:nth-child(1) > ul > li:nth-child(5) > button",
	ApplyFilter: "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div:nth-child(1) > div.popup-overlay > div > div > div.sc-16r8icm-0.bLfjir > div > div.sc-1wbb12d-0.jaoXuA > div.sc-16r8icm-0.kjciSH > div > button.x0o17e-0.cEEOTh.cmc-filter-button",
	ShowResults: "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div:nth-child(1) > div.popup-overlay > div > div > div.sc-16r8icm-0.gVroyI > div.sc-16r8icm-0.sc-1o74f6b-0.fOYpRj > button.x0o17e-0.lgnbfA.cmc-filter-button",

	MarketCap:         "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div:nth-child(1) > div.popup-overlay > div > div > div.sc-16r8icm-0.kBEZJK > div.sc-16r8icm-0.ednGyJ.filter-area > div:nth-child(2) > button",
	CirculatingSupply: "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div:nth-child(1) > div.popup-overlay > div > div > div.sc-16r8icm-0.kBEZJK > div.sc-16r8icm-0.ednGyJ.filter-area > div:nth-child(6) > button",

	MinCirculatingSupply: "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div:nth-child(1) > div.popup-overlay > div > div > div.sc-16r8icm-0.bLfjir > div > div.sc-1wbb12d-0.jaoXuA > div.sc-16r8icm-0.bLdyjf > div.cmc-input-row > input:nth-child(1)",
	MaxCirculatingSupply: "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div:nth-child(1) > div.popup-overlay > div > div > div.sc-16r8icm-0.bLfjir > div > div.sc-1wbb12d-0.jaoXuA > div.sc-16r8icm-0.bLdyjf > div.cmc-input-row > input:nth-child(3)",

	MinMarketCap: "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div:nth-child(1) > div.popup-overlay > div > div > div.sc-16r8icm-0.bLfjir > div > div.sc-1wbb12d-0.jaoXuA > div.sc-16r8icm-0.bLdyjf > div.cmc-input-row > input:nth-child(1)",
	MaxMarketCap: "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div:nth-child(1) > div.popup-overlay > div > div > div.sc-16r8icm-0.bLfjir > div > div.sc-1wbb12d-0.jaoXuA > div.sc-16r8icm-0.bLdyjf > div.cmc-input-row > input:nth-child(3)",
}

type clistview struct {
	CustomizeButton, ATHBtn, ATLBtn, H24Btn, D7Btn, Last7days, ApplyButton, CloseBtn string
}

var CustomListView = &clistview{
	"#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div:nth-child(1) > div.sc-1hxnufv-7.ckiJwf > div.sc-16r8icm-0.sc-36mytl-1.nhbuoh-0.cQmsOq.scroll-initial > div.sc-16r8icm-0.kjciSH.table-control-area > div.sc-1hxnufv-5.fmuRvw > button.x0o17e-0.ewuQaX.sc-1v71trr-13.cz6777-0.cEOpQv",
	"/html/body/div[2]/div/div[2]/div/span/div[1]/div[2]/span[3]",
	"/html/body/div[2]/div/div[2]/div/span/div[1]/div[2]/span[4]",
	"/html/body/div[2]/div/div[2]/div/span/div[2]/div[2]/span[2]",
	"/html/body/div[2]/div/div[2]/div/span/div[2]/div[2]/span[3]",
	"/html/body/div[2]/div/div[2]/div/span/div[6]/div[2]/span[2]",
	"body > div.t8xka3-0.clxZon.modalOpened > div > div.t8xka3-3.jKeeFQ > div > button.x0o17e-0.kzspeM",
	"body > div.t8xka3-0.clxZon.modalOpened > div > div.t8xka3-1.XuJWe.has-title > svg",
}

var BaseXpath = "/html/body/div[1]/div/div[1]/div[2]/div/div[1]/div[5]/table/tbody/tr[%v]"

type coinDets struct {
	Name, Price, MarketCap, Volume, Supply, ATH, ATL string
}

var CoinDetailsPaths = &coinDets {
	"/td[3]/div/a/div/div/div/p",
	"/td[4]/div/a",
	"/td[5]/p/span[2]",
	"/td[6]/div/a/p",
	"/td[7]/div/div[1]/p",
	"/td[8]",
	"/td[9]",
}

/*
Row: /html/body/div[1]/div/div[1]/div[2]/div/div[1]/div[5]/table/tbody/tr[i]
Name: /html/body/div[1]/div/div[1]/div[2]/div/div[1]/div[5]/table/tbody/tr[1]/td[3]/div/a/div/div/div/p
Price: /html/body/div[1]/div/div[1]/div[2]/div/div[1]/div[5]/table/tbody/tr[1]/td[4]/div/a
MarketCap: /html/body/div[1]/div/div[1]/div[2]/div/div[1]/div[5]/table/tbody/tr[1]/td[5]/p/span[2]
Volume: /html/body/div[1]/div/div[1]/div[2]/div/div[1]/div[5]/table/tbody/tr[1]/td[6]/div/a/p
Supply: /html/body/div[1]/div/div[1]/div[2]/div/div[1]/div[5]/table/tbody/tr[1]/td[7]/div/div[1]/p
ATH: /html/body/div[1]/div/div[1]/div[2]/div/div[1]/div[5]/table/tbody/tr[1]/td[8]
ATL: /html/body/div[1]/div/div[1]/div[2]/div/div[1]/div[5]/table/tbody/tr[1]/td[9]


Price, MarketCap, Volume needs to be edited to split currency
Supply Split by space


*/