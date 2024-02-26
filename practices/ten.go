package practices

import (
	"fmt"
	"math"
)

func PrimeNumber(number int) string {
	if number <= 1 {
		return "Need more characters."
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return fmt.Sprintf("%v Not is a prime number", number)
		}
	}

	return fmt.Sprintf("Yes, %v is a prime number", number)
}
