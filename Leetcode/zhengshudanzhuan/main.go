package main

import "fmt"

func main() {
	value := fanzhuan(-9463847412)
	fmt.Println(value)
}

func fanzhuan(x int) int {
	var rece int
	for x != 0 {

		b := x % 10
		if rece > 214748364 || rece == 214748364 && b > 7 {
			return 0
		}
		if rece < -214748364 || rece == -214748364 && b < -8 {
			return 0
		}

		x = x / 10
		rece = rece*10 + b

	}
	return rece
}
