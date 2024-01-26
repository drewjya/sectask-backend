package technologies

import "github.com/go-resty/resty/v2"

func InitRestyClient() *resty.Client {
	return resty.New()
}
