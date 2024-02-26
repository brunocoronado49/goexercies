package practices

import "fmt"

func BiggerNumber(one int, two int, three int) string {
	var aux int

	if one > two {
		aux = one
	} else if one < two {
		aux = two
	}

	if aux > three {
		return fmt.Sprintf("The value bigger is: %v", aux)
	}

	return fmt.Sprintf("The value bigger is: %v", three)
}
