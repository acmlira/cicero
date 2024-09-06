package bridge

import (
	"sync"
	"testing"
)

func TestGetInternalSingleton(t *testing.T) {
	// Verifica se GetInternal() sempre retorna a mesma instância (padrão Singleton)
	instance1 := GetInternal()
	instance2 := GetInternal()

	if instance1 != instance2 {
		t.Errorf("Esperado que ambas as instâncias sejam iguais, mas são diferentes.")
	}
}

func TestReservoirSampling(t *testing.T) {
	// Testando o método ReservoirSampling
	instance := GetInternal()

	words := []string{"apple", "banana", "cherry", "date"}
	sampleSize := 2
	sample := instance.ReservoirSampling(words, sampleSize)

	if len(sample) != sampleSize {
		t.Errorf("Esperado que o tamanho do sample fosse %d, mas obteve %d", sampleSize, len(sample))
	}
}

func TestJoinStrings(t *testing.T) {
	// Testando o método JoinStrings
	instance := GetInternal()

	words := []string{"hello", "world"}
	separator := " "
	joined := instance.JoinStrings(words, separator)

	expected := "hello world"
	if joined != expected {
		t.Errorf("Esperado: %s, mas obteve: %s", expected, joined)
	}
}

func TestGetWords(t *testing.T) {
	// Testando o método GetWords
	instance := GetInternal()

	words := instance.GetWords()
	if len(words) == 0 {
		t.Errorf("Esperado que Words tivesse elementos, mas estava vazio.")
	}
}

func TestGetRandomElement(t *testing.T) {
	// Testando o método GetRandomElement
	instance := GetInternal()

	words := []string{"apple", "banana", "cherry"}
	randomElement := instance.GetRandomElement(words)

	// Verifica se o elemento retornado está na lista original
	found := false
	for _, word := range words {
		if word == randomElement {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("O elemento retornado (%s) não estava na lista original.", randomElement)
	}
}

func TestGenerateNaturalText(t *testing.T) {
	// Testando o método GenerateNaturalText
	instance := GetInternal()

	words := []string{"hello", "world", "this", "is", "a", "test"}
	text := instance.GenerateNaturalText(words)

	// Verifica se o texto não está vazio
	if text == "" {
		t.Errorf("Esperado texto gerado, mas obteve string vazia.")
	}

	// Verifica se o texto termina com ponto final
	if lastChar := text[len(text)-1]; lastChar != '.' {
		t.Errorf("Esperado que o texto terminasse com '.', mas obteve: %s", string(lastChar))
	}
}

func TestGetInternalConcurrentAccess(t *testing.T) {
	// Verifica se o singleton funciona corretamente em acesso concorrente
	var wg sync.WaitGroup
	instanceChannel := make(chan Internal, 10)

	// Executa várias goroutines tentando acessar a instância
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			instance := GetInternal()
			instanceChannel <- instance
		}()
	}

	wg.Wait()
	close(instanceChannel)

	var instanceReference Internal
	for instance := range instanceChannel {
		if instanceReference == nil {
			instanceReference = instance
		} else if instanceReference != instance {
			t.Errorf("Esperado que todas as instâncias fossem iguais, mas são diferentes.")
		}
	}
}
