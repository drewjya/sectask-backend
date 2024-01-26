package parser

import (
	"sectask/libs/try"
	"strings"
	"time"
)

type SamuelStock struct {
	IndividualIndex        int64     `json:"individual_index"`
	Freq                   int64     `json:"freq"`
	Low                    int64     `json:"low"`
	CorporateAction        string    `json:"corporate_action"`
	BestOfferVolume        int64     `json:"best_offer_volume"`
	HiMTD                  int64     `json:"hi_mtd"`
	AveragePrice           float64   `json:"average_price"`
	Low52W                 int64     `json:"low_52_w"`
	BestBidPrice           int64     `json:"best_bid_price"`
	LowYTD                 int64     `json:"low_ytd"`
	PercentChange          float64   `json:"percent_change"`
	BestOfferPrice         int64     `json:"best_offer_price"`
	TradeTime              string    `json:"trade_time"`
	DateBefore             string    `json:"date_before"`
	LowBefore              string    `json:"low_before"`
	Start52W               time.Time `json:"start_52_w"`
	ForeignBuyerValue      int64     `json:"foreign_buyer_value"`
	Sector                 string    `json:"sector"`
	HiBefore               string    `json:"hi_before"`
	Return52W              float64   `json:"return_52_w"`
	LowMTD                 int64     `json:"low_mtd"`
	MarketCapFreeFloat     int64     `json:"market_cap_free_float"`
	MarketCap              int64     `json:"market_cap"`
	Open                   int64     `json:"open"`
	DomesticBuyerValue     int64     `json:"domestic_buyer_value"`
	Close52W               int64     `json:"close_52_w"`
	ReturnYTD              float64   `json:"return_ytd"`
	DomesticSellerValue    int64     `json:"domestic_seller_value"`
	HiYTD                  int64     `json:"hi_ytd"`
	TradeDate              time.Time `json:"trade_date"`
	End52W                 time.Time `json:"end_52_w"`
	Hi52W                  int64     `json:"hi_52_w"`
	BestBidVolume          int64     `json:"best_bid_volume"`
	AvailableForForeigners int64     `json:"available_for_foreigners"`
	Value                  int64     `json:"value"`
	Close                  int64     `json:"close"`
	Prev                   int64     `json:"prev"`
	ForeignSellerValue     int64     `json:"foreign_seller_value"`
	BoardCode              string    `json:"board_code"`
	Volume                 int64     `json:"volume"`
	ReturnMTD              float64   `json:"return_mtd"`
	Hi                     int64     `json:"hi"`
	StockCode              string    `json:"stock_code"`
	Change                 int64     `json:"change"`
	IEP                    float64   `json:"iep"`
	IEV                    float64   `json:"iev"`
}

func ParseStock(s string) SamuelStock {
	arr := strings.Split(s, "|")
	var res SamuelStock
	res.IndividualIndex = try.ArrayStringToInt(arr, 24, 0)
	res.Freq = try.ArrayStringToInt(arr, 23, 0)
	res.Low = try.ArrayStringToInt(arr, 11, 0)
	res.CorporateAction = try.ArrayStringToString(arr, 31, "--")
	res.BestOfferVolume = try.ArrayStringToInt(arr, 30, 0)
	res.HiMTD = try.ArrayStringToInt(arr, 42, 0)
	res.AveragePrice = try.ArrayStringToFloat64(arr, 33, 0)
	res.Low52W = try.ArrayStringToInt(arr, 14, 0)
	res.BestBidPrice = try.ArrayStringToInt(arr, 12, 0)
	res.LowYTD = try.ArrayStringToInt(arr, 39, 0)
	res.PercentChange = try.ArrayStringToFloat64(arr, 20, 0)
	res.BestOfferPrice = try.ArrayStringToInt(arr, 29, 0)
	res.TradeTime = try.ArrayStringToString(arr, 5, "00:00:00")
	res.DateBefore = try.ArrayStringToString(arr, 0, "tidak ada")
	res.LowBefore = try.ArrayStringToString(arr, 0, "tidak ada")
	res.Start52W = try.ArrayStringToTime(arr, 15, time.Time{})
	res.ForeignBuyerValue = try.ArrayStringToInt(arr, 36, 0)
	res.Sector = try.ArrayStringToString(arr, 8, "")
	res.HiBefore = try.ArrayStringToString(arr, 0, "tidak ada")
	res.Return52W = try.ArrayStringToFloat64(arr, 18, 0)
	res.LowMTD = try.ArrayStringToInt(arr, 43, 0)
	res.MarketCapFreeFloat = try.ArrayStringToInt(arr, 41, 0)
	res.MarketCap = try.ArrayStringToInt(arr, 31, 0)
	res.Open = try.ArrayStringToInt(arr, 26, 0)
	res.DomesticBuyerValue = try.ArrayStringToInt(arr, 34, 0)
	res.Close52W = try.ArrayStringToInt(arr, 17, 0)
	res.ReturnYTD = try.ArrayStringToFloat64(arr, 40, 0)
	res.DomesticSellerValue = try.ArrayStringToInt(arr, 35, 0)
	res.HiYTD = try.ArrayStringToInt(arr, 38, 0)
	res.TradeDate = try.ArrayStringToTime(arr, 4, time.Time{})
	res.End52W = try.ArrayStringToTime(arr, 16, time.Time{})
	res.Hi52W = try.ArrayStringToInt(arr, 13, 0)
	res.BestBidVolume = try.ArrayStringToInt(arr, 28, 0)
	res.AvailableForForeigners = try.ArrayStringToInt(arr, 25, 0)
	res.Value = try.ArrayStringToInt(arr, 22, 0)
	res.Close = try.ArrayStringToInt(arr, 27, 0)
	res.Prev = try.ArrayStringToInt(arr, 9, 0)
	res.ForeignSellerValue = try.ArrayStringToInt(arr, 37, 0)
	res.BoardCode = try.ArrayStringToString(arr, 7, "")
	res.Volume = try.ArrayStringToInt(arr, 21, 0)
	res.ReturnMTD = try.ArrayStringToFloat64(arr, 44, 0)
	res.Hi = try.ArrayStringToInt(arr, 10, 0)
	res.StockCode = try.ArrayStringToString(arr, 6, "")
	res.Change = try.ArrayStringToInt(arr, 19, 0)
	res.IEP = try.ArrayStringToFloat64(arr, 45, 0)
	res.IEV = try.ArrayStringToFloat64(arr, 46, 0)
	return res
}
