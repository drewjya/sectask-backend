package try

import (
	"fmt"
	"strconv"
	"time"
)

func ArrayStringToString(array []string, index int, defaultValue string) string {
	var result string
	This(func() {
		result = fmt.Sprintf("%v", array[index])
	}).Catch(func(e E) {
		result = defaultValue
	})
	return result
}

func ArrayStringToInt(array []string, index int64, defaultValue int64) int64 {
	var result int64
	This(func() {
		result, _ = strconv.ParseInt(fmt.Sprintf("%v", array[index]), 10, 64)
	}).Catch(func(e E) {
		result = defaultValue
	})
	return result
}

func ArrayStringToFloat64(array []string, index int, defaultValue float64) float64 {
	var result float64
	This(func() {
		result, _ = strconv.ParseFloat(fmt.Sprintf("%v", array[index]), 64)
	}).Catch(func(e E) {
		result = defaultValue
	})
	return result
}

func ArrayStringToTime(array []string, index int, defaultValue time.Time) time.Time {
	var result time.Time
	This(func() {
		result, _ = time.Parse("2006-01-02", fmt.Sprintf("%v", array[index]))
	}).Catch(func(e E) {
		result = defaultValue
	})
	return result
}
