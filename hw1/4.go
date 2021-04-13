package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Println("Введите год ")
	fmt.Scan(&year)
	if year%4 != 0 {
		fmt.Println("Год невисокосный ")
	} else if year%100 != 0 {
		fmt.Println("Год високосный ")
	} else if year%100 == 0 && year%400 == 0 {
		fmt.Println("Год високосный ")
	}
}
