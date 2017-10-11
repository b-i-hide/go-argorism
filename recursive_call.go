package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i, ":", fact(i))
	}
}

func fact(n int) int {
	if  n == 0 {
		return 1
	} else {
		//階乗計算
		return n * fact(n - 1)
	}
}
