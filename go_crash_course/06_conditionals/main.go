package main

import "fmt"

func main() {
	x := 10
	y := 10
	color := "green"

	//if else
	if x <= y {
		fmt.Printf("%d es menor o igual que %d\n", x, y)
	} else {
		fmt.Printf("%d es mayor que %d\n", x, y)
	}

	//else if
	if color == "red" {
		fmt.Println("el color es red")
	} else if color == "blue" {
		fmt.Println("el color es blue")
	} else {
		fmt.Println("el color no es red ni blue")
	}

	//swith
	switch color {
	case "red":
		fmt.Println("el color es red")
	case "blue":
		fmt.Println("el color es blue")
	default:
		fmt.Println("el color no es red ni blue")
	}

}
