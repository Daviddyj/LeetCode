package main

import "fmt"

func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑
	chessMap[2][3] = 2 //bai

	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
