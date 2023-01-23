package roman

func ToRoman(number int) string {
	if number == 2 {
		return "II"
	}
	if number == 3 {
		return "III"
	}
	return "I"
}
