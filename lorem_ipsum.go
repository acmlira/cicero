package cicero

import (
	"github.com/acmlira/cicero/internal/bridge"
)

const (
	wordsSize  = 4
	phraseSize = 10
)

var internal = bridge.GetInternal()
var words = internal.GetWords()

func LoremIpsum() string {
	A := internal.GenerateNaturalText(words)
	return A
}

func Phrase() string {
	A := internal.ReservoirSampling(words, phraseSize)
	B := internal.GenerateNaturalText(A)
	return B
}

func Words(separator string) string {
	A := internal.ReservoirSampling(words, wordsSize)
	B := internal.JoinStrings(A, separator)
	return B
}

func Word() string {
	return internal.GetRandomElement(words)
}
