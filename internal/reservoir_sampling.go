package internal

import (
	"math/rand"
)

func ReservoirSampling(slice []string, x int) []string {
	selected := make([]string, x)
	copy(selected, slice[:x])

	for i := x; i < len(slice); i++ {
		j := rand.Intn(i + 1)
		if j < x {
			selected[j] = slice[i]
		}
	}

	return selected
}
