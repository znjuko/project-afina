package main

import (
	"fmt"
)

func main() {
	d := []string{
		"asda",
	}
	fmt.Println(d)
	d = d[1:]
	fmt.Println(d)
}
