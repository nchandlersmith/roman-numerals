package roman_test

import (
	roman "roman-numerals"
	"testing"
)

func Test_1ReturnsI(t *testing.T) {
	result := roman.ToRoman(1)
	expected := "I"
	if result != expected {
		t.Errorf("Expected: %v Found %v", expected, result)
	}
}
