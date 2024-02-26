package practices

import (
	"fmt"
	"time"
)

func Weekend() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekeend! Lets drink shit")
	default:
		fmt.Println("F**ckkk!!!!")
	}
}
