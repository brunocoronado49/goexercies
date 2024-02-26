package practices

func Vocals(word string) bool {
	vocals := []string{"a", "e", "i", "o", "u"}
	firstVocal := string(word[0])
	lastVocal := string(word[len(word)-1])

	if contains(vocals, firstVocal) && contains(vocals, lastVocal) && firstVocal == lastVocal {
		return true
	}
	return false
}

func contains(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}
