package rand

import (
	"math/rand"
	"time"
)

func init() {
	// set random number seed
	rand.Seed(time.Now().UnixNano())
}
