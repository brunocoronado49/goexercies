package practices

import "strconv"

func InvertedNumber(number string) int {
	var result string

	for i := len(number) - 1; i >= 0; i-- {
		result += string(number[i])
	}

	invNum, _ := strconv.Atoi(result)
	return invNum
}
