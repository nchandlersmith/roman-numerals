package roman

func ToRoman(number int) string {
	numeral := ""
	if number >= 90 {
		numeral += "XC"
		number -= 90
	}
	for number >= 50 {
		numeral += "L"
		number -= 50
	}
	if number >= 40 {
		numeral += "XL"
		number -= 40
	}
	for number >= 10 {
		numeral += "X"
		number -= 10
	}
	if number == 9 {
		numeral += "IX"
		number -= 9
	}
	if number >= 5 {
		numeral += "V"
		number -= 5
	}
	if number == 4 {
		numeral += "IV"
		number -= 4
	}
	for number < 4 && number > 0 {
		numeral += "I"
		number -= 1
	}
	return numeral
}
