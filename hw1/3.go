package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите одно число ")
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println("Число четное ")
	} else {
		fmt.Println("Число нечетное ")
	}
}
