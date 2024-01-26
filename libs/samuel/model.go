package samuel

import (
	"gorm.io/gorm"
)

const (
	Expired                  = "expired"
	PleaseInputPin           = "please input pin"
	PleaseCheckCS            = "please check with our customer service"
	ProcessingBackOfficeData = "processing back office data"
	NoData                   = "no data"
	InvalidToken             = "invalid_token"
	PinRequired              = "pin_required"
	BackOfficeMaintenance    = "back_office_maintenance"
	SamuelError              = "samuel_error"
)

const (
	URLCashPosition             = "/api/v1/cash-position"
	URLCashPositionFundWithdraw = "/api/v1/cash-position-fund-withdraw"
	URLCompanyProfile           = "/api/v1/company-profile"
	URLCorpActions              = "/api/v1/corp-actions"
	URLMiniFundamental          = "/api/v1/mini-fundamental"
	URLOrder                    = "/api/v1/order"
	URLLogin                    = "/api/v1/login"
	URLOrderStatus              = "/api/v1/order-status"
	URLStockPosition            = "/api/v1/stock-position"
)

type SamuelConfig struct {
	DB                *gorm.DB
	SamuelAccessToken string
	BaseURL           string
	BaseURL2          string
}
