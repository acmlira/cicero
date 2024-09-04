package internal

import (
	"math/rand"
)

func GetRandomElement(slice []string) string {
	if len(slice) == 0 {
		return ""
	}

	index := rand.Intn(len(slice))
	return slice[index]
}
