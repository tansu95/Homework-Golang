package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Введите 3 числа ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	if a+b > c && c+b > a && c+a > b {
		fmt.Println("Треугольник существует ")
	} else {
		fmt.Println("Треугольник не существует ")
	}
	if a == b && a != c {
		fmt.Println("Треугольник равнобедренный ")
	} else if a == b && b == c {
		fmt.Println("Треугольник равносторонний ")
	} else if a != b && b != c {
		fmt.Println("Треугольник разносторониий ")
	}
}
