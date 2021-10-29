package utils

import "time"

type Coin struct {
	Name        string
	Price       int
	AllTimeHigh int
	AllTimeLow  int
	Volume      int64
	MarketCap   int64

	// Created and Updated if we were to use some sort of SQL structure
	// Gorm would be a good pick for storing this to SQL
	CreatedAt time.Time
	UpdatedAt time.Time
}
