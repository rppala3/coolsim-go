package random

import (
	"math/rand"
	"time"
)

func Int(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
