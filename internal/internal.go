package internal

const (
	REGULAR_LOWER_MIN = 97
	REGULAR_LOWER_MAX = 122
	REGULAR_UPPER_MIN = 65
	REGULAR_UPPER_MAX = 90
	REGULAR_DIGIT_MIN = 48
	REGULAR_DIGIT_MAX = 57
)

func inRange(src, min, max uint8) bool {
	return src >= min && src <= max
}

func isRegularLower(src []byte) bool {
	if len(src) != 1 {
		return false
	}
	return inRange(src[0], REGULAR_LOWER_MIN, REGULAR_LOWER_MAX)
}

func isRegularUpper(src []byte) bool {
	if len(src) != 1 {
		return false
	}
	return inRange(src[0], REGULAR_UPPER_MIN, REGULAR_UPPER_MAX)
}

func isRegularDigit(src []byte) bool {
	if len(src) != 1 {
		return false
	}
	return inRange(src[0], REGULAR_DIGIT_MIN, REGULAR_DIGIT_MAX)
}
