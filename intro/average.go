package main

import "fmt"

func avg(xs []float64) float64 {
	var total float64
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

func average() {
	x := []float64{98, 93, 77, 82, 83}
	fmt.Println(avg(x))
}
