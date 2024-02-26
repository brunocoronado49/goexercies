package practices

import "fmt"

func Fibonacci(limit int) {
	a := 0
	b := 1
	fmt.Print(a, b)

	for i := 0; i < limit; i++ {
		next := a + b
		fmt.Print(" ", next, " ")
		a = b
		b = next
	}
}
