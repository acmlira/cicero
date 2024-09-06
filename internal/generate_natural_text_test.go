package internal

import (
	"strings"
	"testing"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TestGenerateNaturalText(t *testing.T) {
	words := []string{"hello", "world", "this", "is", "a", "test"}

	var emptyWords []string
	emptyResult := GenerateNaturalText(emptyWords)
	if emptyResult != "" {
		t.Errorf("Esperado texto vazio, mas obteve %s.", emptyResult)
	}

	result := GenerateNaturalText(words)

	if result == "" {
		t.Errorf("Esperado texto não vazio, mas obteve uma string vazia.")
	}

	if !strings.HasSuffix(result, ".") {
		t.Errorf("Esperado que o texto terminasse com '.', mas obteve: %s", result)
	}

	firstWord := strings.Split(result, " ")[0]
	if !unicode.IsUpper(rune(firstWord[0])) {
		t.Errorf("Esperado que a primeira palavra fosse capitalizada, mas obteve: %s", firstWord)
	}

	sentences := strings.Split(result, ".")
	for _, sentence := range sentences {
		trimmedSentence := strings.TrimSpace(sentence)
		if trimmedSentence != "" {
			firstWordInSentence := strings.Split(trimmedSentence, " ")[0]
			if !unicode.IsUpper(rune(firstWordInSentence[0])) {
				t.Errorf("Esperado que a primeira palavra após um '.' fosse capitalizada, mas obteve: %s", firstWordInSentence)
			}
		}
	}

	for _, word := range words {
		if !strings.Contains(result, word) && !strings.Contains(result, Title(word)) {
			t.Errorf("Esperado que o texto contivesse a palavra: %s", word)
		}
	}

	t.Logf("Texto gerado: %s", result)
}

func Title(word string) string {
	caser := cases.Title(language.English)
	return caser.String(word)
}
