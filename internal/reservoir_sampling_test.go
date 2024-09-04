package internal

import (
	"testing"
)

func TestReservoirSampling(t *testing.T) {
	slice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	x := 3
	selected := ReservoirSampling(slice, x)

	if len(selected) != x {
		t.Errorf("esperado %d elementos, mas obteve %d", x, len(selected))
	}

	for _, el := range selected {
		found := false
		for _, orig := range slice {
			if el == orig {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("elemento %s não encontrado no slice original", el)
		}
	}

	x = len(slice)
	selected = ReservoirSampling(slice, x)

	if len(selected) != x {
		t.Errorf("esperado %d elementos, mas obteve %d", x, len(selected))
	}

	for _, el := range selected {
		found := false
		for _, orig := range slice {
			if el == orig {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("elemento %s não encontrado no slice original", el)
		}
	}

	x = 1
	selected = ReservoirSampling(slice, x)

	if len(selected) != x {
		t.Errorf("esperado %d elementos, mas obteve %d", x, len(selected))
	}

	for _, el := range selected {
		found := false
		for _, orig := range slice {
			if el == orig {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("elemento %s não encontrado no slice original", el)
		}
	}

	slice = []string{}
	x = 0
	selected = ReservoirSampling(slice, x)

	if len(selected) != x {
		t.Errorf("esperado %d elementos, mas obteve %d", x, len(selected))
	}
}
