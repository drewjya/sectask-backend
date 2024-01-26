package random

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateStringRandomNumber() string {
	rand.Seed(time.Now().Unix())
	return strconv.Itoa(rand.Int())
}
