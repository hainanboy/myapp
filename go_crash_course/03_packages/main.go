package main

import (
	"fmt"
	"math"

	// /Users/albertorojo/go/src/github.com/alberto/go_crash_course/03_packages/str_utilities/reverse.go
	"03_packages/str_utilities"
)

func main() {
	fmt.Println("numero original", 4.56)
	fmt.Println("Aqui estoy truncando un numero", math.Floor(4.56))
	fmt.Println("Aqui estoy redondeando el mismo numero", math.Ceil(4.56))
	fmt.Println("\n\n\n", "raiz cuadrada de 16", math.Sqrt(16))
	fmt.Println(str_utilities.Reverse("hola mundo"))

}
