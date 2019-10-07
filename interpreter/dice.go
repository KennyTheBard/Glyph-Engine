package interpreter

import (
	"math/rand"
	"time"
)

func Randomize() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Roll(faces uint) int {
	return rand.Intn(int(faces)) + 1
}
