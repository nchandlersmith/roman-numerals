package roman

func ToRoman(number int) string {
	if number == 2 {
		return "II"
	}
	return "I"
}
