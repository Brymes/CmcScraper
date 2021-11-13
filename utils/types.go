package utils

import (
	"gorm.io/gorm"
)

type ParsedCoin struct {
	gorm.Model

	ID          uint `gorm:"autoIncrement"`
	Name        string
	Price       int
	AllTimeHigh int
	AllTimeLow  int
	MarketCap   int
	Supply      int

	// Volume is 24 hours volume in USD, It could be changed to 7days or 30 days
	//and can also be measured in the Currency Not only in USD
	Volume      int

	// Created and Updated if we were to use some sort of SQL structure
	// Gorm would be a good pick for storing this to SQL
	/*CreatedAt time.Time
	UpdatedAt time.Time*/
}

type RawCoinStruct struct {
	Name        string
	Price       string
	AllTimeHigh string
	AllTimeLow  string

	MarketCap string
	Supply    string

	// Volume is 24 hours volume in USD, It could be changed to 7days or 30 days
	//and can also be measured in the Currency Not only in USD
	Volume string
}

type Config struct {
	MinMarketCap int `json:"minMarketCap"`
	MaxMarketCap int `json:"maxMarketCap"`
	MinSupply    int `json:"minSupply"`
	MaxSupply    int `json:"maxSupply"`
}
