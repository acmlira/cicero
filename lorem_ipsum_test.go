package cicero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoremIpsum(t *testing.T) {
	result := LoremIpsum()

	assert.NotEqual(t, "", result)
}

func BenchmarkLoremIpsum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoremIpsum()
	}
}

func TestPhrase(t *testing.T) {
	result := Phrase()

	assert.NotEqual(t, "", result)
}

func BenchmarkPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phrase()
	}
}

func TestWord(t *testing.T) {
	result := Word()

	assert.NotEqual(t, "", result)
	isPresent := false
	for _, word := range words {
		if word == result {
			isPresent = true
		}
	}
	assert.True(t, isPresent)
}

func BenchmarkWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Word()
	}
}

func TestWords(t *testing.T) {
	result := Words(" ")

	assert.NotEqual(t, "", result)
}

func BenchmarkWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Words(" ")
	}
}
