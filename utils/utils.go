package utils

func ArrayIncludes(array []string, arg string) bool {
	for _, i := range array {
		if i == arg {
			return true
		}
	}
	return false
}
