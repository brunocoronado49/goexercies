package practices

import (
	"fmt"
	"sort"
	"strings"
)

func IsAnagram(wordOne string, wordTwo string) (string, bool) {
	if wordOne == wordTwo {
		return "", false
	}

	if len(wordOne) != len(wordTwo) {
		return "", false
	}

	wordsOnes := strings.Split(wordOne, "")
	wordsTwos := strings.Split(wordTwo, "")

	sort.Strings(wordsOnes)
	sort.Strings(wordsTwos)

	word1 := strings.Join(wordsOnes, "")
	word2 := strings.Join(wordsTwos, "")

	if word1 != word2 {
		return fmt.Sprintf("%s and %s are not anagrams", word1, word2), false

	} else {
		return fmt.Sprintf("%s and %s are anagrams", word1, word2), true
	}
}
