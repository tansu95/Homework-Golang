package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Введите три разных числа ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	if a > b && b > c {
		fmt.Println("Наибольшое число ", a, " , а наименьшое ", c)
	}
	if b > a && b > c && a < c {
		fmt.Println("Наибольшое число ", b, " , а наименьшое ", a)
	}
	if b > a && b > c && a > c {
		fmt.Println("Наибольшое число ", b, " , а наименьшое ", c)
	}
	if c > a && c > b && b > a {
		fmt.Println("Наибольшое число ", c, " , а наименьшое ", a)
	}
	if c > a && c > b && b < a {
		fmt.Println("Наибольшое число ", c, " , а наименьшое ", b)
	}
}
