package main

import "fmt"

func maps(){
	x := make(map[string]int)
	x["key"] = 10

	if name, ok := x["key"]; ok {
		fmt.Println(name, ok)
	  }
}