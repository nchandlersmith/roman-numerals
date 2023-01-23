package roman

func ToRoman(number int) string {
	numeral := ""
	for number >= 10 {
		numeral += "X"
		number -= 10
	}
	if number >= 5 {
		numeral += "V"
		number -= 5
	}
	for i := 1; i <= number; i++ {
		numeral += "I"
	}
	if numeral == "IIII" {
		return "IV"
	}
	if numeral == "VIIII" {
		return "IX"
	}
	return numeral
}
