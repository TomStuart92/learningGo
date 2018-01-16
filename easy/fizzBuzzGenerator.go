package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func fizzBuzzGenerator(triggers map[int]string) func(int) string {
	return func(number int) string {
		var buffer bytes.Buffer
		for key, value := range triggers {
			if number%key == 0 {
				buffer.WriteString(value)
			}
		}
		if buffer.Len() == 0 {
			buffer.WriteString(strconv.Itoa(number))
		}
		return buffer.String()
	}
}

func run() {

	triggers := map[int]string{
		3: "Fizz",
		2: "Fang",
		5: "Buzz",
	}

	generator := fizzBuzzGenerator(triggers)
	fmt.Println(generator(1))
	fmt.Println(generator(2))
	fmt.Println(generator(3))
	fmt.Println(generator(5))
	fmt.Println(generator(6))
	fmt.Println(generator(15))
}
