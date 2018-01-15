package main

import "fmt"

func isEven(number int) bool {
	return number%2 == 0
}

func forLoop() {
	i := 1
	for i < 10 {
		if isEven(i) {
			fmt.Println(i, " is Even")
		} else {
			fmt.Println(i, " is Odd")
		}
		i++
	}
}
