package main

import "fmt"

func main() {
	//define map
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Pedro"] = "pedro@gmail.com"
	emails["Juan"] = "juan@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Pedro"])

	//delete
	delete(emails, "Pedro")
	fmt.Println(len(emails))
}
