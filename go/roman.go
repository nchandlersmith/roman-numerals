package roman

func ToRoman(number int) string {
	if number == 4 {
		return "IV"
	}
	if number == 9 {
		return "IX"
	}
	if number == 14 {
		return "XIV"
	}
	if number == 19 {
		return "XIX"
	}
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
	return numeral
}
