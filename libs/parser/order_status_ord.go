package parser

import (
	"sectask/libs/try"
	"strings"
)

func ParseOrderStatusORD(s string) SamuelOrderStatusB {
	//"bek|JASO005|ZZZZFITRIAH M|BUMI|BUY|OPEN|A|O|-|RG|RG|D|H|110|202212130000001525|-|-|0|2022-12-13|09:56:06|09:56:39||||0|100|300|30000|100|0|0|0|209|209|RT|JASO005|1|LD"

	//userid = data[1]; bek
	//clientcode = data[2]; JASO005
	//clientname = data[3]; ZZZZFITRIAH M
	//stockcode = data[4]; BUMI
	//order_type = data[5]; BUY
	//order_status = data[6]; OPEN
	//order_status_code = data[7]; A
	//order_status_code_new = data[8]; O
	//message = data[9]; -
	//market = data[10]; RG
	//ads = data[11]; RG
	//fd = data[12]; D (F/D)
	//ds = data[13]; H (Data Source)
	//order_no_olt = data[14]; 110
	//order_no_remote = data[14]; 110
	//order_no_jsx = data[15]; 202212130000001525
	//order_date = data[16]; 2022-12-13
	//order_time = data[17]; 09:56:39
	//trade_no = data[18];
	//trade_date = data[19];
	//trade_time = data[20];
	//avg_cost = data[10]; 0
	//qty = data[22]); 100
	//price = data[23]); 300
	//total_value = data[24]; 30000
	//balance = data[25]; 100
	//done_qty = data[26];0
	//done_price = data[27];0
	//done_total = data[28];0
	//ord_start =data[29]; 209 (order start)
	//ord_end = data[30]; 209 (order end)
	//branch = data[31];RT (branch)
	//clientinit = data[32]; JASO005
	//child_total = data[33]; 1
	//price_type = data[34]; LD

	arr := strings.Split(s, "|")
	var res SamuelOrderStatusB
	//res.Userid = try.ArrayStringToString(arr, 0, "")
	res.Clientcode = try.ArrayStringToString(arr, 1, "")
	//res.Clientname = try.ArrayStringToString(arr, 2, "")
	res.Stockcode = try.ArrayStringToString(arr, 3, "")
	res.OrderType = try.ArrayStringToString(arr, 4, "")
	res.OrderStatus = try.ArrayStringToString(arr, 5, "")
	//res.OrderStatusCode = try.ArrayStringToString(arr, 6, "")
	res.Message = try.ArrayStringToString(arr, 8, "")
	//res.Market = try.ArrayStringToString(arr, 9, "")
	//res.AvgCost = try.ArrayStringToString(arr, 9, "")
	//res.Ads = try.ArrayStringToString(arr, 10, "")
	//res.Fd = try.ArrayStringToString(arr, 11, "")
	//res.Ds = try.ArrayStringToString(arr, 12, "")
	res.OrderNoOlt = try.ArrayStringToString(arr, 13, "")
	res.OrderNoRemote = try.ArrayStringToString(arr, 13, "")
	res.OrderNoJsx = try.ArrayStringToString(arr, 14, "")
	//res.GtcNo = try.ArrayStringToString(arr, 24, ")
	//res.AutoNo = try.ArrayStringToString(arr, 24, ")
	//res.OrderDate = try.ArrayStringToString(arr, 15, "")
	//res.OrderTime = try.ArrayStringToString(arr, 16, "")
	//res.MainType = try.ArrayStringToString(arr, 24, ")
	//res.MainTypeText = try.ArrayStringToString(arr, 24, ")
	//res.TradeNo = try.ArrayStringToString(arr, 17, "")
	//res.TradeDate = try.ArrayStringToString(arr, 18, "")
	//res.TradeTime = try.ArrayStringToString(arr, 19, "")
	//res.TotalValue = try.ArrayStringToString(arr, 23, "")
	//res.Balance = try.ArrayStringToString(arr, 24, "")
	res.Qty = try.ArrayStringToString(arr, 25, "")
	res.Price = try.ArrayStringToString(arr, 26, "")
	res.DoneQty = try.ArrayStringToString(arr, 29, "")
	res.DonePrice = try.ArrayStringToString(arr, 30, "")
	res.DoneTotal = try.ArrayStringToString(arr, 31, "")
	//res.Branch = try.ArrayStringToString(arr, 30, "")
	//res.Clientinit = try.ArrayStringToString(arr, 31, "")
	//res.PriceType = try.ArrayStringToString(arr, 33, "")
	//res.ChildTotal = try.ArrayStringToString(arr, 32, "")
	return res
}
