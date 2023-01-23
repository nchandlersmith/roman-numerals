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
func Test_2ReturnsII(t *testing.T) {
	result := roman.ToRoman(2)
	expected := "II"
	if result != expected {
		t.Errorf("Expected: %v Found %v", expected, result)
	}
}
func Test_3ReturnsIII(t *testing.T) {
	result := roman.ToRoman(3)
	expected := "III"
	if result != expected {
		t.Errorf("Expected: %v Found %v", expected, result)
	}
}
