package websocket

// import (
// 	"sectask/domain/samuel"
// 	"sectask/domain/user"
// )

// type GetOrderBookData struct {
// 	Stock              samuel.SamuelStock
// 	UserInformationAll user.UserInformationAllResponse
// 	OrderStatus        []samuel.SamuelOrderStatus
// 	OrderBook          OrderBookBidAsk
// 	OrderBookSamuel    string
// 	RequestOrderType   string
// }

// type GetOrderBookResponse struct {
// 	RequestOrderType string                       `json:"request_order_type"`
// 	UserInformation  user.UserInformationResponse `json:"user_information"`
// 	Stock            user.UserInformationStock    `json:"stock"`
// 	OrderBook        OrderBookBidAsk              `json:"order_book"`
// }

// type OrderBookBidAsk struct {
// 	StockCode   string                `json:"stock_code"`
// 	Bid         []OrderBookBidAskData `json:"bid"`
// 	BidTotalLot int64                 `json:"bid_total_lot"`
// 	Ask         []OrderBookBidAskData `json:"ask"`
// 	AskTotalLot int64                 `json:"ask_total_lot"`
// }

// type OrderBookBidAskData struct {
// 	//IsOrderBook bool                `json:"is_order_book"`
// 	Price int64               `json:"price"`
// 	Lot   int64               `json:"lot"`
// 	User  OrderBookBidAskUser `json:"user"`
// }

// type OrderBookBidAskUser struct {
// 	IsYourPrice bool `json:"is_your_price"`
// 	user.UserOrder
// }
