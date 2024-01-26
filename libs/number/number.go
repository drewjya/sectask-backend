package number

import (
	"fmt"
	"strconv"
)

func DefaultZero(data string) string {
	if data == "" {
		return "0"
	}
	return data
}

func FormattingNumber(data string, currency string) string {

	f, _ := strconv.ParseFloat(data, 64)
	// f = f * MNum
	switch {
	case f >= TNum || f <= -TNum:
		data = currency + fmt.Sprintf("%.2f %v", f/TNum, TKey)
	case f >= BNum || f <= -BNum:
		data = currency + fmt.Sprintf("%.2f %v", f/BNum, BKey)
	case f >= MNum || f <= -MNum:
		data = currency + fmt.Sprintf("%.2f %v", f/MNum, MKey)
	case f >= KNum || f <= -KNum:
		data = currency + fmt.Sprintf("%.2f %v", f/KNum, KKey)
	case f != 0:
		data = currency + fmt.Sprintf("%.2f", f)
	}

	return data
}

// func RoundFloat(val string, precision uint) string {
// 	fval, _ := strconv.ParseFloat(val, 64)
// 	ratio := math.Pow(10, float64(precision))
// 	r := math.Round(fval*ratio) / ratio
// 	s := fmt.Sprintf("%.2f", 123.456)
// 	return
// }
