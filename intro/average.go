package main

import "fmt"

func average() {
	x := [5] float64{98, 93, 77, 82, 83}
	var total float64 = 0

	for _, value := range x {
		total += value
	}

	average := total / float64(len(x))
	fmt.Println(average)
}