package main

import "fmt"

func main() {
	var frutas [2]string
	verduras := [2]string{"Lechuga", "Chayotes"}

	paises := []string{"Cuba", "Argentina", "Mexico", "Polonia", "Dinamarca", "Alemania"}

	frutas[0] = "Naranja"
	frutas[1] = "Peras"

	fmt.Println(frutas)
	fmt.Println(frutas[1])
	fmt.Println(verduras)
	fmt.Println("\n\n", paises)
	fmt.Println(paises[2])
	fmt.Println("longitud del arreglo", len(paises))

}
