package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// define string value to display for FIZZ and BUZZ
const (
	FIZZ = "Fizz"
	BUZZ = "Buzz"
)

func isFizz(number int) bool {
	return number%3 == 0
}

func isBuzz(number int) bool {
	return number%5 == 0
}

func fizzBuzz(number int) {
	var buffer bytes.Buffer

	if isFizz(number) {
		buffer.WriteString(FIZZ)
	}
	if isBuzz(number) {
		buffer.WriteString(BUZZ)
	}
	if buffer.Len() == 0 {
		buffer.WriteString(strconv.Itoa(number))
	}

	fmt.Println(buffer.String())
}
