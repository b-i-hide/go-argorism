package main

import (
	"fmt"
)

func main() {
	fmt.Println(pow(3,2))
	fmt.Println(pow(2, 16))
	fmt.Println(pow(6,19))
	fmt.Println(pow(4.16, 24))
}

func pow(x float64, n int) float64 {
	v := 1.0
	for ;n > 0; n-- {
		v *= x
	}
	return v
}