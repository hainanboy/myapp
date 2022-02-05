package main

import "fmt"

func saludos(name string) string {
	return "Hello " + name
}

func sumaNumeros(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(saludos("Pedrito"))
	fmt.Println("\n\n el resultado de una suma es ", sumaNumeros(5, 6))
}
