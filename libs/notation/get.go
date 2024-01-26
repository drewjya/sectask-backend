package notation

import (
	"strings"
)

type NotationStockDefinition struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

func GetNotation(stockNotation string, stockCode string) []NotationStockDefinition {
	notation := []NotationStockDefinition{}
	if stockCode != "" {
		listNotation := strings.Split(stockNotation, ",")
		for _, n := range listNotation {
			notation = append(notation, NotationStockDefinition{
				Code: n,
				Desc: Notation[n],
			})
		}
	}

	return notation
}
