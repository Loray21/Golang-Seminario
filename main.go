package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("hello wordl")
	r, err := suma(1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func suma(a, b int) (int, error) {
	if a < b {
		return 0, errors.New("el primer valor es menor que el segundo valorzs")
	}
	return a + b, nil
}
