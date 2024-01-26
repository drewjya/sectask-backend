package samuel

type SamuelResponse struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
}

type SamuelStockPositionData struct {
	ERROR struct {
		Msg string `json:"Msg"`
	} `json:"ERROR"`
	A struct {
		Datadate      string                 `json:"datadate"`
		Datatime      string                 `json:"datatime"`
		Clientcode    string                 `json:"clientcode"`
		Dateprev      string                 `json:"dateprev"`
		TotalCost     string                 `json:"total_cost"`
		TotalMarket   string                 `json:"total_market"`
		Ungainloss    string                 `json:"ungainloss"`
		Ungainlosspct string                 `json:"ungainlosspct"`
		B             interface{}            `json:"b"`
		ListB         []SamuelStockPositionB `json:"list_b"`
	} `json:"a"`
}

type SamuelStockPositionB struct {
	Stockcode     string `json:"stockcode"`
	Availtrade    string `json:"availtrade"`
	Balance       string `json:"balance"`
	DoneBuy       string `json:"done_buy"`
	DoneSell      string `json:"done_sell"`
	OpenBuy       string `json:"open_buy"`
	OpenSell      string `json:"open_sell"`
	EndBalance    string `json:"end_balance"`
	AvgCost       string `json:"avg_cost"`
	TotalCost     string `json:"total_cost"`
	MarketPrice   string `json:"market_price"`
	MarketValue   string `json:"market_value"`
	Ungainloss    string `json:"ungainloss"`
	Ungainlosspct string `json:"ungainlosspct"`
	Cap           string `json:"cap"`
	CapValue      string `json:"cap_value"`
	PctTotal      string `json:"pct_total"`
}

type ResponseRealizeGainLoss struct {
	ClientCode  string  `json:"clientCode"`
	TotalAmount float64 `json:"totalAmount"`
	Pct         float64 `json:"pct"`
	Detail      []struct {
		Stock  string  `json:"stock"`
		Amount float64 `json:"amount"`
		Pct    float64 `json:"pct"`
	} `json:"detail"`
}
