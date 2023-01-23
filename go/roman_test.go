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
func Test_4ReturnsIV(t *testing.T) {
	result := roman.ToRoman(4)
	expected := "IV"
	if result != expected {
		t.Errorf("Expected: %v Found %v", expected, result)
	}
}
func Test_5ReturnsV(t *testing.T) {
	result := roman.ToRoman(5)
	expected := "V"
	if result != expected {
		t.Errorf("Expected: %v Found %v", expected, result)
	}
}
func Test_6ReturnsVI(t *testing.T) {
	result := roman.ToRoman(6)
	expected := "VI"
	if result != expected {
		t.Errorf("Expected: %v Found %v", expected, result)
	}
}
func Test_9ReturnsIX(t *testing.T) {
	result := roman.ToRoman(9)
	expected := "IX"
	if result != expected {
		t.Errorf("Expected: %v Found %v", expected, result)
	}
}
func Test_10ReturnsX(t *testing.T) {
	result := roman.ToRoman(10)
	expected := "X"
	if result != expected {
		t.Errorf("Expected: %v Found %v", expected, result)
	}
}
func Test_11ReturnsXI(t *testing.T) {
	result := roman.ToRoman(11)
	expected := "XI"
	if result != expected {
		t.Errorf("Expected: %v Found %v", expected, result)
	}
}
func Test_16ReturnsXVI(t *testing.T) {
	result := roman.ToRoman(16)
	expected := "XVI"
	if result != expected {
		t.Errorf("Expected: %v Found %v", expected, result)
	}
}
func Test_27ReturnsXXVII(t *testing.T) {
	result := roman.ToRoman(27)
	expected := "XXVII"
	if result != expected {
		t.Errorf("Expected: %v Found %v", expected, result)
	}
}
