package internal

import (
	"testing"
)

func TestGetRandomElement(t *testing.T) {
	slice := []string{"apple", "banana", "cherry", "date"}

	elementCount := make(map[string]int)
	for i := 0; i < 1000; i++ {
		element := GetRandomElement(slice)
		if _, exists := elementCount[element]; !exists {
			elementCount[element] = 0
		}
		elementCount[element]++
	}

	for _, fruit := range slice {
		if count, exists := elementCount[fruit]; !exists || count == 0 {
			t.Errorf("Elemento %s não foi retornado.", fruit)
		}
	}

	slice = []string{}
	result := GetRandomElement(slice)
	if result != "" {
		t.Errorf("Esperado uma string vazia, mas obteve '%s'", result)
	}
}
