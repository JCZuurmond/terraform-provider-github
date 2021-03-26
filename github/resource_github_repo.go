package github

func Abs(number int) int {
	var result int
	if number < 0 {
		result = -number
	} else {
		result = number
	}
	return result
}
