package internal

import (
	"math/rand"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const commaWeight = 0.2
const dotWeight = 0.05

func GenerateNaturalText(words []string) string {
	if len(words) == 0 {
		return ""
	}

	var result strings.Builder

	capitalizeNext := true

	for i, word := range words {
		if capitalizeNext {
			caser := cases.Title(language.English)
			word = caser.String(word)
			capitalizeNext = false
		}

		result.WriteString(word)

		if i != len(words)-1 {
			r := rand.Float32()
			if r < dotWeight {
				result.WriteString(".")
				capitalizeNext = true
			} else if r < commaWeight {
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
