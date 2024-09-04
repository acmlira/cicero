package internal

import (
	"math/rand"
	"strings"
)

func GenerateNaturalText(words []string) string {
	if len(words) == 0 {
		return ""
	}

	var result strings.Builder

	capitalizeNext := true

	for i, word := range words {
		if capitalizeNext {
			// TODO refactor this next line
			word = strings.Title(word)
			capitalizeNext = false
		}

		result.WriteString(word)

		if i != len(words)-1 {
			r := rand.Float32()
			if r < 0.05 {
				result.WriteString(".")
				capitalizeNext = true
			} else if r < 0.2 {
				result.WriteString(",")
			}

			result.WriteString(" ")
		}
	}

	if lastChar := result.String()[result.Len()-1]; lastChar != '.' {
		result.WriteString(".")
	}

	return result.String()
}
