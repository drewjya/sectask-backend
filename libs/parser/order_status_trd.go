package parser

import (
	"sectask/libs/try"
	"strings"
)

type SamuelOrderStatusB struct {
	Userid          string `json:"userid"`
	Clientcode      string `json:"clientcode"`
	Clientinit      string `json:"clientinit"`
	Clientname      string `json:"clientname"`
	Stockcode       string `json:"stockcode"`
	OrderType       string `json:"order_type"`
	PriceType       string `json:"price_type"`
	OrderStatus     string `json:"order_status"`
	OrderStatusCode string `json:"order_status_code"`
	Message         string `json:"message"`
	Market          string `json:"market"`
	Ads             string `json:"ads"`
	Fd              string `json:"fd"`
	Ds              string `json:"ds"`
	OrderNoOlt      string `json:"order_no_olt"`
	OrderNoRemote   string `json:"order_no_remote"`
	OrderNoJsx      string `json:"order_no_jsx"`
	GtcNo           string `json:"gtc_no"`
	AutoNo          string `json:"auto_no"`
	OrderDate       string `json:"order_date"`
	OrderTime       string `json:"order_time"`
	MainType        string `json:"main_type"`
	MainTypeText    string `json:"main_type_text"`
	TradeNo         string `json:"trade_no"`
	TradeDate       string `json:"trade_date"`
	TradeTime       string `json:"trade_time"`
	AvgCost         string `json:"avg_cost"`
	Qty             string `json:"qty"`
	Price           string `json:"price"`
	TotalValue      string `json:"total_value"`
	Balance         string `json:"balance"`
	DoneQty         string `json:"done_qty"`
	DonePrice       string `json:"done_price"`
	DoneTotal       string `json:"done_total"`
	Branch          string `json:"branch"`
	ChildTotal      string `json:"child_total"`
}

func ParseOrderStatusTRD(s string) SamuelOrderStatusB {
	//"dsk|BLAC001|BLACKTEST|JSMR|BUY|RG|D|-|D|202|202212130000001056|2022-12-13|09:39:18|202212130000000273|2022-12-13|09:39:50|50|3280|5000|3280|3280|SR|1|12345678"

	//userid = data[1]; dsk
	//clientcode = data[2]; BLAC001
	//clientname = data[3]; BLACKTEST
	//stockcode = data[4]; JSMR
	//order_type = data[5]; BUY
	//market = data[6]; RG
	//fd = data[7]; D (F/D)
	//contra_broker = data[8]; -
	//order_from = data[8]; D
	//order_no_remote = data[12]; 202
	//order_no_jsx = data[12]; 202212130000001056
	//order_date = data[14]; 2022-12-13
	//order_time = data[14]; 09:39:18
	//trade_no = data[17]; 202212130000000273
	//trade_date = data[18]; 2022-12-13
	//trade_time = data[19]; 09:39:50
	//qty = data[19]; 50
	//price = data[19]; 3280
	//order_qty = data[19]; 5000
	//order_price = data[19]; 3280
	//avgcost = data[19]; 3280
	//branch = data[19]; SR
	//seq = data[19]; 1
	//client_init= data[19];12345678

	arr := strings.Split(s, "|")
	var res SamuelOrderStatusB
	res.Userid = try.ArrayStringToString(arr, 0, "")
	res.Clientcode = try.ArrayStringToString(arr, 1, "")
	res.Clientname = try.ArrayStringToString(arr, 2, "")
	res.Stockcode = try.ArrayStringToString(arr, 3, "")
	res.OrderType = try.ArrayStringToString(arr, 4, "")
	//res.OrderStatus = try.ArrayStringToString(arr, 5, "")
	//res.OrderStatusCode = try.ArrayStringToString(arr, 6, "")
	//res.Message = try.ArrayStringToString(arr, 8, "")
	res.Market = try.ArrayStringToString(arr, 5, "")
	//res.Ads = try.ArrayStringToString(arr, 10, "")
	res.Fd = try.ArrayStringToString(arr, 6, "")
	//res.Ds = try.ArrayStringToString(arr, 12, "")
	//res.OrderNoOlt = try.ArrayStringToString(arr, 13, "")
	res.OrderNoRemote = try.ArrayStringToString(arr, 11, "")
	res.OrderNoJsx = try.ArrayStringToString(arr, 11, "")
	//res.GtcNo = try.ArrayStringToString(arr, 24, ")
	//res.AutoNo = try.ArrayStringToString(arr, 24, ")
	res.OrderDate = try.ArrayStringToString(arr, 13, "")
	res.OrderTime = try.ArrayStringToString(arr, 13, "")
	//res.MainType = try.ArrayStringToString(arr, 24, ")
	//res.MainTypeText = try.ArrayStringToString(arr, 24, ")
	res.TradeNo = try.ArrayStringToString(arr, 16, "")
	res.TradeDate = try.ArrayStringToString(arr, 17, "")
	res.TradeTime = try.ArrayStringToString(arr, 18, "")
	res.Qty = try.ArrayStringToString(arr, 18, "")
	res.Price = try.ArrayStringToString(arr, 18, "")
	res.AvgCost = try.ArrayStringToString(arr, 18, "")
	//res.TotalValue = try.ArrayStringToString(arr, 23, "")
	//res.Balance = try.ArrayStringToString(arr, 24, "")
	//res.DoneQty = try.ArrayStringToString(arr, 25, "")
	//res.DonePrice = try.ArrayStringToString(arr, 26, "")
	//res.DoneTotal = try.ArrayStringToString(arr, 27, "")
	res.Branch = try.ArrayStringToString(arr, 18, "")
	res.Clientinit = try.ArrayStringToString(arr, 18, "")
	//res.PriceType = try.ArrayStringToString(arr, 33, "")
	//res.ChildTotal = try.ArrayStringToString(arr, 32, "")
	return res
}
