package main

import "fmt"

func smallestInt(numbers []int) int {
	smallestSeen := numbers[0]

	for _, value := range numbers {
		if value < smallestSeen {
			smallestSeen = value
		}
	}

	fmt.Println("Smallest = ", smallestSeen)
	return smallestSeen
}
