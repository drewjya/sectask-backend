package arr

func SumArrInt(s []int) int {
	res := 0
	for _, i2 := range s {
		res += i2
	}
	return res
}
