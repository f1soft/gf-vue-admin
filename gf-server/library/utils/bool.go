package utils

func BoolToInt(b bool) (i int) {
	if b {
		return 1
	}
	return 0
}