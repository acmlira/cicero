package bridge

import (
	"cicero/internal"
	"fmt"
	"sync"
)

type Internal interface {
	ReservoirSampling([]string, int) []string
	JoinStrings([]string, string) string
	GetWords() []string
	GetRandomElement([]string) string
	GenerateNaturalText([]string) string
}

type internalImpl struct {
	Words []string
}

var internalInstance Internal

var lock = &sync.Mutex{}

func GetInternal() Internal {
	if internalInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if internalInstance == nil {
			fmt.Println("Creating single instance now.")
			internalInstance = &internalImpl{
				Words: internal.Words,
			}
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return internalInstance
}

func (*internalImpl) ReservoirSampling(words []string, count int) []string {
	return internal.ReservoirSampling(words, count)
}

func (*internalImpl) JoinStrings(words []string, separator string) string {
	return internal.JoinStrings(words, separator)
}

func (i *internalImpl) GetWords() []string {
	return i.Words
}

func (*internalImpl) GetRandomElement(words []string) string {
	return internal.GetRandomElement(words)
}

func (*internalImpl) GenerateNaturalText(words []string) string {
	return internal.GenerateNaturalText(words)
}
