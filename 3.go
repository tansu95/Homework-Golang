package main

import "fmt"

func main() {
	var znak string
	var a, b int
	fmt.Println("Введите два числа ")
	fmt.Scan(&a,&b)
	fmt.Println("Введите арифмитическую операцию ")
	fmt.Scan(&znak)
	if znak == "+"{
		fmt.Println(a + b)
	}
	if znak == "-"{
		fmt.Println(a - b)
	}

}
