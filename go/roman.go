package roman

func ToRoman(number int) string {
	numeral := ""
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
	return numeral
}
