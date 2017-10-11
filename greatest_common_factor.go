package main

import "fmt"

func main() {
	fmt.Println(gcc(42, 30))
	fmt.Println(gcc(15, 70))
	fmt.Println(lcf(5, 7))
	fmt.Println(lcf(14, 35))
}

func gcc(bigger_num, smaller_num int) int {
	if smaller_num == 0 {
		return bigger_num
	} else {
		return gcc(smaller_num, bigger_num % smaller_num)
	}
}

func lcf(bigger_num, smaller_num int) int {
	return bigger_num * smaller_num / gcc(bigger_num, smaller_num)
}
