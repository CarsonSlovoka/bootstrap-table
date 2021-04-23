package utils

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func In(label string, set []string) bool {
	for _, val := range set {
		if label == val {
			return true
		}
	}
	return false
}
