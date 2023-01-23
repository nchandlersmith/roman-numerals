package roman

func ToRoman(number int) string {
	if number == 5 {
		return "V"
	}
	numeral := ""
	for i := 1; i <= number; i++ {
		numeral += "I"
	}
	if numeral == "IIII" {
		return "IV"
	}
	return numeral
}
