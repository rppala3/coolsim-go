package random

import (
	"math/rand"
	"time"
)

func PullOutInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
