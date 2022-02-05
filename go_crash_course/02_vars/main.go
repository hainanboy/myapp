package main

import "fmt"

func main() {
	//main types
	//string
	//bool
	//int int8 int16 int32 int64
	//uint uint8 uin16 uint32 uint64 uintptr
	//byte - alias for uint8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128

	//using var
	var name string = "Pedro"
	var name2 = "Juan"
	var numero int = 15
	var numero2 = 30
	var isCool = true
	const cantidad = 5

	// una ves que se creo una constante esto cantidad = 3 generara un
	// error ya que no puede cambiar su valor

	//camino corto para declarar variablees
	nombre := "Luis"

	//nombre3:="Saul"
	//email:="saul@hotmail.com"
	//esto es lo mismo
	nombre3, email := "Saul", "saul@hotmail.com"

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name, numero, numero2, isCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", numero2)
	fmt.Printf("%T\n", isCool)
	fmt.Println(cantidad)
	fmt.Println(nombre)
	fmt.Println(nombre3, email)

}
