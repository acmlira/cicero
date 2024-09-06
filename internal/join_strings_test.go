package internal

import (
	"testing"
)

func TestJoinStrings(t *testing.T) {
	slice := []string{"Hello", "World"}
	separator := " "
	expected := "Hello World"
	result := JoinStrings(slice, separator)
	if result != expected {
		t.Errorf("esperado '%s', mas obteve '%s'", expected, result)
	}

	separator = "-"
	expected = "Hello-World"
	result = JoinStrings(slice, separator)
	if result != expected {
		t.Errorf("esperado '%s', mas obteve '%s'", expected, result)
	}

	separator = ","
	expected = "Hello,World"
	result = JoinStrings(slice, separator)
	if result != expected {
		t.Errorf("esperado '%s', mas obteve '%s'", expected, result)
	}

	slice = []string{}
	separator = " "
	expected = ""
	result = JoinStrings(slice, separator)
	if result != expected {
		t.Errorf("esperado '%s', mas obteve '%s'", expected, result)
	}

	slice = []string{"Hello"}
	separator = " "
	expected = "Hello"
	result = JoinStrings(slice, separator)
	if result != expected {
		t.Errorf("esperado '%s', mas obteve '%s'", expected, result)
	}

	slice = []string{"Hello", "Hello"}
	separator = "-"
	expected = "Hello-Hello"
	result = JoinStrings(slice, separator)
	if result != expected {
		t.Errorf("esperado '%s', mas obteve '%s'", expected, result)
	}
}
