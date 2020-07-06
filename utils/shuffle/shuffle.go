package fisheryates

import (
	"math/rand"
)

func ShuffleInts(input []int) {
	for i := 0; i < len(input); i++ {
		j := rand.Intn(i)
		input[j] = input[i]
	}
}

func Shuffle(input []interface{}) {
}
