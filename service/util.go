package service

func sliceContains(list []string, element string) bool {
	for _, n := range list {
		if element == n {
			return true
		}
	}
	return false
}