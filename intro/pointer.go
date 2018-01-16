package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func incrementPointer() {
	x := 5
	zero(&x)
	fmt.Println(x)
}
