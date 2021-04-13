package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Println("Введите год ")
	fmt.Scan(&year)
	if year%4 == 0 {
		fmt.Println("Год висококосный ")
	} else if year%4 != 0 {
		fmt.Println("Год не висококосный ")
	}
}
