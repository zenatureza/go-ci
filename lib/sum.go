package lib

import "fmt"

func Sum(a float64, b float64) float64 {
	var total = a + b

	fmt.Printf("Sum(%v, %v) = %v\n", a, b, total)

	return total
}
