package lib

func Contains(values [] string, key string) bool {
	for _, value := range values {
		if value == key {
			return true
		}
	}
	return false
}
