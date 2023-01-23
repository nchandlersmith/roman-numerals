package roman

func ToRoman(number int) string {
	numeral := ""
	for i := 1; i <= number; i++ {
		numeral += "I"
	}
	if numeral == "IIII" {
		return "IV"
	}
	return numeral
}
